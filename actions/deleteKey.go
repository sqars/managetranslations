package actions

import (
	"fmt"

	"github.com/abiosoft/ishell"
	"github.com/sqars/managetranslations/utils"
)

// NewDeleteTranslation creates instance of DeleteTranslation struct
func NewDeleteTranslation() *DeleteTranslation {
	return &DeleteTranslation{
		name: "Delete translation key",
	}
}

// DeleteTranslation struct of operation of removing translation key
type DeleteTranslation struct {
	name string
}

// GetName returns name of Action
func (a *DeleteTranslation) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *DeleteTranslation) PromptActionDetails(s *ishell.Shell) error {
	filePaths, err := utils.GetTranslationFilePaths()
	selectedFilePaths := []string{}
	if err != nil {
		return err
	}
	selectedFilePathIdx := s.Checklist(filePaths, "Choose file(s) to delete translation:", []int{})
	for _, filePathIdx := range selectedFilePathIdx {
		selectedFilePaths = append(selectedFilePaths, filePaths[filePathIdx])
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
	modifiedTranslations := removeKey(translationData, keyToRemove)
	err = utils.SaveJSONTranslationData(filePath, modifiedTranslations)
	if err != nil {
		return err
	}
	return nil
}

func removeKey(data utils.Translation, key string) utils.Translation {
	for lang := range data {
		_, ok := data[lang][key]
		if ok {
			delete(data[lang], key)
		} else {
			fmt.Println(fmt.Sprintf(`Cant find translation key: "%s" for lang: "%s"`, key, lang))
		}
	}
	return data
}
