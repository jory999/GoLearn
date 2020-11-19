package hello

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

// Hello for this
func Hello() string {

	//return "Hello World" firm Git form vsCode

	return quote.Hello()
}

// Proverb For loop that iterates over sharks list and prints each string item
func Proverb() string {

	return quoteV3.Concurrency()
}
