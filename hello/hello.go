package main

import (
    "fmt";
 "rsc.io/quote";
 "mransack/greetings"
)
func main(){
    fmt.Println(quote.Go())
    message := greetings.Hello("Maass")
    fmt.Println(message)
}