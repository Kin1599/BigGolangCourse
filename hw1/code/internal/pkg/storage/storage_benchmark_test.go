package storage

import (
	"testing"
)

func BenchmarkSet(b *testing.B) {
	stor, err := NewStorage()
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stor.Set("key", "value")
	}
}

func BenchmarkGet(b *testing.B) {
	stor, err := NewStorage()
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}
	stor.Set("key", "value")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stor.Get("key")
	}
}

func BenchmarkSetAndGet(b *testing.B) {
	stor, err := NewStorage()
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stor.Set("key", "value")
		stor.Get("key")
	}
}
