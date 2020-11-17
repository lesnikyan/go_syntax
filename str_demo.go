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
	// strings.Start

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

/*

For information about UTF-8 strings in Go, see
https://blog.golang.org/strings.

func Compare(a, b string) int
func Contains(s, substr string) bool
func ContainsAny(s, chars string) bool
func ContainsRune(s string, r rune) bool
func Count(s, substr string) int
func EqualFold(s, t string) bool
func Fields(s string) []string
func FieldsFunc(s string, f func(rune) bool) []string
func HasPrefix(s, prefix string) bool
func HasSuffix(s, suffix string) bool
func Index(s, substr string) int
func IndexAny(s, chars string) int
func IndexByte(s string, c byte) int
func IndexFunc(s string, f func(rune) bool) int
func IndexRune(s string, r rune) int
func Join(elems []string, sep string) string
func LastIndex(s, substr string) int
func LastIndexAny(s, chars string) int
func LastIndexByte(s string, c byte) int
func LastIndexFunc(s string, f func(rune) bool) int
func Map(mapping func(rune) rune, s string) string
func Repeat(s string, count int) string
func Replace(s, old, new string, n int) string
func ReplaceAll(s, old, new string) string
func Split(s, sep string) []string
func SplitAfter(s, sep string) []string
func SplitAfterN(s, sep string, n int) []string
func SplitN(s, sep string, n int) []string
func Title(s string) string
func ToLower(s string) string
func ToLowerSpecial(c unicode.SpecialCase, s string) string
func ToTitle(s string) string
func ToTitleSpecial(c unicode.SpecialCase, s string) string
func ToUpper(s string) string
func ToUpperSpecial(c unicode.SpecialCase, s string) string
func ToValidUTF8(s, replacement string) string
func Trim(s, cutset string) string
func TrimFunc(s string, f func(rune) bool) string
func TrimLeft(s, cutset string) string
func TrimLeftFunc(s string, f func(rune) bool) string
func TrimPrefix(s, prefix string) string
func TrimRight(s, cutset string) string
func TrimRightFunc(s string, f func(rune) bool) string
func TrimSpace(s string) string
func TrimSuffix(s, suffix string) string
type Builder struct{ ... }
type Reader struct{ ... }
    func NewReader(s string) *Reader
type Replacer struct{ ... }
    func NewReplacer(oldnew ...string) *Replacer

BUG: The rule Title uses for word boundaries does not handle Unicode punctuation properly.
*/
