package main

import "testing"

func Test_applyMask(t *testing.T) {
	type args struct {
		mask  string
		value uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{ 
			"11 to 73",
			args{
				mask:  "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
				value: 11,
			},
			73,
		},
		{
			"101 to 101",
			args{
				mask:  "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
				value: 101,
			},
			101,
		},
		{
			"0 to 64",
			args{
				mask:  "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
				value: 0,
			},
			64,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyMask(tt.args.mask, tt.args.value); got != tt.want {
				t.Errorf("applyMask() = %v, want %v", got, tt.want)
			}
		})
	}
}
