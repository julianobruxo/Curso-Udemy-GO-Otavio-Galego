package main

import (
	"fmt"
	"time"
)

func password() {
	var p int

	fmt.Println("Type your password")
	for {
		fmt.Scan(&p)
		if p == 12345 {
			fmt.Println("Access granted.\nYou may procceed.")
			break
		} else {
			fmt.Println("Wrong password.\nTry again")
			continue
		}
	}
}
func login() string {
	fmt.Println("Welcome")
	var name string
	time.Sleep(time.Second * 1)
	fmt.Println("Type your name")
	for {
		fmt.Scan(&name)
		switch name {
		case "Juliano":
			{
				fmt.Println("Welcome master. ")
				return name
			}
		case "Suzy":
			{
				fmt.Println("Welcome master's master. ")
				return name
			}
		case "Isa":
			{
				fmt.Println("Welcome young master. ")
				return name
			}

		default:
			{
				fmt.Println("Invalid name.\nTry again")
			}

		}
		continue
	}
}

func main() {
	login()
	password()
}
