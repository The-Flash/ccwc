package wc

import (
	"os"
	"testing"
)

func TestCountAll(t *testing.T) {
	r, err := os.Open("../test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	wc := NewWC(r)
	charCount, wordCount, byteCount, lineCount := wc.CountAll()
	expectedByteCount := 342190
	expectedCharCount := 339292
	expectedLineCount := 7145
	expectedWordCount := 58164
	if byteCount != expectedByteCount {
		t.Fatalf("byteCount: got %d, want %d", byteCount, expectedByteCount)
	}
	if charCount != expectedCharCount {
		t.Fatalf("charCount: got %d, want %d", charCount, expectedCharCount)
	}
	if lineCount != expectedLineCount {
		t.Fatalf("lineCount: got %d, want %d", lineCount, expectedLineCount)
	}
	if wordCount != expectedWordCount {
		t.Fatalf("wordCount: got %d, want %d", wordCount, expectedWordCount)
	}

}

func TestByteCount(t *testing.T) {
	r, err := os.Open("../test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

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

	defer r.Close()

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

	defer r.Close()

	wc := NewWC(r)
	wordCount := wc.WordCount()
	expectedCount := 58164
	if wordCount != expectedCount {
		t.Fatalf("got %d, want %d", wordCount, expectedCount)
	}
}

func TestCharacterCount(t *testing.T) {
	r, err := os.Open("../test.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer r.Close()

	wc := NewWC(r)
	charCount := wc.CharCount()
	expectedCount := 339292
	if charCount != expectedCount {
		t.Fatalf("got %d, want %d", charCount, expectedCount)
	}
}
