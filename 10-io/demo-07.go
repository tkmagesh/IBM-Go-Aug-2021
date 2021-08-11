package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileReader, err := os.Open("data1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	bufferedReader := bufio.NewReader(fileReader)
	sentenceCount := 0
	for {
		sentence, err := bufferedReader.ReadString('.')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		sentenceCount++
		fmt.Printf("Sentence #%d: %s\n", sentenceCount, sentence)
	}
}
