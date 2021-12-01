package main

import (
	"flag"
	"fmt"
)

//protoc --go_out=. ./test/taginject/*.proto

func main() {
	var inp string
	flag.StringVar(&inp, "input", "./proto/auditproto/*.go", "input file path's pattern")
	flag.Parse()
	fmt.Println(inp)
	Parse(inp)
}
