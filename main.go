package main

import (
	"log"

	"github.com/abiosoft/ishell"
	"github.com/sqars/managetranslations/actions"
	"github.com/sqars/managetranslations/config"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	shell := ishell.New()

	actions := []actions.Action{
		actions.NewAddTranslation(config),
		actions.NewDeleteTranslation(config),
		actions.NewUpdateTranslationFromExisting(config),
		actions.NewUpdateFromCSV(config),
	}

	options := []string{}
	for _, action := range actions {
		options = append(options, action.GetName())
	}
	optionSelected := shell.MultiChoice(options, "Choose an option")

	err = actions[optionSelected].PromptActionDetails(shell)
	if err != nil {
		log.Fatalln(err)
	}
}
