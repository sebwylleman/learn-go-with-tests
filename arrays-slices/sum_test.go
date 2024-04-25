package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got %d, want %d, given %v", got, want, numbers)
	}

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	want := []int{6, 15}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}
