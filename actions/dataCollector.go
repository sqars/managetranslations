package actions

import (
	"errors"

	"github.com/sqars/managetranslations/config"
	"github.com/sqars/managetranslations/utils"
)

// PromptShell represents shell inteface for PromptActionDetails
type promptShell interface {
	ReadLine() string
	Println(...interface{})
	Checklist([]string, string, []int) []int
}

type dataCollector interface {
	PromptFiles(promptShell, string, string) ([]string, error)
	PromptText(promptShell, string) (string, error)
}

// NewDataCollector returns new DataCollector instance
func NewDataCollector(conf config.Config) *DataCollector {
	return &DataCollector{
		config: conf,
	}
}

// DataCollector is used to collect necesary input from user, includes app config
type DataCollector struct {
	config config.Config
}

// PromptFiles asks user to select files to work with
func (d *DataCollector) PromptFiles(s promptShell, msg, filePattern string) (selectedFilePaths []string, err error) {
	filePaths, err := GetFilePaths(filePattern)
	selected := []string{}
	if err != nil {
		return nil, err
	}
	filePaths = append([]string{"All files"}, filePaths...)
	selectedFilePathIdx := s.Checklist(filePaths, msg, []int{})
	if utils.ContainsInt(selectedFilePathIdx, 0) {
		selected = filePaths
	} else {
		for _, filePathIdx := range selectedFilePathIdx {
			selected = append(selected, filePaths[filePathIdx])
		}
	}
	selected = utils.FilterStr(selected, "All files")
	if len(selected) == 0 {
		return nil, errors.New("No files selected")
	}
	return selected, nil
}

// PromptText method for getting text from user
func (d *DataCollector) PromptText(s promptShell, msg string) (string, error) {
	s.Println(msg)
	text := s.ReadLine()
	if len(text) == 0 {
		return text, errors.New("No value provided")
	}
	return text, nil
}
