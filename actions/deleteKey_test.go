package actions

import (
	"reflect"
	"testing"
)

func TestDeleteTranslation_removeKey(t *testing.T) {
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
					translationKey: "existing",
				},
			},
			want: Translation{
				"en": map[string]string{},
				"pl": map[string]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &DeleteTranslation{
				name: tt.fields.name,
			}
			if got := a.removeKey(tt.args.data, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteTranslation.removeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
