package main

import (
	"log"

	"github.com/abiosoft/ishell"
	"github.com/sqars/managetranslations/actions"
)

func main() {
	shell := ishell.New()

	actions := []actions.Action{
		actions.NewAddTranslation(),
		actions.NewDeleteTranslation(),
		actions.NewUpdateTranslationFromExisting(),
	}

	options := []string{}
	for _, action := range actions {
		options = append(options, action.GetName())
	}
	optionSelected := shell.MultiChoice(options, "Choose an option")

	err := actions[optionSelected].PromptActionDetails(shell)
	if err != nil {
		log.Fatalln(err)
	}
}
