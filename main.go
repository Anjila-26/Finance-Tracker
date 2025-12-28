package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const logo = `
  ______ _                        _____       
 |  ____(_)                      / ____|      
 | |__   _ _ __   __ _ _ __   ___| |  __  ___  
 |  __| | | '_ \ / _' | '_ \ / _ \ | |_ |/ _ \ 
 | |    | | | | | (_| | | | |  __/ |__| | (_) |
 |_|    |_|_| |_|\__,_|_| |_|\___|\_____|\___/

         • GO • FINANCE • AUTOMATION

`

func main() {
	fmt.Print(logo)

	fmt.Print("Current Money : ")
	scanner := bufio.NewScanner(os.Stdin)

	var money float64
	var err error
	if scanner.Scan() {
		input := scanner.Text()
		money, err = strconv.ParseFloat(strings.TrimSpace(input), 64)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			return
		}

		fmt.Printf("You entered. : $%.2f\n", money)
	}

	fmt.Print("What you wanna do next? : \n")
	fmt.Print("1. Adding.\n")
	fmt.Print("2. Dedecuting.\n")

	options := bufio.NewScanner(os.Stdin)

	options.Scan()

	option := options.Text()
	optionInt, err := strconv.Atoi(strings.TrimSpace(option))

	if err != nil {
		fmt.Println("Invalid option. Please enter a number.")
		return
	}

	fmt.Printf("This is your option: %v\n", optionInt)

	switch optionInt {
	case 1:
		add_cash, err := readAmount("How much money to add ? ")

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			return
		}

		result := add_money(money, add_cash)
		fmt.Printf("Adding money... $%.2f + $%.2f = $%.2f\n", money, add_cash, result)
	case 2:
		subtract, err := readAmount("How much money to deduct ? ")

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			return
		}

		result := deduct_money(money, subtract)
		fmt.Printf("Deducting money... $%.2f - $%.2f = $%.2f\n", money, subtract, result)
	default:
		fmt.Println("Invalid option.")
	}
}

func readAmount(prompt string) (float64, error) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return 0, fmt.Errorf("no input")
	}
	s := strings.TrimSpace(scanner.Text())
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func add_money(x, y float64) float64 {
	return x + y
}

func deduct_money(x, y float64) float64 {
	return x - y
}
