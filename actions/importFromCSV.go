package actions

import (
	"github.com/sqars/managetranslations/config"
	"github.com/sqars/managetranslations/utils"
)

// NewEImportFromCSV creates instance of ExportToCSV struct
func NewEImportFromCSV(conf config.Config) *ImportFromCSV {
	return &ImportFromCSV{
		name: "Import translations from CSV(must be in proper format)",
		conf: conf,
	}
}

// ImportFromCSV struct of operation of importing translations from CSV
type ImportFromCSV struct {
	name string
	conf config.Config
}

// GetName returns name of Action
func (a *ImportFromCSV) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *ImportFromCSV) PromptActionDetails(s promptShell, d dataCollector) (ActionDetails, error) {
	details := ActionDetails{}
	selectedFilePaths, err := d.PromptFiles(
		s, "Select file(s) to import", a.conf.CSVFilePattern,
	)
	if err != nil {
		return details, err
	}
	details.selectedFilesPaths = selectedFilePaths
	return details, nil
}

// PerformAction Performs action using collected ActionDetails
func (a *ImportFromCSV) PerformAction(d ActionDetails) error {
	for _, filePath := range d.selectedFilesPaths {
		csvData, err := readCSVFile(filePath)
		if err != nil {
			return err
		}
		lang := csvData[0][1]
		// dont need first column
		csvData = csvData[1:]
		groupedByFilePath := utils.GroupByIdxValue(csvData, 2)
		for file, valuesToUpdate := range groupedByFilePath {
			data, err := GetJSONTranslationData(file)
			if err != nil {
				return err
			}
			for _, valueToUpdate := range valuesToUpdate {
				data[lang][valueToUpdate[0]] = valueToUpdate[1]
			}
			err = SaveJSONTranslationData(file, data)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
