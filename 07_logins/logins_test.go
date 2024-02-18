package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var ltestOk1 = `4
hello
ozoner
roma
anykey
7
roma
ello
zooner
ankyey
ynakey
amor
rom
`

var ltestOkResult1 = `1
0
1
1
0
0
0
`

func TestLogins(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(ltestOk1))
	out := new(bytes.Buffer)
	err := logins(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != ltestOkResult1 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, ltestOkResult1)
	}
}

var ltestOk2 = `4
hello
ozoner
roma
anykey
1
ynakey
`

var ltestOkResult2 = `0
`

func TestLogin2(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(ltestOk2))
	out := new(bytes.Buffer)
	err := logins(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != ltestOkResult2 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, ltestOkResult2)
	}
}
