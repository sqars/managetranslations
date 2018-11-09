package actions

import (
	"github.com/abiosoft/ishell"
	"github.com/sqars/managetranslations/utils"
)

// NewAddTranslation creates instance of AddTranslation struct
func NewAddTranslation() *AddTranslation {
	return &AddTranslation{
		name: "Add translation key",
	}
}

// AddTranslation struct of operation of adding translation key
type AddTranslation struct {
	name string
}

// GetName returns name of Action
func (a *AddTranslation) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *AddTranslation) PromptActionDetails(s *ishell.Shell) error {
	filePaths, err := utils.GetTranslationFilePaths()
	selectedFilePaths := []string{}
	if err != nil {
		return err
	}
	selectedFilePathIdx := s.Checklist(filePaths, "Choose file(s) to add translation:", []int{})
	for _, filePathIdx := range selectedFilePathIdx {
		selectedFilePaths = append(selectedFilePaths, filePaths[filePathIdx])
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
	modifiedTranslations := addkey(translationData, keyToAdd)
	err = utils.SaveJSONTranslationData(filePath, modifiedTranslations)
	if err != nil {
		return err
	}
	return nil
}

func addkey(data utils.Translation, key string) utils.Translation {
	for lang := range data {
		data[lang][key] = ""
	}
	return data
}
