package main

import (
	"math/rand"
	"reflect"
	"testing"
)

var testSeed = rand.NewSource(125)

func TestFoo(t *testing.T) {
	terr := newTerrain(10, 10, testSeed)
	Foo(terr, 10)

	got := terr.toSlice()
	want := []int{8, 7, 8, 9, 9, 8, 7, 6, 5, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
