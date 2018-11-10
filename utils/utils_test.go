package utils

import (
	"reflect"
	"testing"
)

func TestIndexOfInt(t *testing.T) {
	type args struct {
		in0 []int
		in1 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				in0: []int{1, 2, 3, 4, 5},
				in1: 4,
			},
			want: 3,
		}, {
			name: "Test 2",
			args: args{
				in0: []int{1, 2, 2, 2, 5},
				in1: 2,
			},
			want: 1,
		}, {
			name: "Test 3",
			args: args{
				in0: []int{1, 2, 2, 2, 5},
				in1: 6,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOfInt(tt.args.in0, tt.args.in1); got != tt.want {
				t.Errorf("IndexOfInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt(t *testing.T) {
	type args struct {
		s []int
		v int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				s: []int{1, 2, 2, 2, 5},
				v: 2,
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				s: []int{1, 2, 2, 2, 5},
				v: 7,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt(tt.args.s, tt.args.v); got != tt.want {
				t.Errorf("ContainsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupByIdxValue(t *testing.T) {
	type args struct {
		d   [][]string
		idx int
	}
	tests := []struct {
		name string
		args args
		want map[string][][]string
	}{
		{
			name: "Test 1",
			args: args{
				d: [][]string{
					[]string{"val1", "toGroup1"},
					[]string{"val2", "toGroup1"},
					[]string{"val3", "toGroup1"},
					[]string{"val4", "toGroup1"},
					[]string{"val5", "toGroup2"},
					[]string{"val2", "toGroup2"},
					[]string{"val6", "toGroup3"},
				},
				idx: 1,
			},
			want: map[string][][]string{
				"toGroup1": [][]string{
					[]string{"val1", "toGroup1"},
					[]string{"val2", "toGroup1"},
					[]string{"val3", "toGroup1"},
					[]string{"val4", "toGroup1"},
				},
				"toGroup2": [][]string{
					[]string{"val5", "toGroup2"},
					[]string{"val2", "toGroup2"},
				},
				"toGroup3": [][]string{
					[]string{"val6", "toGroup3"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupByIdxValue(tt.args.d, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupByIdxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
