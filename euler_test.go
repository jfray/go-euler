package main

import (
	"testing"
)

func BenchmarkProjectOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectOne(1000)
	}
}

func TestProjectOne(t *testing.T) {
	expected := 233168
	got := ProjectOne(1000)
	if expected != got {
		t.Errorf("got %d expected %d", got, expected)
	}
}
