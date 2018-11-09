package actions

import (
	"github.com/abiosoft/ishell"
)

// Action represents action to perform on json file
type Action interface {
	GetName() string
	PromptActionDetails(*ishell.Shell) error
	// Perform(filePath, key, translation string) error
}
