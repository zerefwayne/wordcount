package main

import (
	"fmt"
	"log"
)

func main() {

	fileName := "shakespeare.txt"

	fmt.Println("Reading from", fileName)

	words, err := readTxtFile(fileName)

	if err != nil {
		log.Fatalln(err)
	}

	executeSimpleWordCount(words)

}
