package gsa

import (
	"reflect"
	"testing"
)

func TestCountSort(t *testing.T) {
	type args struct {
		x string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test 1",
			args{"abaab"},
			"aaabb",
		},
		{
			"Test 2",
			args{"mississippi"},
			"iiiimppssss",
		},
		{
			"Test 3",
			args{""},
			"",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSort(tt.args.x); got != tt.want {
				t.Errorf("CountSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBucketSort(t *testing.T) {
	type args struct {
		x   string
		idx []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Test 1",
			args{"abaab", []int{0, 1, 2, 3, 4}},
			[]int{0, 2, 3, 1, 4},
		},
		{
			"Test 2",
			args{"abaab", []int{4, 3, 2, 1, 0}},
			[]int{3, 2, 0, 4, 1},
		},
		{
			"Test 3",
			args{"mississippi", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			[]int{1, 4, 7, 10, 0, 8, 9, 2, 3, 5, 6},
		},
		{
			"Test 4",
			args{"mississippi", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
			[]int{10, 7, 4, 1, 0, 9, 8, 6, 5, 3, 2},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BucketSort(tt.args.x, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BucketSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLsdRadixSort(t *testing.T) {
	type args struct {
		x string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LsdRadixSort(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LsdRadixSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsdRadixSort(t *testing.T) {
	type args struct {
		x string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MsdRadixSort(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MsdRadixSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
