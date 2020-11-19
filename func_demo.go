package main

import (
	"fmt"
	"sort"
)

// FuncDemo provides set of cases of define and use functions
// with different formats
func FuncDemo() {

	// arg types
	fun1 := func(x int64, y float32, boolMark bool, nums []int, name string, options map[string]int, abc Abc) {
		pf("(%d, %.3f, %v), nums: %v, name: %s, Abc(%d, %s, %c)", x, y, boolMark, nums, name, abc.a, abc.b, abc.c)
		for k, v := range options {
			pf("opt %s: %d", k, v)
		}
	}

	abc1 := Abc{
		a: 200,
		b: "ABC b-val",
		c: rune('Ъ'),
	}

	opts := map[string]int{
		"opt-a": 1111,
		"opt-b": 2222,
		"opt-c": 3333,
	}

	fun1(100, 111.125009, true, []int{11, 22, 33}, "Big Vasya", opts, abc1)

	// uncount args

	// func result type

	// several result values
	numData := []float64{1, 2, 3, 4, 5}
	min, max, average := minMaxAvrg(numData)
	pf("min = %.3f, max = %.3f, average = %.3f", min, max, average)

	// struct and method
	abc2 := Abc{
		a: 10,
		b: "test ABC",
		c: rune('Ж'),
	}
	p(abc2.info())

	// Closure 1 lvl
	demoClosure()

	// panic, recover, defer
	panicDemo()
}

//panic
func panicDemo() {
	panicHandler := func() {
		msg := recover() // return nil if no panic
		pf("panic->recover: %v", msg)
	}
	defer panicHandler()

	p("Panic demo: before panic")
	panic("Panik demo: PANIC!!!")
	p("Panic demo: after panik") // unreachable
}

// min, max, average = minMaxAvrg(data)
func minMaxAvrg(data []float64) (float64, float64, float64) {
	inData := data[:]
	sort.Float64s(inData)
	sum := 0.0
	for _, n := range inData {
		sum += n
	}
	return inData[0], inData[len(inData)-1], sum / float64(len(inData))
}

// Abc simple type
type Abc struct {
	a int
	b string
	c rune
}

func (v *Abc) info() string {
	return fmt.Sprintf("Abc(%d, %s, %x)", v.a, v.b, v.c)
}

// closure
func demoClosure() {
	f3 := outer(3)
	f5 := outer(5)
	p("f3: ", f3(3), f3(4), f3(5))
	p("f5", f5(6), f5(7), f5(8))
}

func outer(x int) func(int) int {
	z := 0
	// func1
	foo := func() int {
		z++
		return z
	}
	// func2
	inner := func(y int) int {
		return x*100 + foo()*10 + y
	}
	return inner
}
