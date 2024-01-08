package wc

import (
	"bufio"
	"io"
)

type WC struct {
	r io.ReadSeeker
}

func NewWC(r io.ReadSeeker) WC {
	return WC{r: r}
}

func (w WC) ByteCount() int {
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

func (w WC) LineCount() int {
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

func (w WC) WordCount() int {
	count := 0
	w.r.Seek(0, 0)
	scanner := bufio.NewScanner(w.r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count = count + 1
	}
	return count
}

func (w WC) CharCount() int {
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
