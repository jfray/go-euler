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

func BenchmarkProjectTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectTwo(4000000)
	}
}

func TestProjectTwo(t *testing.T) {
	expected := 4613732
	got := ProjectTwo(4000000)
	if expected != got {
		t.Errorf("got %d expected %d", got, expected)
	}
}

func BenchmarkProjectThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectThree(6857)
	}
}

func TestProjectThree(t *testing.T) {
	expected := 6857
	got := ProjectThree(600851475143)
	if expected != got {
		t.Errorf("got %d expected %d", got, expected)
	}
}

func BenchmarkProjectFour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectFour()
	}
}

func TestProjectFour(t *testing.T) {
	expected := 906609
	got := ProjectFour()
	if expected != got {
		t.Errorf("got %d expected %d", got, expected)
	}
}

func BenchmarkProjectFive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectFive(1, 20)
	}
}

func TestProjectFive(t *testing.T) {
	expected := 232792560
	got := ProjectFive(1, 20)
	if expected != got {
		t.Errorf("got %d expected %d", got, expected)
	}
}
