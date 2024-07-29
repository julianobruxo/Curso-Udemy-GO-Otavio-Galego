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
func login() {
	fmt.Println("Welcome")
	var name string
	time.Sleep(time.Second * 1)
	fmt.Println("Type your name")
	for {
		fmt.Scan(&name)
		if name == "Juliano" {
			fmt.Println("Welcome master. ")
			break
		} else if name == "Suzy" {
			fmt.Println("Welcome master's master. ")
			break
		} else if name == "Isa" {
			fmt.Println("Welcome young master. ")
			break
		} else {
			fmt.Println("Invalid name.\nTry again")
			continue
		}
	}
}
func main() {
	login()
	password()
}
