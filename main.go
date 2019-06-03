package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	fmt.Printf("Please pick a word and then I will tell you how often is shows up in each of the following Shakespeare plays: The Tempest, Hamlet, King Lear, Macbeth, and Romeo and Juliet. Please pick a word...")
	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()
	userInputWord := Scanner.Text()
	userInputLowercase := strings.ToLower(userInputWord)

	timeAtStart := time.Now()

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)
	ch5 := make(chan string)

	go processingPlay(userInputLowercase, "ariel.txt", ch1)
	go processingPlay(userInputLowercase, "hamlet.txt", ch2)
	go processingPlay(userInputLowercase, "kinglear.txt", ch3)
	go processingPlay(userInputLowercase, "macbeth.txt", ch4)
	go processingPlay(userInputLowercase, "romeo.txt", ch5)

	tempest := <-ch1
	hamlet := <-ch2
	kinglear := <-ch3
	macbeth := <-ch4
	romeo := <-ch5
	fmt.Println(tempest, hamlet, kinglear, macbeth, romeo)

	fmt.Printf("Time taken to search through plays and give you your results: %v\n", time.Since(timeAtStart))
}

func processingPlay(userWord string, fileName string, x chan string) {
	thePlay, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(fmt.Errorf("failed to read fule: %s", err))
	}
	playWords := strings.Fields(string(thePlay))
	playWordMap := countingWords(playWords)
	specificWordCount, exists := playWordMap[userWord]
	if exists {
		x <- fmt.Sprintf("%v showed up %d time(s) in %v\n", userWord, specificWordCount, fileName)
	} else {
		x <- fmt.Sprintf("sorry %v doesn't exist in %v!\n", userWord, fileName)
	}
}

func countingWords(playWords []string) map[string]int {
	wordCountMap := make(map[string]int)

	for _, word := range playWords {
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		nonPuncWord := reg.ReplaceAllLiteralString(word, "")
		lowerCaseWord := strings.ToLower(nonPuncWord)
		value, exists := wordCountMap[lowerCaseWord]
		if exists {
			wordCountMap[lowerCaseWord] = value + 1
		} else if !exists {
			wordCountMap[lowerCaseWord] = 1
		}
	}
	return wordCountMap
}

func printResultingValues(wordCountMap map[string]int) {
	for k, v := range wordCountMap {
		if v == 1 {
			fmt.Printf("%v showed up 1 time", k)
			fmt.Println()
			fmt.Println()
		} else {
			fmt.Printf("%v showed up %d times", k, v)
			fmt.Println()
			fmt.Println()
		}
	}
}
