package ufap

import (
	"testing"
)

func TestHashTable(t *testing.T) {
	h := NewHashTable()
	h.Put(1, "one")
	h.Put(2, "two")
	h.Put(3, "three")
	h.Put(10465270, "10465270")

	v1, ok := h.Find(1)
	if !ok {
		t.Fatalf("fail")
	}

	if v1 != "one" {
		t.Fatalf("fail")
	}

	v2, ok := h.Find(2)
	if !ok {
		t.Fatalf("fail")
	}

	if v2 != "two" {
		t.Fatalf("fail")
	}

	v3, ok := h.Find(3)
	if !ok {
		t.Fatalf("fail")
	}

	if v3 != "three" {
		t.Fatalf("fail")
	}

	v4, ok := h.Find(10465270)
	if !ok {
		t.Fatalf("fail")
	}

	if v4 != "10465270" {
		t.Fatalf("fail")
	}

}
