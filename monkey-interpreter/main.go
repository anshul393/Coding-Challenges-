package main

import (
	"fmt"

	"github.com/anshul393/coding-challenges/monkey-interpreter/token"
)

func main() {
	token := token.Token{Literal: "123", Type: token.ASSIGN}
	fmt.Printf("%s,%d\n", token.Literal, token.Type)
}
