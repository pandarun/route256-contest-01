package main

import (
	"os"

	quoteV3 "rsc.io/quote/v3"
)

func main() {
	err := marks(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}

func Hello() string {
	return quoteV3.HelloV3()
}

func Proverb() string {
	return quoteV3.Concurrency()
}
