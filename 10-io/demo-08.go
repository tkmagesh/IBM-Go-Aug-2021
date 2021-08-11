package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/* fileReader, err := os.Open("data2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fileReader.Close() */
	//scanner := bufio.NewScanner(fileReader)
	scanner := bufio.NewScanner(os.Stdin)
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		fmt.Printf("Line #%d : %s\n", lineCount, line)
	}
}
