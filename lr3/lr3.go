package main

import (
	//b64 "encoding/base64"
	hex "encoding/hex"
	"fmt"
	"log"
	"crypto/aes"
	"strconv"
	"os"
	"bufio"
)

func DecodeHex(key string) string {
	decoded, err := hex.DecodeString(key)
	if err != nil {
		log.Fatal(err)
	}
	return string(decoded)
}

func GetCountCycles(text string) rune {
	return rune((len(text) / 32) - 1)
}

func DecryptAes128Ecb(data, key []byte) []byte {
	cipher, _ := aes.NewCipher([]byte(key))
	decrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}

func EncryptAes128Ecb(data, key []byte) []byte {
	cipher, _ := aes.NewCipher([]byte(key))
	decrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}

func DecryptBlock(block []byte, str string) string {
	var block_result string = ""
	for j, i := range block {
		t := i ^ str[j]
		j += 1
		if t < 16 {
			block_result += "0"
		}
		block_result += strconv.FormatInt(int64(t), 16)
	}

	return block_result
}

func SolveCBCEncryption(key string, text string) string {
	var cbc string = ""
	key1 := DecodeHex(key)
	cycles := GetCountCycles(text)

	for cycle := range make([]rune, cycles){
		s := DecodeHex(text[(cycle + 1) * 32:(cycle + 2) * 32])
		block := DecryptAes128Ecb([]byte(s), []byte(key1))
		str2 := DecodeHex(text[cycle * 32:((cycle + 1) * 32)])
		str3 := DecryptBlock(block, str2)
		cbc += str3
	}
	return cbc
}

func GetParts(text string) (byte, string) {
	ciper := DecodeHex(text[:32])
	last := len(ciper) - 1
	front := ciper[last]
	back := ciper[:last]
	return front, back
}

func SolveCTRModeEncryption(key string, text string) string {
	var ctr string = ""
	front, back := GetParts(text)

	for c := range make([]int, int((len(text) / 64) - 1)) {
		encrypt := append([]byte(back), front + byte(c))
		message_block := EncryptAes128Ecb(encrypt, []byte(DecodeHex(key)))
		key_block := DecodeHex(text[(c + 1) * 32:(c + 1) * 64])
		ctr += DecryptBlock(message_block, key_block)
	}
	return DecodeHex(ctr)
}

func StartCBC() {
	var key string = "140b41b22a29beb4061bda66b6747e14"
	var text string ="4ca00ff4c898d61e1edbf1800618fb2828a226d160dad07883d04e008a7897ee2e4b7465d5290d0c0e6c6822236e1daafb94ffe0c5da05d9476be028ad7c1d81"

	var cbc string = SolveCBCEncryption(key, text)

	PrintResult(cbc, "./cbc.txt")
}

func StartCTR() {

	var key string = "36f18357be4dbd77f050515c73fcf9f2"
	var text string = "69dda8455c7dd4254bf353b773304eec0ec7702330098ce7f7520d1cbbb20fc388d1b0adb5054dbd7370849dbf0b88d393f252e764f1f5f7ad97ef79d59ce29f5f51eeca32eabedd9afa932908080808080808080808080808080808080808080808080808080808080808008080808080808080808080808080808080808080808080808080808080808080808080808080808080808080808080808080808080808080808080808080808080808080808080"

	var ctr string = SolveCTRModeEncryption(key, text)
	PrintResultsToFile(ctr, "./ctr.txt")
}

func main() {

	StartCBC()
	StartCTR()
}

func PrintResult( b string,  file string) {
	var response string = ""
	for i := 0; i < len(b); i+=2{
		hexCode := string(b[i]) + string(b[i+1])
		number, _ := strconv.ParseInt(hexCode, 16, 8)
		response += fmt.Sprintf("%s", string(number))
	}

	PrintResultsToFile(response, file)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func PrintResultsToFile(response string, file string) {
	f, e := os.Create(file)
	check(e)
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(response)

	w.Flush()

	fmt.Println("Results printed to file")
}
