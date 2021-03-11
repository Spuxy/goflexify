package reader

import (
	"log"
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
		log.Fatal("Config file is missing: ", r.filename)
		return config, nil
	}

	if _, err = toml.DecodeFile(r.filename, &config); err != nil {
		log.Fatal(err)
		return config, err
	}
	return config, nil
}

func CreateReader(filename string) *ReaderFromINI {
	return &ReaderFromINI{filename: filename}
}
