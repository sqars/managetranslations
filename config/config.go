package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config for app
type Config struct {
	JSONFilePattern string
	CSVFilePattern  string
}

// GetConfig reads config file and returns Config struct
func GetConfig() (Config, error) {
	conf := Config{}
	wd, err := os.Getwd()
	if err != nil {
		return conf, err
	}
	f, err := os.Open(wd + "/config.json")
	if err != nil {
		return conf, err
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return conf, err
	}
	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
