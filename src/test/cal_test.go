package main

import "testing"

func Test_addUpper(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3}, 6},
		{"2", args{4}, 10},
		{"3", args{5}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addUpper(tt.args.n); got != tt.want {
				t.Errorf("addUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}
