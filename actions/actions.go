package actions

import "github.com/sqars/managetranslations/utils"

// Action represents action to perform on json file
type Action interface {
	GetName() string
	PromptActionDetails(PromptShell, filesCollector) (ActionDetails, error)
	GetModifierFn() TranslationModifier
}

// ActionDetails represents details for action to perform
type ActionDetails struct {
	selectedFilesPaths []string
	translations       utils.Translation
	translationKey     string
}

// TranslationModifier represents function which modifies translations
type TranslationModifier func(utils.Translation, ActionDetails) utils.Translation

// PerformAction Performs action using collected ActionDetails
func (a *ActionDetails) PerformAction(m TranslationModifier) error {
	for _, filePath := range a.selectedFilesPaths {
		translationData, err := utils.GetJSONTranslationData(filePath)
		if err != nil {
			return err
		}
		modifiedTranslations := m(translationData, *a)
		err = utils.SaveJSONTranslationData(filePath, modifiedTranslations)
		if err != nil {
			return err
		}
	}
	return nil
}
