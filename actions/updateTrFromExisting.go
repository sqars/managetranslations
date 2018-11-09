package actions

import (
	"github.com/abiosoft/ishell"
	"github.com/sqars/managetranslations/config"
	"github.com/sqars/managetranslations/utils"
)

// NewUpdateTranslationFromExisting creates instance of UpdateTranslationFromExisting struct
func NewUpdateTranslationFromExisting(conf config.Config) *UpdateTranslationFromExisting {
	return &UpdateTranslationFromExisting{
		name:   "Update empty translations from existing pool",
		config: conf,
	}
}

// UpdateTranslationFromExisting struct of operation of updating empty translations from existing pool
type UpdateTranslationFromExisting struct {
	name   string
	config config.Config
}

// GetName returns name of Action
func (a *UpdateTranslationFromExisting) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *UpdateTranslationFromExisting) PromptActionDetails(s *ishell.Shell) error {
	selectedFilePaths, err := utils.PromptFiles(
		s, "Select file(s) to update with existing translations", a.config.JSONFilePattern,
	)
	if err != nil {
		return err
	}
	existingPool, err := utils.GetExistingPool()
	if err != nil {
		return err
	}
	for _, path := range selectedFilePaths {
		err := a.Perform(path, existingPool)
		if err != nil {
			return err
		}
	}
	return nil
}

// Perform updates missing translations with existing one from pool
func (a *UpdateTranslationFromExisting) Perform(filePath string, existingPool utils.Translation) error {
	translationData, err := utils.GetJSONTranslationData(filePath)
	if err != nil {
		return err
	}
	modifiedTranslations := utils.UpdateTranslations(translationData, existingPool)
	err = utils.SaveJSONTranslationData(filePath, modifiedTranslations)
	if err != nil {
		return err
	}
	return nil
}
