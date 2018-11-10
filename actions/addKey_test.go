package actions

import (
	"reflect"
	"testing"
)

func TestAddTranslation_addKey(t *testing.T) {
	type fields struct {
		name string
	}
	type args struct {
		data Translation
		d    ActionDetails
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Translation
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
				d: ActionDetails{
					translationKey: "toAdd",
				},
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
			a := &AddTranslation{
				name: tt.fields.name,
			}
			if got := a.addKey(tt.args.data, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTranslation.addKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
