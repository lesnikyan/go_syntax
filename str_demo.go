package main

import "strings"

// StrDemo ...
func StrDemo() {
	// len
	p("Str Len:")
	p(len("Hello!"))
	p(len("1234567890"))
	p(len("абвгдежзиклмн"))

	// multibyte
	p("Multibyte strings:")
	p("len(абвгд) = ", len("абвгд"))
	// char by char
	// char code
	for i, c := range "абвгд" {
		p("> ", i, c, sf("U+%04x", c))
	}

	// starts with

	// repeate

	// replace

	// lower / upper

	// contains

	// join array

	// mapping runes

	// split
	p("Split:")
	p(strings.Split("123,44,5656,777,aaa,bc", ","))

	// substring
	p("Substring, utf and ascii chars:")
	p(substr("Ура! Привет, Vasya! чувак!", 5, 19))

	// text to binary
	p("string - bytes conversion:")
	text := `<a href="http://home.net/?page=1">Первая страница</a>`
	textBytes := []byte(text)
	p(textBytes)

	// binary to text
	textFormBytes := string(textBytes)
	p(textFormBytes)

}

//
func substr(str string, start int, end int) string {
	return string(([]rune(str))[start:end])
}
