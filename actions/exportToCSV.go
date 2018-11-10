package actions

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/sqars/managetranslations/config"
)

// NewExportToCSV creates instance of ExportToCSV struct
func NewExportToCSV(conf config.Config) *ExportToCSV {
	return &ExportToCSV{
		name: "Export translation to CSV",
		conf: conf,
	}
}

// ExportToCSV struct of operation of adding translation key
type ExportToCSV struct {
	name string
	conf config.Config
}

// GetName returns name of Action
func (a *ExportToCSV) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *ExportToCSV) PromptActionDetails(s promptShell, d dataCollector) (ActionDetails, error) {
	details := ActionDetails{}
	selectedFilePaths, err := d.PromptFiles(
		s, "Select file(s) to export", a.conf.JSONFilePattern,
	)
	if err != nil {
		return details, err
	}
	details.selectedFilesPaths = selectedFilePaths
	return details, nil
}

// PerformAction Performs action using collected ActionDetails
func (a *ExportToCSV) PerformAction(d ActionDetails) error {
	for _, lang := range a.conf.Languages {
		langTranslations := [][]string{}
		langTranslations = append(langTranslations, []string{"en", lang, "file path"})
		for _, filePath := range d.selectedFilesPaths {
			trasnaltion, err := GetJSONTranslationData(filePath)
			if err != nil {
				return err
			}
			fileValues := transformToCSVformat(trasnaltion[lang], filePath)
			langTranslations = append(langTranslations, fileValues...)
		}
		err := saveToCSV(langTranslations, lang)
		if err != nil {
			return err
		}
	}
	return nil
}

func transformToCSVformat(vals map[string]string, path string) [][]string {
	output := [][]string{}
	for key, value := range vals {
		output = append(output, []string{key, value, path})
	}
	return output
}

func saveToCSV(contents [][]string, lang string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	f, err := os.Create(wd + fmt.Sprintf("/%s_translations.csv", lang))
	writer := csv.NewWriter(f)
	err = writer.WriteAll(contents)
	if err != nil {
		return err
	}
	return nil
}
