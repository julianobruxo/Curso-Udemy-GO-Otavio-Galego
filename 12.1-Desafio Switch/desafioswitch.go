package main

import (
	"fmt"
	"time"
)

func recognition(user string) string {
	switch {
	case user == "Juliano":
		return "Welcome master. "
	case user == "Suzy":
		return "Welcome master's master."
	case user == "Isa":
		return "Welcome young master"
	default:
		return "Intruder detected"
	}
}

func main() {

	for {
		fmt.Println("Welcome")
		var name string
		fmt.Println("Type your name")
		fmt.Scanln(&name)
		response := recognition(name)
		fmt.Println("Input name: ", response)

		if response == "Intruder detected" {
			fmt.Println("Terminating Program due to unauthorized user.")
			break
		}
		fmt.Println(recognition(name))
		var p int
		fmt.Println("Type your password")
		fmt.Scanln(&p)

		if p == 12345 {
			fmt.Println("Access Granted. You may proceed.")
			time.Sleep(1 * time.Second)
			break

		} else if p != 12345 {
			fmt.Println("Access Denied.\nTry again?(Y/N)")
			time.Sleep(1 * time.Second)

			var answer string
			fmt.Scanln(&answer)

			if answer != "Y" && answer != "y" {
				fmt.Println("Terminating Program")
				break
			} else {

				fmt.Println("Starting over.")
				time.Sleep(1 * time.Second)
				continue
			}
		}

	}
}
