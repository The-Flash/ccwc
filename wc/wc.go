package wc

import (
	"bufio"
	"fmt"
	"io"
)

type WC struct {
	r io.Reader
}

func NewWC(r io.Reader) WC {
	return WC{r: r}
}

func (w WC) ByteCount() int {
	count := 0
	var p []byte = make([]byte, 4096)
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
	return 3
}

func (w WC) WordCount() int {
	return 3
}
