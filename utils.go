package main

import (
	"fmt"
)

func p(a ...interface{}) {
	fmt.Println(a...)
}

func sf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func pf(format string, a ...interface{}) {
	p(sf(format, a...))
}

// ByteStrMap byte -> string
type ByteStrMap func(byte) string

func bmap(data []byte, foo ByteStrMap) []string {
	res := make([]string, len(data))
	for i, v := range data {
		res[i] = foo(v)
	}
	return res
}
