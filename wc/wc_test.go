package wc

import (
	"os"
	"testing"
)

func TestByteCount(t *testing.T) {
	r, err := os.Open("../test.txt")
	if err != nil {
		t.Fatal(err)
	}

	wc := NewWC(r)
	byteCount := wc.ByteCount()
	expectedCount := 342190
	if byteCount != expectedCount {
		t.Fatalf("got %d, want %d", byteCount, expectedCount)
	}
}

func TestLineCount(t *testing.T) {
	r, err := os.Open("../test.txt")
	if err != nil {
		t.Fatal(err)
	}

	wc := NewWC(r)
	lineCount := wc.LineCount()
	expectedCount := 7145
	if lineCount != expectedCount {
		t.Fatalf("got %d, want %d", lineCount, expectedCount)
	}
}

func TestWordCount(t *testing.T) {
	r, err := os.Open("../test.txt")
	if err != nil {
		t.Fatal(err)
	}

	wc := NewWC(r)
	wordCount := wc.WordCount()
	expectedCount := 58164
	if wordCount != expectedCount {
		t.Fatalf("got %d, want %d", wordCount, expectedCount)
	}
}
