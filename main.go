package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"slices"
	"strings"
	"time"
)

var fretLength int = 12
var tuning []string
var statusMessage string

func main() {
	collectInput()
	for {
		x, y := rand.Intn(len(tuning)), rand.Intn(fretLength)
		note := drawNote(x, y, false)

		correct := false
		for !correct {
			var guess string
			fmt.Print("Enter note: ")
			fmt.Scanln(&guess)
			guess = strings.Replace(guess, "s", "#", -1)
			if len(guess) > 0 {
				guess = strings.ToUpper(string(guess[0])) + guess[1:]
			} else {
				guess = strings.ToUpper(guess)
			}

			guessNum, ok := noteToNum[guess]
			if !ok {
				statusMessage = guess + " is not a valid note"
				drawNote(x, y, false)
			} else if guessNum == note {
				statusMessage = "Correct!"
				drawNote(x, y, true)
				time.Sleep(time.Second)
				correct = true
			} else {
				statusMessage = guess + " is incorrect"
				drawNote(x, y, false)
			}
		}

	}
}

func drawNote(x int, y int, drawLetter bool) int {
	clearScreen()
	fretboard := ""
	drawnNote := -1
	for xp, note := range tuning {
		guitarString := note + "\t---"
		currNote := noteToNum[note]
		for yp := range fretLength {

			if currNote == 11 {
				currNote = 0
			} else {
				currNote += 1
			}

			noteStr := numToNote[currNote]
			if drawLetter == false {
				noteStr = "?"
			}
			if x == xp && y == yp {
				drawnNote = currNote
				guitarString += "| " + noteStr + ""
				if len(noteStr) == 1 {
					guitarString += " "
				}
			} else {
				guitarString += "| - "
			}

		}
		guitarString += "|---"
		fretboard += guitarString + "\n"
	}
	if fretLength >= 12 {
		fretboard += `                     3       5       7       9          12`
	} else if fretLength >= 9 {
		fretboard += `                     3       5       7       9`
	} else if fretLength >= 7 {
		fretboard += `                     3       5       7`
	} else if fretLength >= 5 {
		fretboard += `                     3       5`
	} else if fretLength >= 3 {
		fretboard += `                     3`
	}
	fmt.Println(fretboard + "")
	fmt.Println(statusMessage)
	return drawnNote
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func collectInput() {

	tuningInput := flag.String("tuning", "EADGBE", "Guitar tuning and amount of strings")
	fretInput := flag.Int("frets", 12, "Amount of frets")
	flag.Parse()

	tuningValue := *tuningInput
	fretLength = *fretInput

	// Collect tuningInput
	tuningValue = strings.ToUpper(tuningValue)

	// Transform into slice with regex
	// A#GBbZ --> ["A#", "G", "Bb"]
	re := regexp.MustCompile(`[A-G][#b]|[A-G]`)
	notes := re.FindAllString(tuningValue, -1)

	// Verify tuningInput
	if len(notes) == 0 {
		panic("Invalid tuningInput")
	}
	for _, note := range notes {
		_, ok := noteToNum[note]
		if !ok {
			panic(note + " is not a valid note")
		}
	}

	slices.Reverse(notes)
	tuning = notes
}

var noteToNum = map[string]int{
	"C":  0,
	"C#": 1,
	"Db": 1,
	"D":  2,
	"D#": 3,
	"Eb": 3,
	"E":  4,
	"F":  5,
	"F#": 6,
	"Gb": 6,
	"G":  7,
	"G#": 8,
	"Ab": 8,
	"A":  9,
	"A#": 10,
	"Bb": 10,
	"B":  11,
}

var numToNote = map[int]string{
	0: "C",
	1: "C#",
	//	1: "Db",
	2: "D",
	3: "D#",
	//	3: "Eb",
	4: "E",
	5: "F",
	6: "F#",
	//	6: "Gb",
	7: "G",
	8: "G#",
	//	8: "Ab",
	9:  "A",
	10: "A#",
	//	0: "Bb",
	11: "B",
}
