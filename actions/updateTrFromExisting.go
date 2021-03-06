package actions

import "github.com/sqars/managetranslations/config"

// NewUpdateTranslationFromExisting creates instance of UpdateTranslationFromExisting struct
func NewUpdateTranslationFromExisting(conf config.Config) *UpdateTranslationFromExisting {
	return &UpdateTranslationFromExisting{
		name: "Update empty translations from existing pool",
		conf: conf,
	}
}

// UpdateTranslationFromExisting struct of operation of updating empty translations from existing pool
type UpdateTranslationFromExisting struct {
	name string
	conf config.Config
}

// GetName returns name of Action
func (a *UpdateTranslationFromExisting) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *UpdateTranslationFromExisting) PromptActionDetails(s promptShell, d dataCollector) (ActionDetails, error) {
	details := ActionDetails{}
	selectedFilePaths, err := d.PromptFiles(
		s, "Select file(s) to update with existing translations", a.conf.JSONFilePattern,
	)
	if err != nil {
		return details, err
	}
	existingPool, err := GetExistingPool()
	if err != nil {
		return details, err
	}
	details.translations = existingPool
	details.selectedFilesPaths = selectedFilePaths
	return details, nil
}

// PerformAction Performs action using collected ActionDetails
func (a *UpdateTranslationFromExisting) PerformAction(d ActionDetails) error {
	err := modifyAndSaveTranslations(d, a.updateTranslations)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTranslations updates translation from pool(existing translations)
func (a *UpdateTranslationFromExisting) updateTranslations(data Translation, d ActionDetails) Translation {
	return UpdateTranslations(data, d.translations)
}
