package actions

import "github.com/sqars/managetranslations/config"

// NewAddTranslation creates instance of AddTranslation struct
func NewAddTranslation(conf config.Config) *AddTranslation {
	return &AddTranslation{
		name: "Add translation key",
		conf: conf,
	}
}

// AddTranslation struct of operation of adding translation key
type AddTranslation struct {
	name string
	conf config.Config
}

// GetName returns name of Action
func (a *AddTranslation) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *AddTranslation) PromptActionDetails(s promptShell, d dataCollector) (ActionDetails, error) {
	details := ActionDetails{}
	selectedFilePaths, err := d.PromptFiles(
		s, "Select file(s) to add translation key", a.conf.JSONFilePattern,
	)
	if err != nil {
		return details, err
	}
	translationKey, err := d.PromptText(s, "Type translation key to add:")
	if err != nil {
		return details, err
	}
	details.selectedFilesPaths = selectedFilePaths
	details.translationKey = translationKey
	return details, nil
}

// PerformAction Performs action using collected ActionDetails
func (a *AddTranslation) PerformAction(d ActionDetails) error {
	err := modifyAndSaveTranslations(d, a.addKey)
	if err != nil {
		return err
	}
	return nil
}

// Addkey adds empty translation key to translation
func (a *AddTranslation) addKey(data Translation, d ActionDetails) Translation {
	for lang := range data {
		data[lang][d.translationKey] = ""
	}
	return data
}
