package ascii

import (
	"fmt"
	"os"
	"strings"
)

// Ascii prints ASCII art from a given array of characters.
func Ascii(fileArr []string, wordsArr []string, lettersToColor string, colorCode string, outputflag string) {
	var count int
	reset := "\033[0m"
	var outputBuilder strings.Builder

	for _, val := range wordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					if len(lettersToColor) == 0 {
						outputBuilder.WriteString(colorCode + fileArr[int(start)+i] + reset)
					} else if strings.Contains(lettersToColor, string(v)) {
						outputBuilder.WriteString(colorCode + fileArr[int(start)+i] + reset)
					} else {
						outputBuilder.WriteString(fileArr[int(start)+i])
					}
				}
				outputBuilder.WriteString("\n")
			}
		} else {
			count++
			if count < len(wordsArr) {
				outputBuilder.WriteString("\n")
			}
		}
	}

	if IsFlagPassed("output") {
		file, err := os.Create(outputflag)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(outputBuilder.String())
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Print(outputBuilder.String())
	}
}
