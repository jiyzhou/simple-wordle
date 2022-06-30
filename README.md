# Simple Wordle

## Description
A wordle game written in Go.

### Word Filtering
[words_alpha.txt](words_alpha.txt) contains all words. Since only five-letter words are needed in the wordle game, a new list [words_valid.txt](words_valid.txt) was created by running [processWords.go](processWords.go). It simply filters out all words whose length is not five.

### Interaction
Each time the player enters a guess, the program determines whether it hits the target. If not, the guess given will be compared to the target, and the word given by the player will be colored to show its difference from the target. Green means the letter is right where it should be; and yellow means it is contained in the target word though the position is not correct.

### Database Searching
At the beginning of each round, a random word is chosen from [words_valid.txt](words_valid.txt) as the target. Since each line in the list file consists of a five-letter word and a \n symbol, the total number of words can be calculated by dividing 6 into the size of the file. Then a random integer no more than the number of lines is generated, determing where to find the target word.

Before comparing the guess to the target, the validity of the guess must be checked first, that is, if the guess is a valid English word. If so, it should appear in the list [words_valid.txt](words_valid.txt). To search a word in the list efficiently, a trie is built. The key lookup complexity remains proportional to the key size this way.

## Acknowledgments
* [List Of English Words](https://github.com/dwyl/english-words)
* [Using colors with printf](https://stackoverflow.com/questions/5412761/using-colors-with-printf)
* [Trie - Wikipedia](https://en.wikipedia.org/wiki/Trie)
* [Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/solution/)