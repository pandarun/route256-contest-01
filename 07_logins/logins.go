package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func logins(input io.Reader, output io.Writer) error {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(input)
	out = bufio.NewWriter(output)
	defer out.Flush()

	var employeeCount int
	fmt.Fscan(in, &employeeCount)

	var employees = make(map[string]int)
	var employeeSet = make(map[string][]string)

	for j := 0; j < employeeCount; j++ {
		var employee string
		fmt.Fscan(in, &employee)
		employees[employee] = 1

		sorted := getSortedLoginKey(employee)

		employeeSet[sorted] = append(employeeSet[sorted], employee)
	}

	var loginCount int
	fmt.Fscan(in, &loginCount)
	for j := 0; j < loginCount; j++ {
		var login string
		fmt.Fscan(in, &login)

		if _, ok := employees[login]; ok {
			fmt.Fprintln(out, "1")
		} else if candidate, ok := employeeSet[getSortedLoginKey(login)]; ok {

			if isSimilar(candidate, login) {
				fmt.Fprintln(out, "1")
			} else {
				fmt.Fprintln(out, "0")
			}

		} else {
			fmt.Fprintln(out, "0")
		}
	}

	return nil
}

func isSimilar(candidates []string, login string) bool {
	for _, c := range candidates {
		if OneAdjacentPermutationDistance(c, login) {
			return true
		}
	}
	return false
}

func OneAdjacentPermutationDistance(c string, login string) bool {

	if len(c) != len(login) {
		return false
	}

	var diffPositions []int
	var diffCount = 0
	for i := 0; i < len(c); i++ {
		if c[i] != login[i] {
			diffCount++
			diffPositions = append(diffPositions, i)
		}

		if diffCount > 2 {
			return false
		}
	}

	return diffCount == 2 && diffPositions[0]+1 == diffPositions[1]
}

func getSortedLoginKey(employee string) string {
	sorted := []rune(employee)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	return string(sorted)
}
