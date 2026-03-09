package main

import (
	"fmt"
	"mransack/greetings"

	"golang.org/x/example/hello/reverse"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	message := greetings.Hello("Maass")
	fmt.Println(message)
	fmt.Println(reverse.String("Hello"))
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
