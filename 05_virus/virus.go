package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Directory struct {
	Dir     string
	Files   []string
	Folders []Directory
}

var buffer = new(bytes.Buffer)

func virus(input io.Reader, output io.Writer) error {

	buffer.Grow(10 * 1024 * 1024)

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(input)
	scanner := bufio.NewScanner(in)
	out = bufio.NewWriter(output)
	defer out.Flush()

	var dataInputs int
	scanner.Scan()
	dataInputsString := scanner.Text()

	dataInputs, err := parseNumber11(dataInputsString)
	if err != nil {
		return fmt.Errorf("unable to parse data inputs")
	}

	var i = 0
	for i < dataInputs {

		scanner.Scan()
		jsonLinesString := scanner.Text()

		jsonLines, err := parseNumber11(jsonLinesString)
		if err != nil {
			return fmt.Errorf("unable to parse json lines")
		}
		var dir Directory

		j := 0
		for j < jsonLines {
			scanner.Scan()
			line := scanner.Text()

			buffer.WriteString(line)
			j++
		}

		err = json.NewDecoder(buffer).Decode(&dir)
		if err != nil {
			return fmt.Errorf("unable to read directory")
		}

		poisoned := 0
		process(&dir, &poisoned, false)

		fmt.Fprintln(output, poisoned)

		i++
	}

	return nil
}

func process(d *Directory, poisoned *int, parentPoisoned bool) {

	if parentPoisoned {
		*poisoned = *poisoned + len(d.Files)
		for _, f := range d.Folders {
			process(&f, poisoned, true)
		}
	} else {
		containsVirus := false
		for _, f := range d.Files {
			if strings.LastIndex(f, ".hack") != -1 {
				containsVirus = true
				break
			}
		}

		if containsVirus {
			*poisoned = *poisoned + len(d.Files)
		}

		for _, f := range d.Folders {
			process(&f, poisoned, containsVirus)
		}
	}

}

func parseNumber11(item string) (int, error) {
	number, err := strconv.Atoi(item)
	if err != nil {
		fmt.Println("unable to parse int")
		return 0, fmt.Errorf("unable to parse int")
	}
	return number, nil
}
