package main

import "os"

type demoFunc func()

var parts map[string]demoFunc

func main() {
	parts = map[string]demoFunc{
		"str":      StrDemo,
		"encoding": EncodingDemo,
		"func":     FuncDemo,
	}

	if len(os.Args) < 2 {
		runAll()
		return
	}

	partName := os.Args[1]
	targetFunc := parts[partName]
	if targetFunc != nil {
		targetFunc()
	}
}

func runAll() {
	for name, targetFunc := range parts {
		pf("Demo of %s:", name)
		targetFunc()
	}
}
