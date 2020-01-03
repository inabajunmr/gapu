package ufap

import (
	"testing"
)

func TestStackSingleValue(t *testing.T) {
	s := Stack{}
	s.Push("test")
	val, ok := s.Pop().(string)
	if !ok {
		t.Fatalf("failed test for pop()")
	}
	if val != "test" {
		t.Fatal("failed test")
	}
}

func TestStackMultipleValue(t *testing.T) {
	s := Stack{}
	if !s.IsEmpty() {
		t.Fatalf("failed test for IsEmpty()")
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Fatalf("failed test for IsEmpty()")
	}

	s.Push(2)
	if s.IsEmpty() {
		t.Fatalf("failed test for IsEmpty()")
	}

	s.Push(3)
	if s.IsEmpty() {
		t.Fatalf("failed test for IsEmpty()")
	}

	val, ok := s.Pop().(int)
	if !ok {
		t.Fatalf("failed test for pop()")
	}
	if val != 3 {
		t.Fatal("failed test")
	}
	val, ok = s.Pop().(int)
	if !ok {
		t.Fatalf("failed test for pop()")
	}
	if val != 2 {
		t.Fatal("failed test")
	}
	val, ok = s.Pop().(int)
	if !ok {
		t.Fatalf("failed test for pop()")
	}
	if val != 1 {
		t.Fatal("failed test")
	}

	if !s.IsEmpty() {
		t.Fatalf("failed test for IsEmpty()")
	}
}

func TestStackPushPopPush(t *testing.T) {
	s := Stack{}
	s.Push(1)
	val, ok := s.Pop().(int)
	if !ok {
		t.Fatalf("failed test for pop()")
	}
	if val != 1 {
		t.Fatal("failed test")
	}

	s.Push(2)

	val, ok = s.Pop().(int)
	if !ok {
		t.Fatalf("failed test for pop()")
	}
	if val != 2 {
		t.Fatal("failed test")
	}
}
