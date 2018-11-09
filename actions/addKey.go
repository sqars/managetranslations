package actions

import (
	"errors"

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
	name     string
	modifier TranslationModifier
}

// GetName returns name of Action
func (a *AddTranslation) GetName() string {
	return a.name
}

// GetModifierFn returns function which modifies translations for action
func (a *AddTranslation) GetModifierFn() TranslationModifier {
	return a.addKey
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *AddTranslation) PromptActionDetails(s PromptShell, d filesCollector) (ActionDetails, error) {
	details := ActionDetails{}
	selectedFilePaths, err := d.PromptFiles(
		s, "Select file(s) to add translation key", d.getJSONConfig(),
	)
	if err != nil {
		return details, err
	}
	s.Println("Type translation key to add:")
	translationKey := s.ReadLine()
	if len(translationKey) == 0 {
		return details, errors.New("No translation key provided")
	}
	details.selectedFilesPaths = selectedFilePaths
	details.translationKey = translationKey
	return details, nil
}

// Addkey adds empty translation key to translation
func (a *AddTranslation) addKey(data utils.Translation, d ActionDetails) utils.Translation {
	for lang := range data {
		data[lang][d.translationKey] = ""
	}
	return data
}
