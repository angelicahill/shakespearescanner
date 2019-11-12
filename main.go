package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("static"))))

	fmt.Println("Hello and welcome to the Shakespeare Scanner.\n This is a tool which allows you to search a Shakespeare play for a word\n and it will tell you both how many times it shows up, as well as where it shows up.\n")
	for {
		fmt.Println("Please type the title of the play you would like to search for your word...")
		Scanner := bufio.NewScanner(os.Stdin)
		Scanner.Scan()
		userPlayChoice := strings.ToLower(Scanner.Text())
		fmt.Printf("Great, so you want to search in %s correct?\nPlease type yes or no.\n", userPlayChoice)
		Scanner.Scan()
		confirmation := strings.ToLower((Scanner.Text()))
		if confirmation == "yes" {
			fmt.Printf("Great! Please tell me what word you would like to search for in %s?\n", userPlayChoice)
		} else if confirmation == "no" {
			fmt.Println("Ok sorry about that. Please tell me again what play you are looking for...")
			play2ndTry := bufio.NewScanner(os.Stdin)
			play2ndTry.Scan()
			confirmation := strings.ToLower((play2ndTry.Text()))
			fmt.Printf("Let's try this again, is the play you're interested in %s?\nPlease type yes or no.\n", confirmation)
			Scanner := bufio.NewScanner(os.Stdin)
			Scanner.Scan()
			finalCheck := strings.ToLower((Scanner.Text()))
			if finalCheck == "yes" {
				userPlayChoice := confirmation
				fmt.Printf("Great, now you can tell me what word you want to look for in %s?\n", userPlayChoice)
			} else if finalCheck == "no" {
				fmt.Println("Sorry I'm obviously having a bad day. Please try again later. Goodbye!\n")
				break
			}
		}
		Scanner = bufio.NewScanner(os.Stdin)
		Scanner.Scan()
		finalWord := strings.ToLower((Scanner.Text()))
		fmt.Printf("Ok so just to confirm the word you want to search is %s.\n Now searching...\n", finalWord)
		getPlay, err := ioutil.ReadFile(userPlayChoice + ".txt")
		if err != nil {
			fmt.Println("Sorry we do not currently have that play in our database, please try another play.\n")
			return
		}
		playString := string(getPlay)
		playLines := strings.Split(playString, "\n")
		lastActIndex := 0
		acts := []string{}
		for lineIndex, line := range playLines {
			if line == "ACT I" {
				continue
			}
			findActs := strings.HasPrefix(line, "ACT ")
			if findActs == true {
				actNumeral := strings.Fields(line)[1]
				characters := strings.Split(actNumeral, "")
				isNotRoman := false
				for _, character := range characters {
					if character == "V" || character == "I" || character == "X" {
					} else {
						isNotRoman = true
					}
				}
				if isNotRoman == true {
					continue
				} else {
					actLines := playLines[lastActIndex:lineIndex]
					act := strings.Join(actLines, "\n")
					acts = append(acts, act)
				}
			}
		}

		lastAct := playLines[lastActIndex:]
		lastActadd := strings.Join(lastAct, "\n")
		acts = append(acts, lastActadd)

		for actNumber, act := range acts {
			lowerCaseAct := strings.ToLower(act)
			reg, err := regexp.Compile("[^a-zA-Z0-9]+")
			if err != nil {
				log.Fatal(err)
			}
			splitbySpace := reg.ReplaceAllLiteralString(lowerCaseAct, "")
			wordCount := strings.Count(splitbySpace, finalWord)
			fmt.Printf("%s showed up in your play %v times in Act %v\n", finalWord, wordCount, actNumber+1)
		}

		fmt.Println("Would you like to search for a word in another play, or a different word in this play?\n")
		Scanner = bufio.NewScanner(os.Stdin)
		Scanner.Scan()
		anotherSearch := strings.ToLower((Scanner.Text()))
		if anotherSearch == "yes" {
			continue
		} else if anotherSearch == "no" {
			fmt.Println("Ok, hope this can be helpful to you again soon.\n")
			break
		}

	}
}

func processingAct(userWord string, acts []string) {

	userWordMap := countingWords(acts)
	specificWordCount, exists := userWordMap[userWord]
	if exists {
		fmt.Sprintf("%v showed up %d time(s)\n", userWord, specificWordCount)
	} else {
		fmt.Sprintf("sorry %v doesn't exist!\n", userWord)
	}
}

func countingWords(acts []string) map[string]int {
	wordCountMap := make(map[string]int)

	for _, word := range acts {
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
