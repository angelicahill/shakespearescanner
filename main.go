package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"log"
	"net/http"
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

type Play struct {
	XMLName xml.Name `xml:"PLAY"`
	Text    string   `xml:",chardata"`
	TITLE   string   `xml:"TITLE"`
	FM      struct {
		Text string   `xml:",chardata"`
		P    []string `xml:"P"`
	} `xml:"FM"`
	PERSONAE struct {
		Text    string   `xml:",chardata"`
		TITLE   string   `xml:"TITLE"`
		PERSONA []string `xml:"PERSONA"`
		PGROUP  []struct {
			Text     string   `xml:",chardata"`
			PERSONA  []string `xml:"PERSONA"`
			GRPDESCR string   `xml:"GRPDESCR"`
		} `xml:"PGROUP"`
	} `xml:"PERSONAE"`
	SCNDESCR string `xml:"SCNDESCR"`
	PLAYSUBT string `xml:"PLAYSUBT"`
	ACT      []struct {
		Text  string `xml:",chardata"`
		TITLE string `xml:"TITLE"`
		SCENE []struct {
			Text     string   `xml:",chardata"`
			TITLE    string   `xml:"TITLE"`
			STAGEDIR []string `xml:"STAGEDIR"`
			SPEECH   []struct {
				Text    string   `xml:",chardata"`
				SPEAKER []string `xml:"SPEAKER"`
				LINE    []struct {
					Text     string `xml:",chardata"`
					STAGEDIR string `xml:"STAGEDIR"`
				} `xml:"LINE"`
				STAGEDIR []string `xml:"STAGEDIR"`
			} `xml:"SPEECH"`
		} `xml:"SCENE"`
	} `xml:"ACT"`
}

var plays = map[string]string{
	"hamlet":              "https://www.ibiblio.org/xml/examples/shakespeare/hamlet.xml",
	"antonyandcleopatra":  "http://www.ibiblio.org/xml/examples/shakespeare/a_and_c.xml",
	"coriolanus":          "http://www.ibiblio.org/xml/examples/shakespeare/coriolan.xml",
	"juliuscaesar":        "http://www.ibiblio.org/xml/examples/shakespeare/j_caesar.xml",
	"kinglear":            "http://www.ibiblio.org/xml/examples/shakespeare/lear.xml",
	"macbeth":             "http://www.ibiblio.org/xml/examples/shakespeare/macbeth.xml",
	"othello":             "http://www.ibiblio.org/xml/examples/shakespeare/othello.xml",
	"romeoandjuliet":      "http://www.ibiblio.org/xml/examples/shakespeare/r_and_j.xml",
	"timonofathens":       "http://www.ibiblio.org/xml/examples/shakespeare/timon.xml",
	"titusandronicus":     "http://www.ibiblio.org/xml/examples/shakespeare/titus.xml",
	"henryivpart1":        "http://www.ibiblio.org/xml/examples/shakespeare/hen_iv_1.xml",
	"henryivpart2":        "http://www.ibiblio.org/xml/examples/shakespeare/hen_iv_2.xml",
	"henryv":              "http://www.ibiblio.org/xml/examples/shakespeare/hen_v.xml",
	"henryvipart1":        "http://www.ibiblio.org/xml/examples/shakespeare/hen_vi_1.xml",
	"henryvipart2":        "http://www.ibiblio.org/xml/examples/shakespeare/hen_vi_2.xml",
	"henryvipart3":        "http://www.ibiblio.org/xml/examples/shakespeare/hen_vi_3.xml",
	"henryviii":           "http://www.ibiblio.org/xml/examples/shakespeare/hen_viii.xml",
	"kingjohn":            "http://www.ibiblio.org/xml/examples/shakespeare/john.xml",
	"richardii":           "http://www.ibiblio.org/xml/examples/shakespeare/rich_ii.xml",
	"richardiii":          "http://www.ibiblio.org/xml/examples/shakespeare/rich_iii.xml",
	"allswell":            "http://www.ibiblio.org/xml/examples/shakespeare/all_well.xml",
	"asyoulikeit":         "http://www.ibiblio.org/xml/examples/shakespeare/as_you.xml",
	"comedyoferrors":      "http://www.ibiblio.org/xml/examples/shakespeare/com_err.xml",
	"cymbeline":           "http://www.ibiblio.org/xml/examples/shakespeare/cymbelin.xml",
	"loveslabourslost":    "http://www.ibiblio.org/xml/examples/shakespeare/lll.xml",
	"measureformeasure":   "http://www.ibiblio.org/xml/examples/shakespeare/m_for_m.xml",
	"merrywives":          "http://www.ibiblio.org/xml/examples/shakespeare/m_wives.xml",
	"merchantofvenice":    "http://www.ibiblio.org/xml/examples/shakespeare/merchant.xml",
	"midsummerightsdream": "http://www.ibiblio.org/xml/examples/shakespeare/dream.xml",
	"muchadoaboutnothing": "http://www.ibiblio.org/xml/examples/shakespeare/much_ado.xml",
	"pericles":            "http://www.ibiblio.org/xml/examples/shakespeare/pericles.xml",
	"tamingoftheshrew":    "http://www.ibiblio.org/xml/examples/shakespeare/taming.xml",
	"tempest":             "http://www.ibiblio.org/xml/examples/shakespeare/tempest.xml",
	"troilusandcressida":  "http://www.ibiblio.org/xml/examples/shakespeare/troilus.xml",
	"twelfthnight":        "http://www.ibiblio.org/xml/examples/shakespeare/t_night.xml",
	"twogentelman":        "http://www.ibiblio.org/xml/examples/shakespeare/two_gent.xml",
	"winterstale":         "http://www.ibiblio.org/xml/examples/shakespeare/win_tale.xml",
}

func main() {
	http.HandleFunc("/appendix/", appendixHandler)
	http.HandleFunc("/run2", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getPlay(playChoice string) (Play, error) {
	resp, err := http.Get(playChoice)
	if err != nil {
		fmt.Print("Failed to read text file.", err)
		return Play{}, err
	}
	defer resp.Body.Close()
	var p Play
	err = xml.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		fmt.Print("Failed to read text file.", err)
		return Play{}, err

	}
	return p, nil
}
