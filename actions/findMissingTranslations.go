package actions

import (
	"fmt"

	"github.com/sqars/managetranslations/config"
)

// NewFindMissingTranslations creates instance of FindMissingTranslations struct
func NewFindMissingTranslations(conf config.Config) *FindMissingTranslations {
	return &FindMissingTranslations{
		name: "Find missing translations",
		conf: conf,
	}
}

// FindMissingTranslations struct of operation of finding all missing translations
type FindMissingTranslations struct {
	name string
	conf config.Config
}

// GetName returns name of Action
func (a *FindMissingTranslations) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *FindMissingTranslations) PromptActionDetails(s promptShell, d dataCollector) (ActionDetails, error) {
	details := ActionDetails{}
	paths, err := GetFilePaths(a.conf.JSONFilePattern)
	if err != nil {
		return details, err
	}
	details.selectedFilesPaths = paths
	return details, nil
}

// PerformAction Performs action using collected ActionDetails
func (a *FindMissingTranslations) PerformAction(d ActionDetails) error {
	for _, filePath := range d.selectedFilesPaths {
		trasnaltion, err := GetJSONTranslationData(filePath)
		if err != nil {
			return err
		}
		for lang, values := range trasnaltion {
			for key, value := range values {
				if len(value) == 0 {
					fmt.Println(fmt.Sprintf(`Missing translation for "%s" for lang "%s" in %s`, key, lang, filePath))
				}
			}
		}
	}
	return nil
}
