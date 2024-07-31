package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {

	var years int
	var amount float64
	annualInterestRate := 4.5

    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Investment Calculator")    
    fmt.Println("Enter the amount to be invested: ")    
    for {
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        if _, err := fmt.Sscanf(input, "%f", &amount); err == nil {
            break
        }
        fmt.Println("Invalid input. Please enter a valid amount:")
    }

    fmt.Println("Enter the annual interest rate (default value is 4.5): ")    
    for {
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        if input == "" {
            break
        }
        if _, err := fmt.Sscanf(input, "%f", &annualInterestRate); err == nil {
            break
        }
        fmt.Println("Invalid input. Please enter a valid annual interest rate:")
    }

    fmt.Println("Enter the number of years: ")    
    for {
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        if _, err := fmt.Sscanf(input, "%d", &years); err == nil {
            break
        }
        fmt.Println("Invalid input. Please enter a valid number of years:")
    }

    // Calculate the future value
    futureValue := amount * math.Pow((1 + annualInterestRate/100), float64(years))
    fmt.Printf("The future value of the investment is $%.2f\n", futureValue)
}