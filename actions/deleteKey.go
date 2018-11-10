package actions

import (
	"errors"
	"fmt"

	"github.com/sqars/managetranslations/config"
)

// NewDeleteTranslation creates instance of DeleteTranslation struct
func NewDeleteTranslation(conf config.Config) *DeleteTranslation {
	return &DeleteTranslation{
		name: "Delete translation key",
		conf: conf,
	}
}

// DeleteTranslation struct of operation of removing translation key
type DeleteTranslation struct {
	name string
	conf config.Config
}

// GetName returns name of Action
func (a *DeleteTranslation) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *DeleteTranslation) PromptActionDetails(s promptShell, d dataCollector) (ActionDetails, error) {
	details := ActionDetails{}
	selectedFilePaths, err := d.PromptFiles(
		s, "Choose file(s) to delete translation:", a.conf.JSONFilePattern,
	)
	if err != nil {
		return details, err
	}
	s.Println("Type translation key to delete:")
	translationKey := s.ReadLine()
	if len(translationKey) == 0 {
		return details, errors.New("No translation key provided")
	}
	details.selectedFilesPaths = selectedFilePaths
	details.translationKey = translationKey
	return details, nil
}

// PerformAction Performs action using collected ActionDetails
func (a *DeleteTranslation) PerformAction(d ActionDetails) error {
	err := modifyAndSaveTranslations(d, a.removeKey)
	if err != nil {
		return err
	}
	return nil
}

// RemoveKey removes providen key from translation
func (a *DeleteTranslation) removeKey(data Translation, d ActionDetails) Translation {
	for lang := range data {
		_, ok := data[lang][d.translationKey]
		if ok {
			delete(data[lang], d.translationKey)
		} else {
			fmt.Println(fmt.Sprintf(`Cant find translation key: "%s" for lang: "%s"`, d.translationKey, lang))
		}
	}
	return data
}
