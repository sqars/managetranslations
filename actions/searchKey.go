package actions

import (
	"errors"
)

// NewSearchKey creates instance of SearchKey struct
func NewSearchKey() *SearchKey {
	return &SearchKey{
		name: "Search translation key",
	}
}

// SearchKey struct of operation of adding translation key
type SearchKey struct {
	name     string
	modifier TranslationModifier
}

// GetModifierFn returns function which modifies translations for action
func (a *SearchKey) GetModifierFn() TranslationModifier {
	return a.modifier
}

// GetName returns name of Action
func (a *SearchKey) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *SearchKey) PromptActionDetails(s PromptShell, d filesCollector) (ActionDetails, error) {
	details := ActionDetails{}
	s.Println("Type translation key to search:")
	translationKey := s.ReadLine()
	if len(translationKey) == 0 {
		return details, errors.New("No translation key provided")
	}
	return details, nil
}
