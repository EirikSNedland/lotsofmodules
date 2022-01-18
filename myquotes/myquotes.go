package myquotes

import (
	"fmt"
	"rsc.io/quote"
)

func Display() string {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Opt())
	return quote.Opt()
}
