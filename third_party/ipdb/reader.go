package ipdb

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"
	"unsafe"
)

const ChineseLanguage = "CN"
const EnglishLanguage = "EN"

const IPv4 = 0x01
const IPv6 = 0x02

var (
	ErrFileSize = errors.New("IP Database file size error.")
	ErrMetaData = errors.New("IP Database metadata error.")
	ErrReadFull = errors.New("IP Database ReadFull error.")

	ErrDatabaseError = errors.New("database error")

	ErrIPFormat = errors.New("Query IP Format error.")

	ErrNoSupportLanguage = errors.New("language not support")
	ErrNoSupportIPv4     = errors.New("IPv4 not support")
	ErrNoSupportIPv6     = errors.New("IPv6 not support")

	ErrDataNotExists = errors.New("data is not exists")

	ErrFillToNil = errors.New("the fill object is nil")
)

type MetaData struct {
	Build     int64          `json:"build"`
	IPVersion uint16         `json:"ip_version"`
	Languages map[string]int `json:"languages"`
	NodeCount int            `json:"node_count"`
	TotalSize int            `json:"total_size"`
	Fields    []string       `json:"fields"`
}

type reader struct {
	fileSize  int
	nodeCount int
	v4offset  int

	meta MetaData
	data []byte
}

func newReader(name string) (Reader, error) {
	var err error
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat(name)
	if err != nil {
		return nil, err
	}
	fileSize := int(fileInfo.Size())
	if fileSize < 4 {
		return nil, ErrFileSize
	}
	body, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, ErrReadFull
	}
	var meta MetaData
	metaLength := int(binary.BigEndian.Uint32(body[0:4]))
	if fileSize < (4 + metaLength) {
		return nil, ErrFileSize
	}
	if err := json.Unmarshal(body[4:4+metaLength], &meta); err != nil {
		return nil, err
	}
	if len(meta.Languages) == 0 || len(meta.Fields) == 0 {
		return nil, ErrMetaData
	}
	if fileSize != (4 + metaLength + meta.TotalSize) {
		return nil, ErrFileSize
	}

	db := &reader{
		fileSize:  fileSize,
		nodeCount: meta.NodeCount,

		meta: meta,
		data: body[4+metaLength:],
	}

	if db.v4offset == 0 {
		node := 0
		for i := 0; i < 96 && node < db.nodeCount; i++ {
			if i >= 80 {
				node = db.readNode(node, 1)
			} else {
				node = db.readNode(node, 0)
			}
		}
		db.v4offset = node
	}

	return db, nil
}

func (db *reader) Fill(addr, language string, f FillObject) error {
	if f == nil {
		return ErrFillToNil
	}

	values, mask, err := db.find1(addr, language)
	if err == nil {
		fileds := make([]string, len(db.meta.Fields))
		copy(fileds, db.meta.Fields)
		f.FillIPDB(fileds, values, mask)
		return err
	}

	return nil
}

func (db *reader) FindMap(addr, language string) (map[string]string, int, error) {
	values, mask, err := db.find1(addr, language)
	if err == nil {
		mFiled := make(map[string]string, len(db.meta.Fields))
		for i := range db.meta.Fields {
			mFiled[db.meta.Fields[i]] = values[i]
		}
		return mFiled, mask, err
		// return mFiled, mask, nil
	}

	return nil, mask, err
}

func (db *reader) Find(addr, language string) ([]string, int, error) {
	return db.find1(addr, language)
}

func (db *reader) find0(addr string) ([]byte, int, error) {

	var err error
	var node int
	var maskLength int

	ipv := net.ParseIP(addr)
	if ip := ipv.To4(); ip != nil {
		if !db.IsIPv4Support() {
			return nil, maskLength, ErrNoSupportIPv4
		}

		node, maskLength, err = db.search(ip, 32)
	} else if ip := ipv.To16(); ip != nil {
		if !db.IsIPv6Support() {
			return nil, maskLength, ErrNoSupportIPv6
		}

		node, maskLength, err = db.search(ip, 128)
	} else {
		return nil, maskLength, ErrIPFormat
	}

	if err != nil || node < 0 {
		return nil, maskLength, err
	}

	body, err := db.resolve(node)
	if err != nil {
		return nil, maskLength, err
	}

	return body, maskLength, nil
}

func (db *reader) find1(addr, language string) ([]string, int, error) {

	off, ok := db.meta.Languages[language]
	if !ok {
		return nil, 0, ErrNoSupportLanguage
	}

	body, maskLength, err := db.find0(addr)
	if err != nil {
		return nil, maskLength, err
	}

	str := (*string)(unsafe.Pointer(&body))
	tmp := strings.Split(*str, "\t")

	if (off + len(db.meta.Fields)) > len(tmp) {
		return nil, maskLength, ErrDatabaseError
	}

	return tmp[off : off+len(db.meta.Fields)], maskLength, nil
}

func (db *reader) search(ip net.IP, bitCount int) (int, int, error) {

	var node int
	var maskLength int

	if bitCount == 32 {
		node = db.v4offset
	} else {
		node = 0
	}

	for bitOffset := 0; bitOffset < bitCount; bitOffset++ {
		if node > db.nodeCount {
			break
		}
		maskLength = bitOffset
		node = db.readNode(node, ((0xFF&int(ip[bitOffset>>3]))>>uint(7-(bitOffset%8)))&1)
	}

	if node > db.nodeCount {
		return node, maskLength, nil
	}

	return -1, maskLength, ErrDataNotExists
}

func (db *reader) readNode(node, index int) int {
	off := node*8 + index*4
	return int(binary.BigEndian.Uint32(db.data[off : off+4]))
}

func (db *reader) resolve(node int) ([]byte, error) {
	resolved := node - db.nodeCount + db.nodeCount*8
	if resolved >= db.fileSize {
		return nil, ErrDatabaseError
	}

	size := int(binary.BigEndian.Uint16(db.data[resolved : resolved+2]))
	if (resolved + 2 + size) > len(db.data) {
		return nil, ErrDatabaseError
	}
	bytes := db.data[resolved+2 : resolved+2+size]

	return bytes, nil
}

func (db *reader) IsIPv4Support() bool {
	return (int(db.meta.IPVersion) & IPv4) == IPv4
}

func (db *reader) IsIPv6Support() bool {
	return (int(db.meta.IPVersion) & IPv6) == IPv6
}

func (db *reader) Build() time.Time {
	return time.Unix(db.meta.Build, 0).In(time.UTC)
}

func (db *reader) Languages() []string {
	ls := make([]string, 0, len(db.meta.Languages))
	for k := range db.meta.Languages {
		ls = append(ls, k)
	}
	return ls
}

func (db *reader) Fields() []string {
	fields := make([]string, len(db.meta.Fields))
	copy(fields, db.meta.Fields)
	return fields
}
