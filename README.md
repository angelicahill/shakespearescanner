# Shakespeare Word Scanner

**What is this program and what is its purpose?**

This program is a Shakespeare word scanner. 

The purpose of this program is to make it easy for a user to search Shakespeare plays for specific words by giving them the ability to input any word and find out quickly and easily how many times it shows up in Shakespeare's core portfolio of play. 

***This is a MVP of a larger scanner which will scan the entirety of Shakespeare’s works, including his sonnets, and return how many times a word shows up.** 

**How to run the program?**

If you would like to run this program yourself in the terminal all you have to do is: 
- Clone this repository.  
- Open your terminal and navigate to this directory before running “go run main.go” 
- This will then start the server running and if you navigate to `localhost:8080` you should see the homepage of the app. 

**How would you use this?**

This tool could be extremely useful to students in high school and beyond during their education in English Literature. 

Being able to identify words that repeat numerous times in a play e.g “love” or “death” can help identify themes both within specific plays (e.g. the word “love” in Romeo and Juliet), as well as across Shakespeare’s works by genre (The word “war” in his history plays, or “fop” in his comedies) or his works in their entirety. 

This is also a tool as it’s root that could be applied to other bodies of work, such as Dickens, Chaucer, or any other author and/or playwright. 

I can also see this being a good program for application to non-academic purposes e.g. scanning for your name in a body of works, or scanning history books for your birthday to see what happened on that day. 

I could go on, there are a wealth of use cases for this simple but fun mini-program. 

**Why did I personally write this program?**

I wrote this program in order to practice all I have learnt thus far in relation to writing functions, ranging over text files, creating channels, and running my program concurrently using go routines. It was developed as my skillset has developed.

It also required me to refactor my code a number of times from being an extremely long 200+ line long program to 93 lines, and then longer again as I developed and refined it which was good practice.

I plan to interate on this MVP in the future as I build out my Go skillset further. 

**How did it work oringially in V1?** 

This program works by scanning through a series of five .txt files concurrently looking for whatever string was input by the user following making that input lowercase and taking out and ignoring any punctuation or spacing as it ranges over the strings. 

The program counts each instance of the user input until it’s finished running through the text file and then sends that final integer to the channel corresponding to that play and then returns that value input into the string to the terminal for the user to see.  

After setting out the basic structure of the main function I created a function that creates a map “wordCountMap” into which I input the results of ranging over the words in the play, eliminating any numbers, spaces, and/or punctation, making it lowercase, and then returning the number of instances found: 

```
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

```

Then I wrote a minor function which enables the program to say the word shows up 1 time, as the previous logic would have been syntactically incorrect in that is assumed plural and would have have said, for example, “love shows up 1 times in The Tempest”: 

```
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

```

The final function to bring everything together was as follows: 

```
func processingPlay(userWord string, fileName string, x chan string) {
    thePlay, err := ioutil.ReadFile(fileName)
    if err != nil {
        panic(fmt.Errorf("failed to read fle: %s", err))
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

```

A minor part of the program I included is that it tells the user how long it took to scan and get the final values. 

**How does it now work in V2?** 
The basic functionality of the app works. it is now a browser only application. 

Next step is going to be to do more reading around this area of development, especially focusing on template in Go (https://golang.org/pkg/text/template/) and errors (https://golang.org/pkg/errors/). I also plan to then add more styling to http://localhost:8080/run via a separate html file. 

**THIS PART IS THE PART I AM CURRENTLY WORKING ON**
The resources I am looking at are as follows: 
- https://gowebexamples.com/
- https://developer.mozilla.org/en-US/docs/Learn/HTML/Forms/Your_first_HTML_form
- https://thenewstack.io/make-a-restful-json-api-go/
- https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da
- https://github.com/google/jsonapi
- https://github.com/Michael-F-Ellis/goht
- https://meshstudio.io/blog/2017-11-06-serving-html-with-golang/
- https://www.alexedwards.net/blog/serving-static-sites-with-go
- https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
- https://dev.to/hauxe/golang-http-serve-static-files-correctly-2oj2
- https://tutorialedge.net/golang/creating-simple-web-server-with-golang/

