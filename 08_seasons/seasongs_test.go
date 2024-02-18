package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var stestOk1 = `5
8
2 3 4 5 4 3 2 1
4
99 99 99 99
1
123456789
12
1 3 10 8 7 4 6 5 11 3 9 2
4
100 99 98 1000
`

var stestOkResult1 = `1 1 1 0 0 0 0 0
0 0 0 0
0
3 1 0 0 0 0 0 0 0 0 0 0
0 0 0 0
`

func TestSeasons(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(stestOk1))
	out := new(bytes.Buffer)
	err := seasons(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != ltestOkResult1 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, stestOkResult1)
	}
}

var stestOk2 = `1
8
2 3 4 5 4 3 2 1
`

var stestOkResult2 = `1 1 1 0 0 0 0 0
`

func TestSeasons2(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(stestOk2))
	out := new(bytes.Buffer)
	err := seasons(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != ltestOkResult2 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, stestOkResult2)
	}
}
