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

	runTarget(os.Args[1])
}

func runAll() {
	for name := range parts {
		runTarget(name)
	}
}

func runTarget(name string) {
	targetFunc := parts[name]
	pf("\n------ Demo of %s: -------", name)
	if targetFunc != nil {
		targetFunc()
	}
}
