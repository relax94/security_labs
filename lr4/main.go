package main

import (
	"os"
	"log"
	"fmt"
	"crypto/sha256"
	"container/list"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(path) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func ReadFilePart(file *os.File, part rune) ([]byte, bool) {
	data := make([]byte, part)
	count, err := file.Read(data)
	if err == nil {
		return data[:count], false
	} else {
		return nil, true
	}
}

func ElementAt (l *list.List, index int) *list.Element {
	var listIndex int = 0
	for e := l.Front(); e != nil; e = e.Next() {
		if listIndex == index {
			return e
		}
		listIndex++
		// do something with e.Value
	}
	return nil
}


func GenerateBlocks(file *os.File, partSize rune) *list.List {
	blocksList := list.New()
	for {
		part, isEof := ReadFilePart(file, partSize)
		if !isEof {
			blocksList.PushBack(part)
		} else {
			blocksList.PushBack([]byte{})
			break
		}
	}

	return blocksList
}

func GetSHA256(block []byte) []byte {
	sha_256 := sha256.New()
	sha_256.Write(block)
	return sha_256.Sum(nil)
}


func Solve(blocks *list.List) []byte {
	blockIndex := blocks.Len() - 3
	resultBlock := ElementAt(blocks, blockIndex + 1).Value.([]byte)

	lastBlockIndex := -1
	for lastBlockIndex != blockIndex {
		shBlock := GetSHA256(resultBlock)
		mainBlock := ElementAt(blocks, blockIndex).Value.([]byte)
		resultBlock = append(mainBlock, shBlock...)
		blockIndex -= 1
	}

	return GetSHA256(resultBlock)
}

func main() {

	var file *os.File = OpenFile("d://test.mp4")
	defer file.Close()

	var blocks *list.List = GenerateBlocks(file, 1024)
	var result []byte = Solve(blocks)

	fmt.Printf("h0:\t%x\n", result)

}
