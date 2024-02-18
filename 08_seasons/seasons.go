package main

import (
	"bufio"
	"fmt"
	"io"
)

func seasons(input io.Reader, output io.Writer) error {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(input)
	out = bufio.NewWriter(output)
	defer out.Flush()

	var dataInputs int
	fmt.Fscan(in, &dataInputs)

	var i = 0
	for i < dataInputs {

		var n int
		fmt.Fscan(in, &n)

		prices := make([]int, n, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &prices[j])
		}

		fmt.Fprintln(out, prices)

		var seasonNumber int
		result := make([]int, n, n)
		var seasonStart, seasonEnd int

		highSeason := false
		lowSeason := false

		inSeason := false
		for j := 1; j < n; j++ {
			seasonNumber = 0
			inSeason := highSeason || lowSeason
			if prices[j] > prices[j-1] && !inSeason {
				seasonNumber++
				highSeason = true
				seasonStart = j
			} else if prices[j] < prices[j-1] && inSeason {
				lowSeason = true
				seasonEnd = j
			} else if j == n-1 && inSeason {
				seasonNumber++
				lowSeason = false
				highSeason = false

				seasonEnd = j
			}

			if inSeason {
			}

			prevSeason := inSeason
		}

		fmt.Fprintln(out, result)

		i++
	}

	return nil
}

type Seasons struct {
	seasons []int
}
