package actions

import (
	"github.com/abiosoft/ishell"
	"github.com/sqars/managetranslations/config"
	"github.com/sqars/managetranslations/utils"
)

// NewAddTranslation creates instance of AddTranslation struct
func NewAddTranslation(conf config.Config) *AddTranslation {
	return &AddTranslation{
		name:   "Add translation key",
		config: conf,
	}
}

// AddTranslation struct of operation of adding translation key
type AddTranslation struct {
	name   string
	config config.Config
}

// GetName returns name of Action
func (a *AddTranslation) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *AddTranslation) PromptActionDetails(s *ishell.Shell) error {
	selectedFilePaths, err := utils.PromptFiles(
		s, "Select file(s) to add translation key", a.config.JSONFilePattern,
	)
	if err != nil {
		return err
	}
	s.Println("Type translation key to add:")
	translationKey := s.ReadLine()
	for _, path := range selectedFilePaths {
		err := a.Perform(path, translationKey)
		if err != nil {
			return err
		}
	}
	return nil
}

// Perform adds translation key to file
func (a *AddTranslation) Perform(filePath, keyToAdd string) error {
	translationData, err := utils.GetJSONTranslationData(filePath)
	if err != nil {
		return err
	}
	modifiedTranslations := utils.Addkey(translationData, keyToAdd)
	err = utils.SaveJSONTranslationData(filePath, modifiedTranslations)
	if err != nil {
		return err
	}
	return nil
}
