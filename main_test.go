package main

import (
	"math/rand"
	"testing"
  "reflect"
)
var testSeed = rand.NewSource(125)

func TestFoo(t *testing.T) {
  line := Foo(testSeed, 10, 10, 10)

	got := line
	want := []int{10, 9, 8, 10, 10, 10, 8, 9, 7, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
