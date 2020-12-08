package main

import (
	"reflect"
	"testing"
)

func Test_readBag(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *bag
	}{
		{
			name: "contains single other bag",
			args: args{"dark indigo bags contain 2 clear indigo bags."},
			want: &bag{
				name: "dark indigo",
				contains: map[string]int{
					"clear indigo": 2,
				},
			},
		},
		{
			name: "contains 2 other bags",
			args: args{"mirrored gold bags contain 2 pale blue bags, 1 dark violet bag."},
			want: &bag{
				name: "mirrored gold",
				contains: map[string]int{
					"pale blue":   2,
					"dark violet": 1,
				},
			},
		},
		{
			name: "contains no other bags",
			args: args{"shiny crimson bags contain no other bags."},
			want: &bag{
				name: "shiny crimson",
				contains: map[string]int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readBag(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readBag() = %v, want %v", got, tt.want)
			}
		})
	}
}
