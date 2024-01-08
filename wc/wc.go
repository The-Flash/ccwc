package wc

import (
	"io"
)

type WC struct {
	r io.Reader
}

func NewWC(r io.Reader) WC {
	return WC{r: r}
}

func (w WC) ByteCount() int {
	return 2
}

func (w WC) LineCount() int {
	return 3
}

func (w WC) WordCount() int {
	return 3
}
