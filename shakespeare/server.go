package shakespeare

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

//TODO: Add testing using "net/http/httptest" package.

type appendixPage struct {
	Title string
	Info  string
}

func AppendixHandler(w http.ResponseWriter, r *http.Request) {
	p := appendixPage{Title: "Appendix", Info: "Information on why I created this app"}
	t, _ := template.ParseFiles("appendixpage.html")
	t.Execute(w, p)
}

func Run2(w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("word")
	play := r.FormValue("play")

	url, ok := plays[play]
	if !ok {
		fmt.Println("Play has not been found. Please try another play.")
		return
	}
	p, err := getPlay(url)
	if err != nil {
		fmt.Print("Failed to read text file.", err)
		return
	}
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
	word = strings.ToLower(word)
	for _, act := range p.ACT {
		wordCount := 0
		for _, scene := range act.SCENE {
			for _, speech := range scene.SPEECH {
				for _, line := range speech.LINE {
					wordCount = wordCount + strings.Count(strings.ToLower(line.Text), word)
				}

			}
		}
		answer := fmt.Sprintf("<h5>%s showed up in your play %v times in %v</h5>\n", word, wordCount, act.TITLE)
		x = x + answer
	}

	w.Write([]byte(x))
}
