package main

type BankAccount struct {
	accountNumber int
	holder string
	balance int
	history map[int]int
}

func (bankAccount BankAccount) widthdraw(amount int) {
	bankAccount.balance -= amount
}

func (bankAccount BankAccount) deposit(amount int) {
	bankAccount.balance +=  amount
}