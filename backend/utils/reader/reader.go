package reader

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

const filename string = "properties.ini"

type Config struct {
	DbPassword string
	DbDatabase string
	DbUser     string
	DbUsername string
	DbHost     string
	DbTable    string
	DbPort     string
}

func ReadGivenFileIntoMap() (config Config, err error) {
	_, err = os.Stat(filename)
	if err != nil {
		log.Fatal("Config file is missing: ", filename)
		return config, nil
	}

	if _, err = toml.DecodeFile(filename, &config); err != nil {
		log.Fatal(err)
		return config, err
	}
	return config, nil
}
