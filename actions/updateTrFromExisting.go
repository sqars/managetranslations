package actions

// NewUpdateTranslationFromExisting creates instance of UpdateTranslationFromExisting struct
func NewUpdateTranslationFromExisting() *UpdateTranslationFromExisting {
	return &UpdateTranslationFromExisting{
		name: "Update empty translations from existing pool",
	}
}

// UpdateTranslationFromExisting struct of operation of updating empty translations from existing pool
type UpdateTranslationFromExisting struct {
	name     string
	modifier TranslationModifier
}

// GetName returns name of Action
func (a *UpdateTranslationFromExisting) GetName() string {
	return a.name
}

// GetModifierFn returns function which modifies translations for action
func (a *UpdateTranslationFromExisting) GetModifierFn() TranslationModifier {
	return a.updateTranslations
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *UpdateTranslationFromExisting) PromptActionDetails(s promptShell, d dataCollector) (ActionDetails, error) {
	details := ActionDetails{}
	selectedFilePaths, err := d.PromptFiles(
		s, "Select file(s) to update with existing translations", d.getJSONConfig(),
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

// UpdateTranslations updates translation from pool(existing translations)
func (a *UpdateTranslationFromExisting) updateTranslations(data Translation, d ActionDetails) Translation {
	return UpdateTranslations(data, d.translations)
}
