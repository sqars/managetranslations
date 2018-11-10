package actions

import (
	"fmt"

	"github.com/sqars/managetranslations/config"
)

// NewSearchKey creates instance of SearchKey struct
func NewSearchKey(conf config.Config) *SearchKey {
	return &SearchKey{
		name: "Search translation key",
		conf: conf,
	}
}

// SearchKey struct of operation of adding translation key
type SearchKey struct {
	name string
	conf config.Config
}

// GetName returns name of Action
func (a *SearchKey) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *SearchKey) PromptActionDetails(s promptShell, d dataCollector) (ActionDetails, error) {
	details := ActionDetails{}
	searchKey, err := d.PromptText(s, "Type translation key to search:")
	if err != nil {
		return details, err
	}
	paths, err := GetFilePaths(a.conf.JSONFilePattern)
	if err != nil {
		return details, err
	}
	details.selectedFilesPaths = paths
	details.translationKey = searchKey
	return details, nil
}

// PerformAction searchs for translation key within all files
func (a *SearchKey) PerformAction(d ActionDetails) error {
	for _, filePath := range d.selectedFilesPaths {
		trasnaltion, err := GetJSONTranslationData(filePath)
		if err != nil {
			return err
		}
		for lang, values := range trasnaltion {
			_, ok := values[d.translationKey]
			if ok {
				fmt.Println(fmt.Sprintf(`Found "%s" lang "%s" in %s`, d.translationKey, lang, filePath))
			}
		}
	}
	return nil
}
