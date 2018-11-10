package actions

// Action represents action to perform on json file
type Action interface {
	GetName() string
	PromptActionDetails(promptShell, dataCollector) (ActionDetails, error)
	GetModifierFn() TranslationModifier
}

// ActionDetails represents details for action to perform
type ActionDetails struct {
	selectedFilesPaths []string
	translations       Translation
	translationKey     string
}

// PerformAction Performs action using collected ActionDetails
func (a *ActionDetails) PerformAction(m TranslationModifier) error {
	for _, filePath := range a.selectedFilesPaths {
		translationData, err := GetJSONTranslationData(filePath)
		if err != nil {
			return err
		}
		modifiedTranslations := m(translationData, *a)
		err = SaveJSONTranslationData(filePath, modifiedTranslations)
		if err != nil {
			return err
		}
	}
	return nil
}
