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

func Test_csvToTranslationFormat(t *testing.T) {
	type args struct {
		data [][]string
	}
	tests := []struct {
		name string
		args args
		want Translation
	}{
		{
			name: "Test case 1",
			args: args{
				data: [][]string{
					[]string{"", "en", "pl"},
					[]string{"CSV_tr_1", "en_tr_1", "pl_tr_1"},
					[]string{"CSV_tr_2", "en_tr_2", "pl_tr_2"},
				},
			},
			want: Translation{
				"en": map[string]string{
					"CSV_tr_1": "en_tr_1",
					"CSV_tr_2": "en_tr_2",
				},
				"pl": map[string]string{
					"CSV_tr_1": "pl_tr_1",
					"CSV_tr_2": "pl_tr_2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := csvToTranslationFormat(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("csvToTranslationFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
