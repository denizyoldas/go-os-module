package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Dosya olusturuldu : \n")
	// writeOsWrite()
	// writeFile()
	// readFile()i
	readFileSeek()
}


func writeOsWrite() {
	err := os.WriteFile("test.txt", []byte("test text buralarda"), os.ModePerm)
	if err != nil {
		fmt.Printf("hata: %v\n", err)
	}
}

func writeFile() {
	f, err := os.Create("test2.txt")
	if err != nil {
		fmt.Printf("hata: %v\n", err)
	}

	for i := 0; i < 10; i++ {
		f.Write([]byte("data\n"))
	}

	f.Close()
}
