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
	want := []int{9, 7, 6, 7, 6, 5, 4, 3, 2, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkFoo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		terr := newTerrain(100, 100, testSeed)
		Foo(terr, 100)
	}
}
