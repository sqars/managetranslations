package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Translation represents translation format written in JSON file
type Translation map[string]map[string]string

// GetTranslationFilePaths search for translation file paths
// in working directory and returns it
func GetTranslationFilePaths() (paths []string, err error) {
	filePaths := []string{}
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path + "/*i18n*")
		paths, err := filepath.Glob(path + "/*i18n*")
		if err != nil {
			return err
		}
		filePaths = append(filePaths, paths...)
		return nil
	}
	filepath.Walk(wd, walkFn)
	if err != nil {
		return nil, err
	}
	return filePaths, nil
}

// GetJSONTranslationData opens JSON file an returns struct with translation
func GetJSONTranslationData(path string) (translation Translation, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, _ := ioutil.ReadAll(f)
	data := Translation{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// SaveJSONTranslationData saves modified translation data to file
func SaveJSONTranslationData(path string, data Translation) error {
	bytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, bytes, 0660)
	if err != nil {
		return err
	}
	return nil
}
