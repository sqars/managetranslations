package actions

import (
	"github.com/sqars/managetranslations/utils"
)

// NewUpdateFromCSV creates instance of UpdateFromCSV struct
func NewUpdateFromCSV() *UpdateFromCSV {
	return &UpdateFromCSV{
		name: "Update empty translations from CSV file",
	}
}

// UpdateFromCSV struct of operation of updating empty translations from CSV file
type UpdateFromCSV struct {
	name string
}

// GetName returns name of Action
func (a *UpdateFromCSV) GetName() string {
	return a.name
}

// GetModifierFn returns function which modifies translations for action
func (a *UpdateFromCSV) GetModifierFn() TranslationModifier {
	return a.updateTranslations
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *UpdateFromCSV) PromptActionDetails(s PromptShell, d filesCollector) (ActionDetails, error) {
	details := ActionDetails{}
	selectedTransationFilePaths, err := d.PromptFiles(
		s,
		"Select translation file(s) to update from CSV",
		d.getJSONConfig(),
	)
	if err != nil {
		return details, err
	}
	selectedCSVFilePaths, err := d.PromptFiles(
		s,
		"Select CSV file(s) with translations",
		d.getCSVConfig(),
	)
	if err != nil {
		return details, err
	}
	csvData, err := utils.GetCSVTranslationData(selectedCSVFilePaths)
	if err != nil {
		return details, err
	}
	details.selectedFilesPaths = selectedTransationFilePaths
	details.translations = csvData
	return details, nil
}

// UpdateTranslations updates translation from pool(existing translations)
func (a *UpdateFromCSV) updateTranslations(data utils.Translation, d ActionDetails) utils.Translation {
	return utils.UpdateTranslations(data, d.translations)
}
