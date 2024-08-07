package main

import (
	"errors"
	"fmt"
	"time"
)

func calc(n1, n2 float64, op string) (float64, error) {
	switch op {
	case "1":
		return n1 + n2, nil
	case "2":
		return n1 - n2, nil
	case "3":
		return n1 * n2, nil
	case "4":
		if n2 == 0 {
			return 0, errors.New("Cannot divide it by zero")
		}
		return n1 / n2, nil
	default:
		return 0, errors.New("Invalid Operation")
	}
}

func main() {
	fmt.Println("Initiating Basic Calculator 2.0")

	var input1, input2 float64
	var op string
	var cont string

	for {
		fmt.Println("Please select the type of calculation\nyou wish to use:\n1.Sum\n2.Subtraction\n3.Multiplying\n4.Division\n5.Exit")
		fmt.Scanln(&op)
		if op == "5" {
			fmt.Println("Terminating the calculator.")
			time.Sleep(time.Second * 1)
			break
		}

		if op != "1" && op != "2" && op != "3" && op != "4" {
			fmt.Println("Invalid option. Try again.")
			continue
		}
		fmt.Println("Please type the first number:")
		fmt.Scanln(&input1)
		fmt.Println("Please type the second number:")
		fmt.Scanln(&input2)

		result, err := calc(input1, input2, op)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("The result is:", result)
		}

		fmt.Println("Do you want to perform another calculation? (yes/no)")
		fmt.Scanln(&cont)

		if cont != "yes" && cont != "y" {
			fmt.Println("Terminating the calculator.")
			time.Sleep(time.Second * 1)
			break
		}

	}

}
