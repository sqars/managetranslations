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

func TestAddkey(t *testing.T) {
	type args struct {
		data Translation
		key  string
	}
	tests := []struct {
		name string
		args args
		want Translation
	}{
		{
			name: "Test Case 1",
			args: args{
				data: Translation{
					"en": map[string]string{
						"existing": "en_existing",
					},
					"pl": map[string]string{
						"existing": "pl_existing",
					},
				},
				key: "toAdd",
			},
			want: Translation{
				"en": map[string]string{
					"existing": "en_existing",
					"toAdd":    "",
				},
				"pl": map[string]string{
					"existing": "pl_existing",
					"toAdd":    "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Addkey(tt.args.data, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Addkey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveKey(t *testing.T) {
	type args struct {
		data Translation
		key  string
	}
	tests := []struct {
		name string
		args args
		want Translation
	}{
		{
			name: "Test Case 1",
			args: args{
				data: Translation{
					"en": map[string]string{
						"existing": "en_existing",
					},
					"pl": map[string]string{
						"existing": "pl_existing",
					},
				},
				key: "existing",
			},
			want: Translation{
				"en": map[string]string{},
				"pl": map[string]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveKey(tt.args.data, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateTranslations(t *testing.T) {
	type args struct {
		data Translation
		pool Translation
	}
	tests := []struct {
		name string
		args args
		want Translation
	}{
		{
			name: "test case 1 update from existing",
			args: args{
				data: Translation{
					"en": map[string]string{
						"dataval": "en_dataval",
						"missing": "",
					},
					"pl": map[string]string{
						"dataval": "pl_dataval",
						"missing": "",
					},
				},
				pool: Translation{
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
			want: Translation{
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
				data: Translation{
					"en": map[string]string{
						"dataval": "en_dataval",
						"missing": "",
					},
					"pl": map[string]string{
						"dataval": "pl_dataval",
						"missing": "",
					},
				},
				pool: Translation{
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
			want: Translation{
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
			if got := UpdateTranslations(tt.args.data, tt.args.pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateTranslations() = %v, want %v", got, tt.want)
			}
		})
	}
}
