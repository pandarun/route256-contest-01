package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var tTestOk1 = `5
MRCMD
MDD
M
MDMRCMD
MMDD
`

var tTestOkResult1 = `YES
NO
NO
YES
NO
`

func TestTasks1(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(tTestOk1))
	out := new(bytes.Buffer)
	err := tasks(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != tTestOkResult1 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, tTestOkResult1)
	}
}
