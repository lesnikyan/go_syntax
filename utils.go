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
