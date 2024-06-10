package ascii

import (
	"fmt"
	"os"
	"strings"
)

// Ascii prints ASCII art from a given array of characters.
// The characters are extracted from a predefined file array.
// The function takes in four arguments: fileArr (a slice of strings representing the file array),
// wordsArr (a slice of strings representing the words to be printed),
// lettersToColor (a string representing the letters to be colored),
// and color (a string representing the color to be applied).
func Ascii(fileArr []string, wordsArr []string, lettersToColor string, colorCode string, outputflag string) {
	var count int
	reset := "\033[0m"

	for _, val := range wordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					if len(lettersToColor) == 0 {
						if IsFlagPassed("output") {
							err := os.WriteFile(outputflag, []byte(fileArr[int(start)+i]), 0644)
							if err != nil {
								fmt.Println(err)
								return
							}
						} else {
							fmt.Print(colorCode + fileArr[int(start)+i] + reset)
						}

					} else if strings.Contains(lettersToColor, string(v)) {
						if IsFlagPassed("output") {
							err := os.WriteFile(outputflag, []byte(fileArr[int(start)+i]), 0644)
							if err != nil {
								fmt.Println(err)
								return
							}
						} else {
							fmt.Print(colorCode + fileArr[int(start)+i] + reset)
						}
					} else {
						if IsFlagPassed("output") {
							err := os.WriteFile(outputflag, []byte(fileArr[int(start)+i]), 0644)
							if err != nil {
								fmt.Println(err)
								return
							}
						} else {
							fmt.Print(fileArr[int(start)+i])
						}
					}
				}

				if IsFlagPassed("output") {
					err := os.WriteFile(outputflag, []byte("\n"), 0644)
					if err != nil {
						fmt.Println(err)
						return
					}
				} else {
					fmt.Println()
				}
			}
		} else {
			count++
			if count < len(wordsArr) {
				if IsFlagPassed("output") {
					err := os.WriteFile(outputflag, []byte("\n"), 0644)
					if err != nil {
						fmt.Println(err)
						return
					}
				} else {
					fmt.Println()
				}
			}
		}
	}
}
