package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/susg/autowordle/internal/config"
	"github.com/susg/autowordle/internal/orchestrator"
	"github.com/susg/autowordle/internal/reader"
	"github.com/susg/autowordle/internal/validate"
	"github.com/susg/autowordle/internal/words"
)

func main() {
	fmt.Println("*******************************************")
	fmt.Println("*************** Auto Wordle ***************")
	fmt.Println("*******************************************")
	appCfg := config.GetConfig()
	wm := words.StartWordManager(reader.NewFileReader(), appCfg)
	for {
	start:
		fmt.Print("\nEnter the word length (enter 'q' to quit): ")
		var wordLengthStr string
		fmt.Scanf("%s", &wordLengthStr)
		wordLengthStr = strings.ToLower(wordLengthStr)
		if strings.ToLower(wordLengthStr) == "q" {
			return
		}

		wordLength, err := strconv.Atoi(wordLengthStr)
		if err != nil {
			fmt.Println("Invalid word length. Please enter a number.")
			continue
		}

		v, err := validate.NewWordleValidator(wordLength, appCfg)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		orch := orchestrator.NewWordleOrchestratorImpl(wordLength, wm, v)
		for {
			var input []string
			fmt.Print("\nEnter the input string (enter 'q' to start new game): ")
			for range wordLength {
				var str string
				fmt.Scanf("%s", &str)
				str = strings.ToLower(str)
				if strings.ToLower(str) == "q" {
					goto start
				}
				input = append(input, str)
			}
			output, err := orch.GenerateWords(input)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			fmt.Println("\nFiltered words: ", output)
		}
	}
}
