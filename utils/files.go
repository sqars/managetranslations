package utils

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/abiosoft/ishell"
)

// GetFilePaths search for required files and returns file paths
// in working directory
func GetFilePaths(filePattern string) (paths []string, err error) {
	filePaths := []string{}
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		paths, err := filepath.Glob(path + filePattern)
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

// GetJSONTranslationData opens JSON file and returns struct with translation
func GetJSONTranslationData(path string) (translation Translation, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	data := Translation{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetCSVTranslationData opens CSV files and returns struct with translation
func GetCSVTranslationData(paths []string) (translation Translation, err error) {
	translations := []Translation{}
	for _, filePath := range paths {
		fileData, err := GetSingleCSVTranslationData(filePath)
		if err != nil {
			return nil, err
		}
		translations = append(translations, fileData)
	}
	return mergeTranslations(translations), nil
}

// GetSingleCSVTranslationData opens CSV file and returns struct with translation
func GetSingleCSVTranslationData(path string) (translation Translation, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data := [][]string{}
	reader := csv.NewReader(bufio.NewReader(f))
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		data = append(data, line)
	}
	return csvToTranslationFormat(data), nil
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

// PromptFiles asks user to select files to work with
func PromptFiles(s *ishell.Shell, msg, filePattern string) (selectedFilePaths []string, err error) {
	filePaths, err := GetFilePaths(filePattern)
	selected := []string{}
	if err != nil {
		return nil, err
	}
	selectedFilePathIdx := s.Checklist(filePaths, msg, []int{})
	for _, filePathIdx := range selectedFilePathIdx {
		selected = append(selected, filePaths[filePathIdx])
	}
	return selected, nil
}
