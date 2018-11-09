package actions

import (
	"errors"
	"reflect"
	"testing"
)

type filesCollectorMock struct {
	jsonConfig          string
	selectedFilesResult []string
}

func (f filesCollectorMock) PromptFiles(shell PromptShell, msg string, conf string) ([]string, error) {
	if msg != "Select file(s) to add translation key" || conf != "mockJSONConfig" {
		return nil, errors.New("Arguments not match")
	}
	return f.selectedFilesResult, nil
}

func (f filesCollectorMock) getJSONConfig() string {
	return f.jsonConfig
}

func (f filesCollectorMock) getCSVConfig() string {
	return f.jsonConfig
}

type shellMock struct {
	readLineResult  string
	checklistResult []int
}

func (s shellMock) Println(...interface{}) {}
func (s shellMock) Checklist(opts []string, msg string, init []int) []int {
	return s.checklistResult
}
func (s shellMock) ReadLine() string {
	return s.readLineResult
}

func TestAddTranslation_PromptActionDetails(t *testing.T) {
	type args struct {
		s PromptShell
		d filesCollector
	}
	tests := []struct {
		name    string
		args    args
		want    ActionDetails
		wantErr bool
	}{
		{
			name: "Should pass without errors",
			args: args{
				d: filesCollectorMock{
					jsonConfig:          "mockJSONConfig",
					selectedFilesResult: []string{"filePath1", "filePath2"},
				},
				s: shellMock{
					readLineResult: "test",
				},
			},
			want: ActionDetails{
				selectedFilesPaths: []string{"filePath1", "filePath2"},
				translationKey:     "test",
			},
			wantErr: false,
		}, {
			name: "Should throw error when PromptFiles fails",
			args: args{
				d: filesCollectorMock{
					jsonConfig:          "failconfig",
					selectedFilesResult: []string{"filePath1", "filePath2"},
				},
				s: shellMock{
					readLineResult: "test",
				},
			},
			wantErr: true,
		}, {
			name: "Should throw error when no input specified",
			args: args{
				d: filesCollectorMock{
					jsonConfig:          "mockJSONConfig",
					selectedFilesResult: []string{"filePath1", "filePath2"},
				},
				s: shellMock{
					readLineResult: "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AddTranslation{}
			got, err := a.PromptActionDetails(tt.args.s, tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTranslation.PromptActionDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTranslation.PromptActionDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
