package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const charPerLine = 6 // each line consists of a 5-letter word and a \n
const wordLength = 5
const maxTries = 6 // the player can try no more than 6 times
const colorGreen = "\033[1;32m%s\033[0m"
const colorYellow = "\033[1;33m%s\033[0m"
const colorWhite = "\033[1;37m%s\033[0m"

// a trie is a tree data structure used to efficiently store and retrieve keys in a dataset of strings
type Trie struct {
	children [26]*Trie // each link corresponds to one of the letter values
	isEnd    bool      // specifies whether the node corresponds to the end of the key
}

func main() {
	file, err := os.Open("words_valid.txt")
	checkErr(err)
	defer file.Close()

	// construct the trie
	var trie Trie
	trie.InitizeTrie(file)

	// get the file size
	stat, err := file.Stat()
	checkErr(err)

	line := stat.Size() / charPerLine      // calculate the number of lines
	rand.Seed(time.Now().UTC().UnixNano()) // seed the random generator
	chosenLine := rand.Intn(int(line))     // pick a random number from 0 to the maximum line
	offset := chosenLine * charPerLine     // the offset from the beginning
	_, err = file.Seek(int64(offset), 0)   // use seek() to set offset for the next Read
	checkErr(err)

	chosenWord := make([]byte, wordLength) // initize a fixed size slice to hold the target word
	_, err = file.Read(chosenWord)
	checkErr(err)
	targetWord := string(chosenWord[:wordLength]) // convert bytes into string

	fmt.Printf("Hey welcome to the wordle game! You have %v tries to gess the wordle. Each guess must be a valid %v letter word in lowercase. Hit the return to submit it.\n", maxTries, wordLength)

	i := 1
	for i <= maxTries {
		fmt.Printf("#%v Try:\n", i) // prompt
		var guess string
		fmt.Scanln(&guess) // read user input

		// check if the length is right
		if len(guess) != wordLength {
			fmt.Printf("Please enter a %v-letter word!\n", wordLength)
			continue
		}

		if guess == "CHEAT" { // so this gives the player the result at once
			fmt.Printf("The word is %s. Keep it secret!\n", targetWord)
			continue
		}

		// check if the guess is a valid word
		if trie.Search(guess) {
			i++
			if guess == targetWord { // yeah! bingo
				fmt.Println("Congratulations! You win!")
				break
			}
			for j := 0; j < wordLength; j++ { // compare the guess with the target word
				if guess[j] == targetWord[j] { // the character is right where it should be, mark it as green
					fmt.Printf(colorGreen, guess[j:j+1])
				} else if strings.Contains(targetWord, guess[j:j+1]) { // the current character appears in the target word, mark it as yellow
					fmt.Printf(colorYellow, guess[j:j+1])
				} else { // a miss
					fmt.Printf(colorWhite, guess[j:j+1])
				}
			}
			fmt.Println("")
		} else {
			fmt.Println("Not a word!")
		}
	}

	fmt.Printf("The word is %v.\n", targetWord)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// construct a trie from the given words database
func (t *Trie) InitizeTrie(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text() // for each word in the text file
		node := t
		for _, ch := range word {
			ch -= 'a'                     // lowercase latin letters
			if node.children[ch] == nil { // a link does not exist
				node.children[ch] = &Trie{} // create a new node and link it with the parent's link matching the current key character
			}
			node = node.children[ch]
		}
		node.isEnd = true // repeat until the last character
	}
}

// returns true if the string word is in the trie, and false otherwise
func (t *Trie) Search(word string) bool {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil { // if there is no link correponds to the current key character, the word doesn't exist in the trie
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd // not necessary for fixed length words, just return true works fine
}
