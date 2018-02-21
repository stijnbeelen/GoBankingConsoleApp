package main

import "fmt"

type BankAccount struct {
	accountNumber int
	holder string
	balance int
	password string
}

func deposit(accChan chan *BankAccount, amChan <-chan int) {
	for true {
		ba := <-accChan
		ba.balance += <-amChan
	}
}

func withdraw(accChan chan *BankAccount, amChan <-chan int) {
	for true {
		ba := <-accChan
		ba.balance -= <-amChan
	}
}

func login() (username, password string) {

	fmt.Print("Gebruikersnaam: ")
	fmt.Scan(&username)
	fmt.Print("Wachtwoord: ")
	fmt.Scan(&password)

	return
}