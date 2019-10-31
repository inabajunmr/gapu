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
	s.Push(1)
	s.Push(2)
	s.Push(3)
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
