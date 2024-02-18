package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var rTestOk1 = `2
5 5
.....
.#A#.
...B.
.#.#.
.....
7 9
.........
.#.#.#.#.
..AB.....
.#.#.#.#.
.........
.#.#.#.#.
.........
`

var rTestOkResult1 = `aaa..
.#A#.
...Bb
.#.#b
....b
aaa......
.#a#.#.#.
..ABb....
.#.#b#.#.
....b....
.#.#b#.#.
....bbbbb
`

func TestTRobots1(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(rTestOk1))
	out := new(bytes.Buffer)
	err := robots(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != rTestOkResult1 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, rTestOkResult1)
	}
}

var rTestOk2 = `1
3 3
B..
.#.
..A
`

var rTestOkResult2 = `B..
.#.
..A
`

func TestTRobots2(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(rTestOk2))
	out := new(bytes.Buffer)
	err := robots(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != rTestOkResult2 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, rTestOkResult2)
	}
}

var rTestOk4 = `1
5 5
.....
.#.#B
.....
.#.#.
...A.
`

var rTestOkResult4 = `bbbbb
.#.#B
.....
.#.#.
...Aa
`

func TestRobots4(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(rTestOk4))
	out := new(bytes.Buffer)
	err := robots(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != rTestOkResult4 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, rTestOkResult4)
	}

}

var rTestOk5 = `1
3 3
...
.#B
A..
`

var rTestOkResult5 = `bbb
.#B
Aaa
`

func TestRobots5(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(rTestOk5))
	out := new(bytes.Buffer)
	err := robots(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != rTestOkResult5 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, rTestOkResult5)
	}

}

var rTestOk6 = `1
3 3
...
.#.
.BA`

var rTestOkResult6 = `b..
b#.
bBA
`

func TestRobots6(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(rTestOk6))
	out := new(bytes.Buffer)
	err := robots(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != rTestOkResult6 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, rTestOkResult6)
	}
}
