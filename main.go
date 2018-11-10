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

	appActions := []actions.Action{
		actions.NewSearchKey(),
		actions.NewAddTranslation(),
		actions.NewDeleteTranslation(),
		actions.NewUpdateTranslationFromExisting(),
		actions.NewUpdateFromCSV(),
	}

	dataCollector := actions.NewDataCollector(config)

	options := []string{}
	for _, action := range appActions {
		options = append(options, action.GetName())
	}
	optionSelected := shell.MultiChoice(options, "Choose an option")

	selectedAction := appActions[optionSelected]

	actionDetails, err := selectedAction.PromptActionDetails(shell, dataCollector)
	if err != nil {
		log.Fatalln(err)
	}
	err = actionDetails.PerformAction(selectedAction.GetModifierFn())
	if err != nil {
		log.Fatalln(err)
	}
}
