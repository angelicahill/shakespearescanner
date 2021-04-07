package shakespeare

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type appendixPage struct {
	Title string
	Info  string
}

// AppendixHandler handles the Appendix page of my App.
func AppendixHandler(w http.ResponseWriter, r *http.Request) {
	p := appendixPage{Title: "Appendix", Info: "Information on why I created this app"}
	t, err := template.ParseFiles("../appendixpage.html")
	if err != nil {
		w.WriteHeader(500)
		fmt.Printf("The error is as follow; %v", err)
		return
	}
	_ = t.Execute(w, p)
}

// Run2 takes user input values and processes the information to provide the user with their desired information.
func Run2(w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("word")
	play := r.FormValue("play")

	url, ok := plays[play]
	if !ok {
		w.WriteHeader(404)
		//Explicitly ignoring value with "_ = x"
		_, _ = w.Write([]byte("<h3>Play not found. Please try another play.</h3>"))
		fmt.Printf("Play has not been found.")
		return
	}
	p, err := getPlay(url)
	//TODO: get rid of url peram and re-write with packaging - fixed path not url?!
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("<h3>Failed to read text file. This is an issue on the App side. Please try again later.</h3>"))
		fmt.Println("File has not been found.")
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
		%s
		</body>
	
	</html>
	`
	answers := []string{}
	word = strings.ToLower(word)
	for _, act := range p.ACT {
		wordCount := 0
		for _, scene := range act.SCENE {
			for _, speech := range scene.SPEECH {
				for _, line := range speech.LINE {
					wordCount += strings.Count(strings.ToLower(line.Text), word)
				}

			}
		}
		answer := fmt.Sprintf("<h5>%s showed up in your play %v times in %v</h5>", word, wordCount, act.TITLE)
		answers = append(answers, answer)
	}
	x = fmt.Sprintf(x, strings.Join(answers, "\n"))
	_, _ = w.Write([]byte(x))
}
