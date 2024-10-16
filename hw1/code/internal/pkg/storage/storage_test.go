package storage

import (
	"testing"
)

type TestCase struct {
	key   string
	value string
	kind  string
}

func TestSetAndGet(t *testing.T) {
	cases := []TestCase{
		{"avocado", "super", "S"},
		{"banana", "123", "D"},
		{"dima", "people", "S"},
		{"ilya", "19", "D"},
		{"money", "inf", "S"},
		{"love", "none", "S"},
	}

	stor, err := NewStorage()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, c := range cases {
		stor.Set(c.key, c.value)
		got := stor.Get(c.key)
		if got == nil || *got != c.value {
			t.Errorf("Get(%q) == %q, want %q", c.key, *got, c.value)
		}
		gotKind := stor.GetKind(c.key)
		if gotKind != c.kind {
			t.Errorf("GetKind(%q) == %q, want %q", c.key, gotKind, c.kind)
		}
	}
}
