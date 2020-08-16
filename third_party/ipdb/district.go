package ipdb

import "os"

// District contains the ipdb reader
type District struct {
	Reader
}

// NewDistrict return an IPDB.
func NewDistrict(name string) (IPDB, error) {
	r, err := newReader(name)
	if err == nil {
		return &District{r}, nil
	}
	return nil, err
}

// Reload database from input path
func (db *District) Reload(name string) error {

	_, err := os.Stat(name)
	if err != nil {
		return err
	}

	reader, err := newReader(name)
	if err != nil {
		return err
	}

	db.Reader = reader

	return nil
}

// Close the database, clean all data.
func (db *District) Close() error {
	return nil
}
