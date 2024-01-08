package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/The-Flash/ccwc/wc"
)

func main() {
	wcPtr := flag.Bool("w", false, "Count words")
	ccPtr := flag.Bool("m", false, "Count characters")
	lcPtr := flag.Bool("l", false, "Count lines")
	bcPtr := flag.Bool("c", false, "Count bytes")
	flag.Parse()

	var f *os.File

	filepath := flag.Arg(0)
	f, err := os.Open(filepath)
	if err != nil {
		f = os.Stdin
	}
	if f == nil {
		panic("could not open file")
	}
	defer f.Close()
	libWc := wc.NewWC(f)
	if *wcPtr {
		formatOutput(-1, -1, libWc.WordCount(), -1, filepath)
		return
	}
	if *ccPtr {
		formatOutput(-1, -1, -1, libWc.CharCount(), filepath)
		return
	}
	if *lcPtr {
		formatOutput(-1, libWc.LineCount(), -1, -1, filepath)
		return
	}
	if *bcPtr {
		formatOutput(libWc.ByteCount(), -1, -1, -1, filepath)
		return
	}

	if !(*wcPtr || *ccPtr || *lcPtr || *bcPtr) {
		formatOutput(libWc.ByteCount(), libWc.LineCount(), libWc.WordCount(), -1, filepath)
		return
	}

}

func formatOutput(byteCount int, lineCount int, wordCount int, charCount int, filepath string) {
	output := ""
	if byteCount > -1 {
		output += fmt.Sprint(byteCount) + " "
	}
	if lineCount > -1 {
		output += fmt.Sprint(lineCount) + " "
	}
	if wordCount > -1 {
		output += fmt.Sprint(wordCount) + " "
	}
	if charCount > -1 {
		output += fmt.Sprint(charCount) + " "
	}
	output += " " + filepath
	fmt.Println(output)
}
