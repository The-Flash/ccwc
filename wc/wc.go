package wc

import (
	"bufio"
	"bytes"
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
	scanner.Split(scanLines)
	for scanner.Scan() {
		line := scanner.Text() + "\n"
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

// scanLines is a split function for a Scanner that returns each line of
// text, stripped of any trailing end-of-line marker. The returned line may
// be empty. The end-of-line marker is one optional carriage return followed
// by one mandatory newline. In regular expression notation, it is `\n`.
// The last non-empty line of input will be returned even if it has no
// newline.
func scanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
