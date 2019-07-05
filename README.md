# Shakespeare Word Scanner

**What is this program and what is its purpose?**

This program is a Shakespeare word scanner. 

The purpose of this program is to make it easy for a user to search Shakespeare plays for specific words by giving them the ability to input any word into the terminal and find out quickly and easily how many times it shows up in a group of, in this case for the MVP, five Shakespeare plays: Romeo and Juliet, The Tempest, Macbeth, King Lear and Hamlet.

This is a MVP of a larger scanner which will scan the entirety of Shakespeare’s works, including his sonnets, and return how many times a word shows up. 

**How would you use this?**

This tool could be extremely useful to students in high school and beyond during their education in English Literature. 

Being able to identify words that repeat numerous times in a play e.g “love” or “death” can help identify themes both within specific plays (e.g. the word “love” in Romeo and Juliet), as well as across Shakespeare’s works by genre (The word “war” in his history plays, or “fop” in his comedies) or his works in their entirety. 

This is also a tool as it’s root that could be applied to other bodies of work, such as Dickens, Chaucer, or any other author and/or playwright. 

I can also see this being a good program for application to non-academic purposes e.g. scanning for your name in a body of works, or scanning history books for your birthday to see what happened on that day. 

I could go on, there are a wealth of use cases for this simple but fun mini-program. 

**Why did I personally write this program?**

I wrote this program in order to practice all I have learnt thus far in relation to writing functions, ranging over text files, creating channels, and running my program concurrently using go routines. 

It also required me to refactor my code a number of times from being an extremely long 200+ line long program to this version at 93 lines which was good practice.

I plan to integrate on this MVP in the future as I build out my Go skillset further. 

**How does it work?** 

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

**How to run program?**

If you would like to run this program yourself in the terminal all you have to do is: 
- Clone this repository.  
- Open your terminal and navigate to this directory before running “go run shakespearescanner.go” 
- The terminal should then display the following: “Please pick a word and then I will tell you how often is shows up in each of the following Shakespeare plays: The Tempest, Hamlet, King Lear, Macbeth, and Romeo and Juliet. Please pick a word...” 
- Input a word. 
- The program will run and display something like the following: 

```
love showed up 12 time(s) in ariel.txt
 love showed up 67 time(s) in hamlet.txt
 love showed up 51 time(s) in kinglear.txt
 love showed up 19 time(s) in macbeth.txt
 love showed up 135 time(s) in romeo.txt

Time taken to search through plays and give you your results: 660.80686ms
```

