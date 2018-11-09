package actions

import (
	"github.com/abiosoft/ishell"
	"github.com/sqars/managetranslations/config"
	"github.com/sqars/managetranslations/utils"
)

// NewDeleteTranslation creates instance of DeleteTranslation struct
func NewDeleteTranslation(conf config.Config) *DeleteTranslation {
	return &DeleteTranslation{
		name:   "Delete translation key",
		config: conf,
	}
}

// DeleteTranslation struct of operation of removing translation key
type DeleteTranslation struct {
	name   string
	config config.Config
}

// GetName returns name of Action
func (a *DeleteTranslation) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *DeleteTranslation) PromptActionDetails(s *ishell.Shell) error {
	selectedFilePaths, err := utils.PromptFiles(
		s, "Choose file(s) to delete translation:", a.config.JSONFilePattern,
	)
	if err != nil {
		return err
	}
	s.Println("Type translation key to delete:")
	translationKey := s.ReadLine()
	for _, path := range selectedFilePaths {
		err := a.Perform(path, translationKey)
		if err != nil {
			return err
		}
	}
	return nil
}

// Perform removes translation key to file
func (a *DeleteTranslation) Perform(filePath, keyToRemove string) error {
	translationData, err := utils.GetJSONTranslationData(filePath)
	if err != nil {
		return err
	}
	modifiedTranslations := utils.RemoveKey(translationData, keyToRemove)
	err = utils.SaveJSONTranslationData(filePath, modifiedTranslations)
	if err != nil {
		return err
	}
	return nil
}
