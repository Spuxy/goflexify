package reader

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DbPassword string
	DbDatabase string
	DbUser     string
	DbUsername string
	DbHost     string
	DbTable    string
	DbPort     string
}
type Reader interface {
	ReadGivenFileIntoMap() (Config, error)
}
type ReaderFromINI struct {
	filename string
}

func (r *ReaderFromINI) ReadGivenFileIntoMap() (config Config, err error) {
	_, err = os.Stat(r.filename)
	if err != nil {
		return config, errors.New("Config file is missing")
	}

	if _, err = toml.DecodeFile(r.filename, &config); err != nil {
		return config, err
	}
	return config, nil
}

func CreateReader(filename string) *ReaderFromINI {
	return &ReaderFromINI{filename: filename}
}
