package main

import (
	"fmt"

	"strconv"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

// Colors
var yellow = color.New(color.FgYellow).SprintFunc()
var cyan_underline = color.New(color.FgCyan, color.Bold, color.Underline).SprintFunc()
var bold = color.New(color.Bold, color.FgYellow)

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

	// 1. Pick Hourly, Daily, Monthly
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

	// 2. Pick Currency and Cash
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
		} else if key == keyboard.KeyEnter {
			break
		}

		printCurrencySelect(currencySelections, currencyIdx, selections, currentSelection, salaryInput)
	}

	fmt.Println("")
	fmt.Println("")

	// 3. Calculate
	salary, err := strconv.ParseFloat(salaryInput, 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return
	}
	calculateSalary(currencySelections[currencyIdx], selections[currentSelection], salary)

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

/*
	@param currency USD or PHP

@param timely Hourly, Daily, Monthly
*/
func calculateSalary(currency string, timely string, salary float64) {
	USD_EXCHANGERATE := 55.75
	if currency == "USD" {
		salary *= USD_EXCHANGERATE
	}

	// salary * 5 days a week, 4 weeks a month, 12 months a year
	if timely == "Daily" {
		salary = salary * 5 * 4 * 12
	} else if timely == "Hourly" {
		salary = salary * 8 * 5 * 4 * 12
	} else if timely == "Monthly" {
		salary = salary * 12
	}

	fmt.Printf("You are expected to make PHP %.2f/year", salary)
}
