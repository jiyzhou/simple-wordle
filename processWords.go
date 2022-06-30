package main

import (
	"bufio"
	"log"
	"os"
)

// select valid words for the wordle game!
func processWords() {
	inputFile, err := os.Open("words_alpha.txt") // a text of English words that only have letters, no numbers or symbols
	checkErr(err)
	defer inputFile.Close()

	outputFile, err := os.Create("words_valid.txt") // filter out words whose length is not 5
	checkErr(err)
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		word := scanner.Text()
		if len(word) == 5 {
			_, err := outputFile.WriteString(word + "\n")
			checkErr(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
