package actions

import (
	"github.com/abiosoft/ishell"
	"github.com/sqars/managetranslations/utils"
)

// NewUpdateTranslationFromExisting creates instance of UpdateTranslationFromExisting struct
func NewUpdateTranslationFromExisting() *UpdateTranslationFromExisting {
	return &UpdateTranslationFromExisting{
		name: "Update empty translations from existing pool",
	}
}

// UpdateTranslationFromExisting struct of operation of updating empty translations from existing pool
type UpdateTranslationFromExisting struct {
	name string
}

// GetName returns name of Action
func (a *UpdateTranslationFromExisting) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *UpdateTranslationFromExisting) PromptActionDetails(s *ishell.Shell) error {
	filePaths, err := utils.GetTranslationFilePaths()
	selectedFilePaths := []string{}
	if err != nil {
		return err
	}
	selectedFilePathIdx := s.Checklist(filePaths, "Choose file(s) to update translations:", []int{})
	for _, filePathIdx := range selectedFilePathIdx {
		selectedFilePaths = append(selectedFilePaths, filePaths[filePathIdx])
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
	modifiedTranslations := updateTranslations(translationData, existingPool)
	err = utils.SaveJSONTranslationData(filePath, modifiedTranslations)
	if err != nil {
		return err
	}
	return nil
}

func updateTranslations(data, pool utils.Translation) utils.Translation {
	for lang, keys := range data {
		for key, translation := range keys {
			existingTranslation, ok := pool[lang][key]
			if len(translation) == 0 && ok {
				if len(existingTranslation) > 0 {
					data[lang][key] = existingTranslation
				}
			}
		}
	}
	return data
}
