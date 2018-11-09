package utils

import (
	"reflect"
	"testing"
)

func Test_mergeTranslations(t *testing.T) {
	type args struct {
		translations []Translation
	}
	tests := []struct {
		name string
		args args
		want Translation
	}{
		{
			name: "Test case 1",
			args: args{
				translations: []Translation{
					Translation{
						"en": map[string]string{
							"dataval": "en_dataval",
							"missing": "en_missing",
						},
						"pl": map[string]string{
							"dataval": "pl_dataval",
							"missing": "pl_missing",
						},
					},
					Translation{
						"en": map[string]string{
							"dataval2": "en_dataval2",
							"missing":  "",
						},
						"pl": map[string]string{
							"dataval2": "pl_dataval2",
							"missing":  "",
						},
					},
				},
			},
			want: Translation{
				"en": map[string]string{
					"dataval":  "en_dataval",
					"dataval2": "en_dataval2",
					"missing":  "en_missing",
				},
				"pl": map[string]string{
					"dataval":  "pl_dataval",
					"dataval2": "pl_dataval2",
					"missing":  "pl_missing",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTranslations(tt.args.translations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTranslations() = %v, want %v", got, tt.want)
			}
		})
	}
}
