package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type appendixPage struct {
	Title string
	Info  string
}

func appendixHandler(w http.ResponseWriter, r *http.Request) {
	p := appendixPage{Title: "Appendix", Info: "Information on why I created this app"}
	t, _ := template.ParseFiles("appendixpage.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/appendix/", appendixHandler)
	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		word := r.FormValue("word")
		play := r.FormValue("play")

		getPlay, err := ioutil.ReadFile(play + ".txt")
		if err != nil {
			return // TODO: CREATE ERROR - tell the user they're wrong, even though the user should never come here.
		}
		acts := sortingActs(getPlay)
		x := `<!DOCTYPE html>
		<html>
			<head>
				<title>Shakespeare Play Scanner</title>
				<link rel= "stylesheet" type="text/css" href="shakespearescanner2.css"/>
			</head>
		
			<body>
		<h1>Result</h1>
			</body>
		
		</html>
		`
		for actNumber, act := range acts {
			lowerCaseAct := strings.ToLower(act)
			reg, err := regexp.Compile("[^a-zA-Z0-9]+")
			if err != nil {
				log.Fatal(err)
			}
			splitbySpace := reg.ReplaceAllLiteralString(lowerCaseAct, "")
			wordCount := strings.Count(splitbySpace, word)
			answer := fmt.Sprintf("<h5>%s showed up in your play %v times in Act %v</h5>\n", word, wordCount, actNumber+1)
			x = x + answer
		}
		w.Write([]byte(x))
	})

	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8080", nil))

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

func sortingActs(getPlay []byte) []string {
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
	return acts
}
