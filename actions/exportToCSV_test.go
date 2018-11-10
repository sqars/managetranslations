package actions

import (
	"reflect"
	"testing"
)

func Test_transformToCSVformat(t *testing.T) {
	type args struct {
		vals map[string]string
		path string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test 1",
			args: args{
				vals: map[string]string{
					"key":  "value",
					"key2": "value2",
				},
				path: "some path",
			},
			want: [][]string{
				[]string{"key", "value", "some path"},
				[]string{"key2", "value2", "some path"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transformToCSVformat(tt.args.vals, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transformToCSVformat() = %v, want %v", got, tt.want)
			}
		})
	}
}
