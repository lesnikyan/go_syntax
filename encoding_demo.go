package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"log"
	"net/url"
	// "encoding/gob"
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
	str := "Hello Юникондый name!.."
	str64 := base64.StdEncoding.EncodeToString([]byte(str))
	pf("encoded to base64: %s", str64)
	str64decBytes, _ := base64.StdEncoding.DecodeString(str64)
	str64dec := string(str64decBytes)
	pf("decoded from base64: %s", str64dec)

	// struct (arr, map) src
	allSrc := testStructData()

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

	// struct (arr, map) -> bin encode / decode
	serializationDemo()

	// bytes - unicode

	// zip

	//

}

func testStructData() AllData {
	return AllData{
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
}

// struct, arr, map, to bytes
func serializationDemo() {
	var buff bytes.Buffer

	// 1. array encoding as interface
	enc := gob.NewEncoder(&buff)
	data := []int{11, 22, 33}
	data2 := interface{}(data)
	err := enc.Encode(&data2)
	if err != nil {
		log.Fatal("Encode error.", err)
	}
	p("Gob.enc = ", bmap(buff.Bytes(), func(x byte) string { return sf("%02x", x) }))

	// decoding

	var arrRes interface{}
	dec := gob.NewDecoder(&buff)
	err = dec.Decode(&arrRes)
	// binArrInt = []int(binArrIntItr)
	if err != nil {
		log.Fatal("Decode error.", err)
	}
	p("Decoded array:\n", arrRes)

	// 2. map of arrays encoding

	buff.Reset()
	enc2 := gob.NewEncoder(&buff)

	mapSrc := map[string][]int{
		"Aaa": []int{1, 11, 111},
		"Bbb": []int{2, 22, 222},
		"Ccc": []int{3, 33, 333},
	}

	if err2 := enc2.Encode(&mapSrc); err2 != nil {
		log.Fatal("Encoding of map error.", err2)
	}

	// decoding
	var mapRes map[string][]int
	dec2 := gob.NewDecoder(&buff)

	if err2 := dec2.Decode(&mapRes); err2 != nil {
		log.Fatal("Decoding map error.", err2)
	}
	// test decoded
	p("Decoded map data:")
	for key, val := range mapRes {
		pf("%s : %v", key, val)
	}

	// 3. complex struct encoding
	buff.Reset()
	srtSrc := testStructData()
	// using same encoder and decoder

	if err3 := enc2.Encode(&srtSrc); err3 != nil {
		log.Fatal("Encoding struct error.", err3)
	}

	//decoding
	var srtRes AllData
	if err3 := dec2.Decode(&srtRes); err3 != nil {
		log.Fatal("Decoding struct error.", err3)
	}
	// test decoded
	p("Decoded struct:")
	p("Name", srtRes.Name)
	pf("Color: %02x, %02x, %02x", srtRes.Color.R, srtRes.Color.G, srtRes.Color.B)
	pf("Nums: %v,\n Vals: %v,\n Guests: %v,\n Options: %v",
		srtRes.Nums, srtRes.Vals, srtRes.Guests, srtRes.Options)

}

// func binDecode

func zipDemo() {

}
