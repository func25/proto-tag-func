package main

import "flag"

//protoc --go_out=. ./test/taginject/*.proto

func main() {
	var inp string
	flag.StringVar(&inp, "input", "./proto/auditproto/*.go", "input file path's pattern")
	flag.Parse()

	Parse(inp)
}
