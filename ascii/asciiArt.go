package ascii

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// Ascii prints ASCII art from a given array of characters.
func Ascii(fileArr []string, wordsArr []string, lettersToColor string, colorCode string, outputflag string, alignflag string) {
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
			return
		}
	} else if IsFlagPassed("align") {
		if alignflag == "right" {
			AlignRight(outputBuilder.String())
		} else if alignflag == "left" {
			fmt.Print(outputBuilder.String())
		} else if alignflag == "center" {
			AlineCenter(outputBuilder.String())
		} else if alignflag == "justify" {
			AlignJustify(outputBuilder.String(), wordsArr)
		}
	} else {
		fmt.Print(outputBuilder.String())
	}
}

func AlignRight(word string) {
	lines := strings.Split(word, "\n")
	var fullArt []string

	for _, line := range lines {
		// Calculate the space needed for each line individually
		lineWidth := GetArtWidth(line)
		space := TerminalWidth() - lineWidth
		if space > 0 {
			fullArt = append(fullArt, strings.Repeat(" ", space)+line)
		} else {
			fullArt = append(fullArt, line)
		}
	}
	fmt.Println(strings.Join(fullArt, "\n"))
}

func AlineCenter(word string) {
	lines := strings.Split(word, "\n")
	var fullArt []string

	for _, line := range lines {
		space := (TerminalWidth() - GetArtWidth(line)) / 2
		if space > 0 {
			fullArt = append(fullArt, strings.Repeat(" ", space)+line)
		} else {
			fullArt = append(fullArt, line)
		}
	}
	fmt.Println(strings.Join(fullArt, "\n"))
}

func AlignJustify(word string, wordsArr []string) {
	str := strings.Join(wordsArr, " ")
	inputSpaces := len(strings.Fields(str)) - 1

	lines := strings.Split(word, "\n")
	var fullArt []string

	fmt.Println(inputSpaces)
	for _, line := range lines {
		if inputSpaces <= 0 {
			fullArt = append(fullArt, line)
		} else {
			space := (TerminalWidth() - GetArtWidth(word)) / inputSpaces
			if space > 0 {
				fullArt = append(fullArt, line+strings.Repeat(" ", space))

			} else {
				fullArt = append(fullArt, line)
			}
		}
	}
	fmt.Println(strings.Join(fullArt, "\n"))
}

func TerminalWidth() int {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error capturing the terminal width", err)
		return 80
	}

	width, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		fmt.Println("Error converting the terminal width to integer", err)
		return 80
	}
	return width
}

func removeANSI(line string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(line, "")
}

func GetArtWidth(str string) int {
	arr := strings.Split(str, "\n")
	max := 0
	for _, line := range arr {
		// Remove ANSI escape codes
		cleanLine := removeANSI(line)
		// fmt.Printf("Line: '%s', Length: %d\n", cleanLine, len(cleanLine))
		if len(cleanLine) > max {
			max = len(cleanLine)
		}
	}
	return max
}
