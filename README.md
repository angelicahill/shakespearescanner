# Shakespeare Word Scanner

**What is this program and what is its purpose?**

This program is a Shakespeare word scanner. 

The purpose of this program is to make it easy for a user to search Shakespeare plays for specific words by giving them the ability to input any word and find out quickly and easily how many times it shows up in Shakespeare's core portfolio of plays (they must use specific input stylings currently for the program to work as seen below and in key, which can be found on the apps' static site): 

### Tragedy
Antony and Cleopatra = antonyandcleopatra
Coriolanus = Coriolanus
Hamlet = Hamlet
Julius Caesar = juliuscaesar
King Lear = kinglear
Macbeth = Macbeth
Othello = Othello
Romeo and Juliet = romeoandjuliet
Timon on Athens = timonofathens
Titus Andronicus = titusandronicus

### History
Henry IV, part 1 = henryivpart1
Henry IV, part 2 = henryivpart2
Henry V = Henryv
Henry VI, part 1 = henryvipart1
Henry VI, part 2 = henryvipart2
Henry VI, part 3 = henryvipart3
Henry VIII = henryviii
King John = kingjohn
Richard II = richardii
Richard III = richardiii

### Comedy
All's Well That Ends Well = allswell
As You Like It = asyoulikeit
The Comedy of Errors = comedyoferrors
Cymbeline = Cymbeline
Love’s Labours Lost = loveslabourslost
Measure for Measure = measureformeasure
The Merry Wives of Windsor = merry wives
The Merchant of Venice = merchantofvenice
A Midsummer Night’s Dream = midsummerightsdream
Much Ado About Nothing = muchadoaboutnothing
Pericles, Prince of Tyre = Pericles
Taming of the Shrew = tamingoftheshrew
The Tempest = tempest
Troilus and Cressida = troilusandcressida
Twelfth Night = twelfthnight
Two Gentlemen of Verona = twogentelman
Winters Tale = winterstale

***This is a MVP of a larger scanner which will scan the entirety of Shakespeare’s works, including his sonnets, and return how many times a word shows up.** 

**How to run the program?**

If you would like to run this program yourself in the terminal all you have to do is: 
- Clone this repository.  
- Open your terminal and navigate to this directory before running “go run main.go” 
- The terminal should then display the following: “Would you like to play around in the terminal or on my website?” 
- Input your choice "terminal" or "web" and then you can either input the play name and word in the text fields on the site, or enter then as prompted in the terminal app. 
- The program will run and display something like the following: 

```
love showed up in your play 13 times in Act 1
love showed up in your play 29 times in Act 2
love showed up in your play 60 times in Act 3
love showed up in your play 74 times in Act 4
love showed up in your play 84 times in Act 5

```

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

Firstly I give the user the option to run and play with the scanner in the terminal or in web: 

```
for {
		fmt.Println("Would you like to play around in the terminal or on my website?")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userTerminalorWebChoice := strings.ToLower(scanner.Text())
		if userTerminalorWebChoice == "terminal" {
			terminalVersion()
		} else if userTerminalorWebChoice == "website" {
			websiteVersion()
		} else {
			fmt.Println("Please select web or terminal. Thank you!")
			continue
		}
	}

```

Then depending on their choice they either are taken straight into the terminal where they can use the scanner there, or if they select "web" the should navigate to `localhost:8080` where they will be able to input the play and word they are interestd in into text fields and then press the "SUBMIT" button and call the API I will build which will return the values they are interested in. **THIS PART IS THE PART I AM CURRENTLY WORKING ON**

