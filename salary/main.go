package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

// Colors
var yellow = color.New(color.FgYellow).SprintFunc()
var cyan_underline = color.New(color.FgCyan, color.Bold, color.Underline).SprintFunc()
var bold = color.New(color.Bold)

func main() {
	fmt.Printf("💵 %s to the %s App\n", yellow("Welcome"), cyan_underline("Salary Estimator"))
	fmt.Println("Estimate your yearly salary")
	fmt.Println("")

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	listenForInputs()

}
func printSelections(selections []string, currentSelection int) {
	fmt.Printf("\033[%dA\033[0K", len(selections))

	for i, selection := range selections {
		if i == currentSelection {
			color.Set(color.FgYellow)
			fmt.Printf("👉   %s \t\n", selection)
		} else {
			color.Unset()
			fmt.Printf("     %s \t\n", selection)
		}
	}
	color.Unset()
}

func listenForInputs() {
	var currentSelection int
	fmt.Println("Select (press [^] [v] to choose):")
	selections := []string{"Hourly", "Daily", "Monthly"}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	printSelections(selections, currentSelection)

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyArrowUp {
			if currentSelection == 0 {
			} else {
				currentSelection -= 1
			}
			printSelections(selections, currentSelection)
		} else if key == keyboard.KeyArrowDown {
			if currentSelection == len(selections)-1 {
			} else {
				currentSelection += 1
			}
			printSelections(selections, currentSelection)
		} else if key == keyboard.KeyEnter {
			break
		}
	}

	fmt.Println("")
	fmt.Println("")

	var currencyIdx int = 1
	currencySelections := []string{"USD", "PHP"}
	var salaryInput string = ""

	printCurrencySelect(currencySelections, currencyIdx, selections, currentSelection, salaryInput)

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyArrowLeft {
			if currencyIdx == 0 {

			} else {
				currencyIdx -= 1
			}
		} else if key == keyboard.KeyArrowRight {
			if currencyIdx == len(currencySelections)-1 {

			} else {
				currencyIdx += 1
			}
		} else if char >= '0' && char <= '9' {
			salaryInput += string(char)
		}
		printCurrencySelect(currencySelections, currencyIdx, selections, currentSelection, salaryInput)
	}

}

func printCurrencySelect(currencySelections []string, currentIdx int, timelySelections []string, timelyIdx int, salaryInput string) {
	fmt.Print("\033[2K\033[1F")

	fmt.Print("[<]")
	for i, selection := range currencySelections {
		if i == currentIdx {
			bold.Printf(" %s ", selection)
		} else {
			fmt.Printf(" %s ", selection)
		}
		if i < len(currencySelections)-1 {
			fmt.Printf("|")
		}
	}
	fmt.Print("[>]\n")

	fmt.Print("✨ What's your")
	color.Set(color.FgYellow)
	fmt.Printf(" %s ", timelySelections[timelyIdx])
	color.Unset()
	fmt.Printf("Salary (%s): %s", currencySelections[currentIdx], salaryInput)
}
