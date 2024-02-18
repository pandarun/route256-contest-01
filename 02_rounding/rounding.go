package main

import (
	"bufio"
	"fmt"
	"io"
)

func rounding(input io.Reader, output io.Writer) error {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(input)
	out = bufio.NewWriter(output)
	defer out.Flush()

	var dataInputs int
	_, err := fmt.Fscan(in, &dataInputs)
	if err != nil {
		return fmt.Errorf("unable to read data inputs")
	}

	var i = 0
	for i < dataInputs {

		var goodsSaleAmount int
		var commission float64

		_, err := fmt.Fscan(in, &goodsSaleAmount, &commission)

		totalCommissions := 0.0
		j := 0
		var goodsPrice float64

		for j < goodsSaleAmount {

			_, err := fmt.Fscan(in, &goodsPrice)
			if err != nil {
				return fmt.Errorf("unable to read goods price")
			}

			currentCommission := goodsPrice * commission / 100.0

			totalCommissions += currentCommission - float64(int(currentCommission))
			j++
		}

		_, err = fmt.Fprintf(output, "%.2f\n", totalCommissions)
		if err != nil {
			return fmt.Errorf("unable to write to output")
		}

		i++
	}

	return nil
}
