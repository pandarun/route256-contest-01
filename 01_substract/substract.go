package main

import (
	"bufio"
	"fmt"
	"io"
)

func substract(input io.Reader, output io.Writer) error {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(input)
	out = bufio.NewWriter(output)
	defer out.Flush()

	var a, b int
	fmt.Fscan(in, &a, &b)
	fmt.Fprint(out, a-b)

	return nil
}
