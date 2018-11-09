package actions

import (
	"github.com/abiosoft/ishell"
	"github.com/sqars/managetranslations/config"
	"github.com/sqars/managetranslations/utils"
)

// NewUpdateFromCSV creates instance of UpdateFromCSV struct
func NewUpdateFromCSV(conf config.Config) *UpdateFromCSV {
	return &UpdateFromCSV{
		name:   "Update empty translations from CSV file",
		config: conf,
	}
}

// UpdateFromCSV struct of operation of updating empty translations from CSV file
type UpdateFromCSV struct {
	name   string
	config config.Config
}

// GetName returns name of Action
func (a *UpdateFromCSV) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *UpdateFromCSV) PromptActionDetails(s *ishell.Shell) error {
	selectedTransationFilePaths, err := utils.PromptFiles(
		s,
		"Select translation file(s) to update from CSV",
		a.config.JSONFilePattern,
	)
	if err != nil {
		return err
	}
	selectedCSVFilePaths, err := utils.PromptFiles(
		s,
		"Select CSV file(s) with translations",
		a.config.CSVFilePattern,
	)
	if err != nil {
		return err
	}
	csvData, err := utils.GetCSVTranslationData(selectedCSVFilePaths)
	if err != nil {
		return err
	}
	for _, path := range selectedTransationFilePaths {
		err := a.Perform(path, csvData)
		if err != nil {
			return err
		}
	}
	return nil
}

// Perform updates missing translations with existing one from pool
func (a *UpdateFromCSV) Perform(filePath string, csvData utils.Translation) error {
	translationData, err := utils.GetJSONTranslationData(filePath)
	if err != nil {
		return err
	}
	modifiedTranslations := utils.UpdateTranslations(translationData, csvData)
	err = utils.SaveJSONTranslationData(filePath, modifiedTranslations)
	if err != nil {
		return err
	}
	return nil
}
