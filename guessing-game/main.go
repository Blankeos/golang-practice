package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	var secretNumber = rand.Intn(100)
	fmt.Println("I have generated a random number. Try to guess it!")
	var won = false
	var reader = bufio.NewReader(os.Stdin)

	for {
		if won {
			break
		}

		fmt.Println("--------------------------------------")

		fmt.Printf("Please input your guess: ")

		guess, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read line.")
		}

		parsedGuess, err := strconv.Atoi(strings.TrimSpace(guess))
		if err != nil {
			fmt.Printf("You need to type a number. Found: %s\n", guess)
			continue
		}

		fmt.Printf("You have guessed %s\n", guess)

		// Win
		if parsedGuess < secretNumber {
			fmt.Println("Higher!")
		} else if parsedGuess == secretNumber {
			won = true
			fmt.Println("You have won!")
		} else {
			fmt.Println("Lower!")
		}
	}
}
