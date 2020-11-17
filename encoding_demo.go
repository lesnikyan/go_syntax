package main

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
)

type Rgb struct {
	R byte
	G byte
	B byte
}

type AllData struct {
	Name    string
	Size    int64
	Weight  float32
	Color   Rgb
	Nums    []int
	Vals    []float32
	Guests  []string
	Options map[string]int
}

// EncodingDemo ...
func EncodingDemo() {

	// bin - base64
	// wbuff := strings.Builder
	str := "Hello Юникондый name!.."
	str64 := base64.StdEncoding.EncodeToString([]byte(str))
	pf("encoded to base64: %s", str64)
	str64decBytes, _ := base64.StdEncoding.DecodeString(str64)
	str64dec := string(str64decBytes)
	pf("decoded from base64: %s", str64dec)

	// struct (arr, map) src
	allSrc := AllData{
		Name:   "Vasya",
		Size:   1324567890,
		Weight: 12345.54321,
		Color: Rgb{
			R: 0xcc,
			G: 0x44,
			B: 0x22},
		Nums:    []int{111, 222, 333, 444, 555},
		Vals:    []float32{10, 20.2, 30.33, 40.444, 50.550055},
		Guests:  []string{"Arlekin", "Barmaley", "Condor", "Dront"},
		Options: map[string]int{"optOne": 1001, "opt-two": 2002, "opt-3": 3003}}

	// struct (arr, map) - bin

	// struct - json
	allJson, err := json.Marshal(&allSrc)
	if err != nil {
		p("Json enc error:", err)
	}
	p("Struct-2-JSON:")
	p(string(allJson))
	allJsonInd, _ := json.MarshalIndent(&allSrc, "*", "- ")
	p("Struct-2-JSON-Indent:")
	p(string(allJsonInd))

	// json-struct
	unDat := AllData{}
	json.Unmarshal(allJson, &unDat)

	strColor := func(c Rgb) string {
		return sf("Rgb(%02x, %02x, %02x)", c.R, c.G, c.B)
	}

	pf("unDat: N: %s, W: %f, C: %s, \n\tV: %s", unDat.Name, unDat.Weight, strColor(unDat.Color), unDat.Vals)

	// url-encoding
	toUrlStr := "My name is Vasya. / A + B = C; \ntrue && 1 || 2, ## @abc. %% ( 50% )."
	urlFormated := url.QueryEscape(toUrlStr)
	p("uslFormated: ", urlFormated)
	urlDecoded, _ := url.QueryUnescape(urlFormated)
	p("urlDecoded:", urlDecoded)

	// bytes - unicode

	// zip

	//

}

func serializationDemo() {
	// struct, arr, map, to bytes
}
