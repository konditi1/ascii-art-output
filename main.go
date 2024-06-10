package main

import (
	"ascii/ascii"
	"flag"
	"fmt"
	//"os"
)

// main function is used to read the banner file and print the ascii art
// based on the arguments passed
func main() {
	// for _, v := range os.Args {
	// 	if v == "--color" || v == "-color" {
	// 		fmt.Println("PPPP","Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
	// 		return
	// 	}
	// }
	var filename string
	var colorflag string
	var outputflag string
	var lettersTocolor string
	var colorCode string

	flag.StringVar(&filename, "filename", "standard", "name for the files")
	flag.StringVar(&colorflag, "color", "reset", "color for color input")
	flag.StringVar(&outputflag, "output", "banner.txt", "name of the file to store the ascii art")
	flag.Parse()
	argsPassed := flag.Args()

	bannerContent, err := ascii.GetFileName(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	if ascii.IsFlagPassed("color") {
		if len(argsPassed) == 1 {
			ascii.Color(colorflag, lettersTocolor, argsPassed, bannerContent, outputflag)
		} else if len(argsPassed) == 2 && (argsPassed[1] == "standard" || argsPassed[1] == "shadow" || argsPassed[1] == "thinkertoy") {
			bannerContent, _ := ascii.GetFileName(argsPassed[1])
			argsPassed = argsPassed[:1]
			ascii.Color(colorflag, lettersTocolor, argsPassed, bannerContent, outputflag)
		} else if len(argsPassed) == 2 && (argsPassed[1] == "\\standard" || argsPassed[1] == "\\shadow" || argsPassed[1] == "\\thinkertoy") {
			argsPassed[1] = argsPassed[1][1:]
			ascii.Color(colorflag, lettersTocolor, argsPassed, bannerContent, outputflag)
		} else if len(argsPassed) == 3 {
			bannerContent, err := ascii.GetFileName(argsPassed[2])
			if err != nil {
				fmt.Println(err)
				return
			}
			argsPassed = argsPassed[:2]
			ascii.Color(colorflag, lettersTocolor, argsPassed, bannerContent, outputflag)
		} else if len(argsPassed) == 2 {
			ascii.Color(colorflag, lettersTocolor, argsPassed, bannerContent, outputflag)
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		}

		// If flag color is not passed
	} else {
		if len(argsPassed) == 2 {
			bannerContent, err := ascii.GetFileName(argsPassed[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			argsPassed = argsPassed[:1]
			colorCode = ""
			ascii.Art(argsPassed, bannerContent, lettersTocolor, colorCode, 0, outputflag)
		} else if len(argsPassed) == 1 {
			colorCode = ""
			ascii.Art(argsPassed, bannerContent, lettersTocolor, colorCode, 0, outputflag)
		} else {
			fmt.Println("PPPP", "Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		}
	}
}
