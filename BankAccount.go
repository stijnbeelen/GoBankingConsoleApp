package main

type BankAccount struct {
	accountNumber int
	holder string
	balance int
	history map[int]int
}
