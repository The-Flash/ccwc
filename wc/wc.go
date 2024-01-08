package wc

import (
	"bufio"
	"io"
	"strings"
)

type WC struct {
	r io.ReadSeeker
}

func NewWC(r io.ReadSeeker) WC {
	return WC{r: r}
}

func (w *WC) ByteCount() int {
	count := 0
	var p []byte = make([]byte, 4096)
	w.r.Seek(0, 0)
	reader := bufio.NewReader(w.r)

	for {
		n, err := reader.Read(p)
		if err != nil {
			break
		}
		count = count + n
	}

	return count
}

func (w *WC) LineCount() int {
	count := 0
	w.r.Seek(0, 0)
	reader := bufio.NewReader(w.r)
	for {
		_, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		count = count + 1
	}
	return count
}

func (w *WC) WordCount() int {
	count := 0
	w.r.Seek(0, 0)
	scanner := bufio.NewScanner(w.r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count = count + 1
	}
	return count
}

func (w *WC) CharCount() int {
	count := 0
	w.r.Seek(0, 0)
	reader := bufio.NewReader(w.r)
	for {
		_, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		count = count + 1
	}
	return count
}

func (w *WC) CountAll() (int, int, int, int) {
	charCount := 0
	wordCount := 0
	byteCount := 0
	lineCount := 0
	scanner := bufio.NewScanner(w.r)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		wordCount += countWords(line)
		byteCount += countBytes(line)
		charCount += countChars(line)
	}
	return charCount, wordCount, byteCount, lineCount
}

func countWords(line string) int {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}
	return wordCount
}

func countBytes(line string) int {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanBytes)
	byteCount := 0
	for scanner.Scan() {
		byteCount++
	}
	return byteCount
}

func countChars(line string) int {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanRunes)
	charCount := 0
	for scanner.Scan() {
		charCount++
	}
	return charCount
}
