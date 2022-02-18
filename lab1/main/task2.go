package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Constants for stopwords file
	const MaxWordsNumber = 10000
	const MaxStopwordsNumber = 20
	const StopwordsFileName = "stopwords.txt"

	var stopwords [MaxStopwordsNumber]string
	stopwordsNumber := 0

	//---------------------------------------------------------------------------------------------------------------------
	// Working with stopwords file
	file, err := os.Open(StopwordsFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
StopLoop1:
	if scanner.Scan() {
		textLine := scanner.Text()
		ind := 0
		word := ""

	StopLoop2:
		character := string(textLine[ind])

		//end of the current word
		if character == "." || character == " " || character == "," {
			stopwords[stopwordsNumber] = word

			stopwordsNumber++
			ind++

			word = ""
			goto StopLoop2
		}

		//add char to the current word
		word += character
		ind++

		//end of the current line
		if ind >= len(textLine) {
			stopwords[stopwordsNumber] = word
			stopwordsNumber++
			word = ""
			goto StopLoop1
		}

		goto StopLoop2
	}

	//---------------------------------------------------------------------------------------------------------------------
	//Constants for input file
	const DisplayWordsNumber = 25
	const TextFileName = "input.txt"

	lowLetters := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q",
		"r", "s", "t", "u", "v", "w", "x", "y", "z"}
	upperLetters := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q",
		"R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	var words [MaxWordsNumber]string
	var wordNumber [MaxWordsNumber]int
	WordsNumber := 0

	const MaxPages = 10000
	var wordsPages [MaxWordsNumber][MaxPages]int

	//---------------------------------------------------------------------------------------------------------------------
	//Working with input file
	file, err = os.Open(TextFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)

	pageNum := 0
	lineNum := 0

PageLoop:
	pageNum++
	lineNum = 0

InputLoop1:
	lineNum++
	if scanner.Scan() {
		textLine := scanner.Text()
		nextLine := false
		ind := 0
		word := ""
		if lineNum > 45 {
			goto PageLoop
		}

	InputLoop2:
		//end of the current line
		if ind >= len(textLine) {
			goto InputLoop1
		}

		character := string(textLine[ind])
		letterInd := 0

		//end of the current word
		if character == "." || character == " " || character == "," {
			ind++
			goto WordLoop1
		}

		//if letter is upper turn it to lower
	UpperToLowerLoop:
		if character == upperLetters[letterInd] {
			character = lowLetters[letterInd]
		} else if letterInd < 25 {
			letterInd++
			goto UpperToLowerLoop
		}

		//add char to the current word
		word += character
		ind++

		if ind >= len(textLine) {
			nextLine = true
			goto WordLoop1
		}
		goto InputLoop2

		//Add word to list
	WordLoop1:
		wordInd := 0
	WordLoop2:
		if words[wordInd] == "" {
			WordsNumber++
			words[wordInd] = word
			wordNumber[wordInd]++
			wordsPages[wordInd][0]++
			wordsPages[wordInd][1] = pageNum
			word = ""
			if nextLine {
				goto InputLoop1
			}
			goto InputLoop2
		} else if words[wordInd] == word {
			wordNumber[wordInd]++
			if wordsPages[wordInd][wordsPages[wordInd][0]] != pageNum {
				wordsPages[wordInd][0]++
				wordsPages[wordInd][wordsPages[wordInd][0]] = pageNum
			}
			word = ""
			goto InputLoop2
		} else if word == "" {
			if nextLine {
				goto InputLoop1
			}
			goto InputLoop2
		}
		wordInd++
		goto WordLoop2
	}

	//bubble sort the list
	i, j := 0, 0

	if WordsNumber <= 1 {
		goto Sorted
	}

SortLoop1:
	j = 0

	//end of the list
	if i+1 >= WordsNumber {
		goto Sorted
	}

SortLoop2:
	if words[j] > words[j+1] {
		temp := wordNumber[j]
		wordNumber[j] = wordNumber[j+1]
		wordNumber[j+1] = temp

		temp1 := words[j]
		words[j] = words[j+1]
		words[j+1] = temp1

		temp2 := wordsPages[j]
		wordsPages[j] = wordsPages[j+1]
		wordsPages[j+1] = temp2
	}

	if j+i+2 >= WordsNumber {
		i++
		goto SortLoop1
	}
	j++

	goto SortLoop2

Sorted:
	displayNumber := 0
	displayInd := 0
	stopInd := 0

DisplayLoop:
	stopInd = 0

	//check if the word is a stopword
CheckLoop:
	if words[displayInd] == stopwords[stopInd] {
		displayInd++
		goto DisplayLoop
	}
	stopInd++
	if stopInd < stopwordsNumber {
		goto CheckLoop
	}

	//display the word + number of repeat
	displayNumber++
	printed := 0
	fmt.Print(displayNumber, ".", words[displayInd], "---")

PrintCycle:
	printed++
	fmt.Print(wordsPages[displayInd][printed], " ")
	if printed < wordsPages[displayInd][0] {
		goto PrintCycle
	}
	fmt.Print("\n")
	displayInd++
	if displayInd < WordsNumber {
		goto DisplayLoop
	}

}
