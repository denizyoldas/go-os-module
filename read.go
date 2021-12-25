package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
)

const FILE = "test.txt"


func readFile() {
	readTest, err := os.ReadFile(FILE)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(readTest))
	}
}

func readBuffer() {
	readTest, err := os.Open(FILE)
	if err != nil {
		fmt.Println(err)
	} else {
		bufRead := bufio.NewReader(readTest)
		io.Copy(os.Stdout, bufRead)
	}
}

func readFileSeek() {
	f, _ := os.Open(FILE)
	f.Seek(5, 1)
	readByte := make([]byte, 4)
	f.Read(readByte)

	fmt.Println(string(readByte))
}
