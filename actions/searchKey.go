package actions

// NewSearchKey creates instance of SearchKey struct
func NewSearchKey() *SearchKey {
	return &SearchKey{
		name: "Search translation key",
	}
}

// SearchKey struct of operation of adding translation key
type SearchKey struct {
	name string
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
	details.translationKey = searchKey
	return details, nil
}
