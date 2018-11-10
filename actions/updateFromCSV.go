package actions

// NewUpdateFromCSV creates instance of UpdateFromCSV struct
func NewUpdateFromCSV() *UpdateFromCSV {
	return &UpdateFromCSV{
		name: "Update empty translations from CSV file",
	}
}

// UpdateFromCSV struct of operation of updating empty translations from CSV file
type UpdateFromCSV struct {
	name string
}

// GetName returns name of Action
func (a *UpdateFromCSV) GetName() string {
	return a.name
}

// PromptActionDetails propmts for action details and runs Perform with arguments
func (a *UpdateFromCSV) PromptActionDetails(s promptShell, d dataCollector) (ActionDetails, error) {
	details := ActionDetails{}
	selectedTransationFilePaths, err := d.PromptFiles(
		s,
		"Select translation file(s) to update from CSV",
		d.getJSONConfig(),
	)
	if err != nil {
		return details, err
	}
	selectedCSVFilePaths, err := d.PromptFiles(
		s,
		"Select CSV file(s) with translations",
		d.getCSVConfig(),
	)
	if err != nil {
		return details, err
	}
	csvData, err := GetCSVTranslationData(selectedCSVFilePaths)
	if err != nil {
		return details, err
	}
	details.selectedFilesPaths = selectedTransationFilePaths
	details.translations = csvData
	return details, nil
}

// PerformAction Performs action using collected ActionDetails
func (a *UpdateFromCSV) PerformAction(d ActionDetails) error {
	err := modifyAndSaveTranslations(d, a.updateTranslations)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTranslations updates translation from pool(existing translations)
func (a *UpdateFromCSV) updateTranslations(data Translation, d ActionDetails) Translation {
	return UpdateTranslations(data, d.translations)
}
