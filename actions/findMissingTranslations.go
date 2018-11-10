package actions

import (
	"fmt"
)

// NewFindMissingTranslations creates instance of FindMissingTranslations struct
func NewFindMissingTranslations() *FindMissingTranslations {
	return &FindMissingTranslations{
		name: "Find missing translations",
	}
}

// FindMissingTranslations struct of operation of finding all missing translations
type FindMissingTranslations struct {
	name string
}

// GetName returns name of Action
func (a *FindMissingTranslations) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *FindMissingTranslations) PromptActionDetails(s promptShell, d dataCollector) (ActionDetails, error) {
	details := ActionDetails{}
	paths, err := GetFilePaths(d.getJSONConfig())
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
