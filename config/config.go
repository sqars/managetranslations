package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config for app
type Config struct {
	JSONFilePattern string
	CSVFilePattern  string
}

// newDefault creates new Config instance with default values
func newDefault() Config {
	return Config{
		JSONFilePattern: "/*i18n*",
		CSVFilePattern:  "/*.csv",
	}
}

// GetConfig reads config file and returns Config struct
func GetConfig() (Config, error) {
	conf := Config{}
	wd, err := os.Getwd()
	if err != nil {
		return conf, err
	}
	path := wd + "/config.json"
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Cannot find config file. Creating one.")
		f, err := os.Create(path)
		if err != nil {
			return conf, err
		}
		conf = newDefault()
		bytes, err := json.MarshalIndent(conf, "", "\t")
		if err != nil {
			return conf, err
		}
		f.Write(bytes)
	} else {
		bytes, err := ioutil.ReadAll(f)
		if err != nil {
			return conf, err
		}
		err = json.Unmarshal(bytes, &conf)
		if err != nil {
			return conf, err
		}
	}
	f.Close()
	return conf, nil
}
