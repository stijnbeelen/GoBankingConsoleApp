package main

import "sync"

type BankAccount struct {
	accountNumber int
	holder string
	balance int
	history map[int]int
	password string
}

func (bankAccount *BankAccount) withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	bankAccount.balance -= amount
}

func (bankAccount *BankAccount) deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	bankAccount.balance +=  amount
}

func (payer *BankAccount) pay (receiver *BankAccount, amount int, wg *sync.WaitGroup) {
	payer.withdraw(amount, wg)
	receiver.deposit(amount, wg)
}
