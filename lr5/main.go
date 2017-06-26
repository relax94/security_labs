package main

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func DecodeHex(key string) string {
	decoded, err := hex.DecodeString(key)
	if err != nil {
		fmt.Println("Errorr")
	}
	return string(decoded)
}

func MakeRequest(url string) int {
	resp, err := http.Get(url)
	fmt.Println(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	//fmt.Println(resp.StatusCode)
	if resp.StatusCode == 404 {
		return 1
	}
	return 0
}

func MakeCharsArray() []int {
	var char_array []int = []int{}

	char_array = append(char_array, 32)
	for start := 97; start < 123; start++ {
		char_array = append(char_array, start)
	}

	for start := 65; start < 91; start++ {
		char_array = append(char_array, start)
	}
	return char_array
}

func ConvertStringToArray(text string) []string {
	var char_array []string = []string{}
	for i := range text {
		char_array = append(char_array, string(text[i]))
	}
	return char_array
}

func CustomConverter(text string) []string {
	var char_array []string = []string{}
	for i := 0; i < len(text); i += 2 {
		ind := i
		char_array = append(char_array, text[ind:ind+2])
	}
	return char_array
}

func Reverse(s string) (ret string) {
	for _, v := range s {
		defer func(r rune) { ret += string(r) }(v)
	}
	return
}

func PaddingOracleAttack(iv string, ct string) string {
	iva := CustomConverter(iv)

	iv_index := len(iva) - 1
	j := 0x1

	//fmt.Println(iv_index)
	bf_range := MakeCharsArray()

	for k := range make([]int, 16) {
		fmt.Println("k=", k)
		temp := iva[iv_index]
		for i := range bf_range {
			num, e := strconv.ParseInt(iva[iv_index], 16, 64)
			if e != nil {
				panic(e)
			}
			nxor := int(num) ^ bf_range[i] ^ j
			xor := strconv.FormatInt(int64(nxor), 16)
			if e != nil {
				panic(e)
			}
			iva[iv_index] = xor
			url := base_url + "" + strings.Join(iva, "") + ct
			//fmt.Println(nxor)
			result := MakeRequest(url)

			if result == 1 {
				plain_text += string(bf_range[i])
				break
			}
			iva[iv_index] = temp
		}

		iv_index -= 1
		j += 1
		list_index := 0
		end_index := len(iva) - 1

		for end_index > iv_index && list_index < len(plain_text) {
			if k == 15 {
				break
			}
			ord, e := strconv.ParseInt(iva[end_index], 16, 32)
			if e != nil {
				panic(e)
			}
			iva[end_index] = strconv.FormatInt(int64(int(ord)^j^(j-1)), 16)
			end_index -= 1
			list_index += 1
		}
	}
	return plain_text
}

var base_url string = "http://crypto-class.appspot.com/po?er="
var plain_text string = ""

func main() {

	//fmt.Println(ConvertStringToArray("sdfsdfsdf"))

	cipher_text := "f20bdba6ff29eed7b046d1df9fb7000058b1ffb4210a580f748b4ac714c001bd4a61044426fb515dad3f21f18aa577c0bdf302936266926ff37dbf7035d5eeb4"
	iv := cipher_text[:32]
	first_block := cipher_text[32:64]

	fmt.Println("PR ", Reverse(PaddingOracleAttack(iv, first_block)))

}
