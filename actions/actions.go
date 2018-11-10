package actions

// Action represents action to perform on json file
type Action interface {
	GetName() string
	PromptActionDetails(promptShell, dataCollector) (ActionDetails, error)
	PerformAction(ActionDetails) error
}

// ActionDetails represents details for action to perform
type ActionDetails struct {
	selectedFilesPaths []string
	translations       Translation
	translationKey     string
}
