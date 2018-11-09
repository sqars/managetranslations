package actions

import (
	"reflect"
	"testing"

	"github.com/sqars/managetranslations/utils"
)

func Test_updateTranslations(t *testing.T) {
	type args struct {
		data utils.Translation
		pool utils.Translation
	}
	tests := []struct {
		name string
		args args
		want utils.Translation
	}{
		{
			name: "test case 1 update from existing",
			args: args{
				data: utils.Translation{
					"en": map[string]string{
						"dataval": "en_dataval",
						"missing": "",
					},
					"pl": map[string]string{
						"dataval": "pl_dataval",
						"missing": "",
					},
				},
				pool: utils.Translation{
					"en": map[string]string{
						"val1":          "en_translation1",
						"missing":       "en_missing",
						"somethingelse": "en_somethingelse",
					},
					"pl": map[string]string{
						"val1":          "pl_translation1",
						"missing":       "pl_missing",
						"somethingelse": "pl_somethingelse",
					},
				},
			},
			want: utils.Translation{
				"en": map[string]string{
					"dataval": "en_dataval",
					"missing": "en_missing",
				},
				"pl": map[string]string{
					"dataval": "pl_dataval",
					"missing": "pl_missing",
				},
			},
		}, {
			name: "test case 2 when pool dont have translation",
			args: args{
				data: utils.Translation{
					"en": map[string]string{
						"dataval": "en_dataval",
						"missing": "",
					},
					"pl": map[string]string{
						"dataval": "pl_dataval",
						"missing": "",
					},
				},
				pool: utils.Translation{
					"en": map[string]string{
						"val1":          "en_translation1",
						"missing":       "",
						"somethingelse": "en_somethingelse",
					},
					"pl": map[string]string{
						"val1":          "pl_translation1",
						"missing":       "",
						"somethingelse": "pl_somethingelse",
					},
				},
			},
			want: utils.Translation{
				"en": map[string]string{
					"dataval": "en_dataval",
					"missing": "",
				},
				"pl": map[string]string{
					"dataval": "pl_dataval",
					"missing": "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateTranslations(tt.args.data, tt.args.pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateTranslations() = %v, want %v", got, tt.want)
			}
		})
	}
}
