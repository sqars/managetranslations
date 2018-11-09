package actions

import (
	"errors"

	"github.com/sqars/managetranslations/config"
	"github.com/sqars/managetranslations/utils"
)

// PromptShell represents shell inteface for PromptActionDetails
type PromptShell interface {
	ReadLine() string
	Println(...interface{})
	Checklist([]string, string, []int) []int
}

type filesCollector interface {
	getJSONConfig() string
	getCSVConfig() string
	PromptFiles(PromptShell, string, string) ([]string, error)
}

// NewDataCollector returns new DataCollector instance
func NewDataCollector(conf config.Config) *DataCollector {
	return &DataCollector{
		config: conf,
	}
}

type DataCollector struct {
	config config.Config
}

func (d *DataCollector) getCSVConfig() string {
	return d.config.CSVFilePattern
}

func (d *DataCollector) getJSONConfig() string {
	return d.config.JSONFilePattern
}

// PromptFiles asks user to select files to work with
func (d *DataCollector) PromptFiles(s PromptShell, msg, filePattern string) (selectedFilePaths []string, err error) {
	filePaths, err := utils.GetFilePaths(filePattern)
	selected := []string{}
	if err != nil {
		return nil, err
	}
	selectedFilePathIdx := s.Checklist(filePaths, msg, []int{})
	for _, filePathIdx := range selectedFilePathIdx {
		selected = append(selected, filePaths[filePathIdx])
	}
	if len(selected) == 0 {
		return nil, errors.New("No files selected")
	}
	return selected, nil
}
