package ipdb

import "time"

// FillObject fill the ipdb search result
type FillObject interface {
	FillIPDB([]string, []string, int)
}

// Reader hold data which read from .ipdb, and privode same function to access the data
type Reader interface {
	Build() time.Time
	IsIPv4Support() bool
	IsIPv6Support() bool
	Fields() []string
	Languages() []string

	Fill(addr, language string, f FillObject) error
	Find(addr, language string) ([]string, int, error)
	FindMap(addr, language string) (map[string]string, int, error)
}

// IPDB an database contains ip location
type IPDB interface {
	Reader
	Close() error
	Reload(name string) error
}
