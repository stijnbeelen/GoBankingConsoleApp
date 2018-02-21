package main

import "sync"

type BankAccount struct {
	accountNumber int
	holder string
	balance int
	password string
}

func (bankAccount *BankAccount) withdraw(amount int, wg *sync.WaitGroup) {
	mux.Lock()
	defer mux.Unlock()
	defer wg.Done()
	bankAccount.balance -= amount
}

func (bankAccount *BankAccount) deposit(amount int, wg *sync.WaitGroup) {
	mux.Lock()
	defer mux.Unlock()
	defer wg.Done()

	bankAccount.balance +=  amount
}

func (payer *BankAccount) pay (receiver *BankAccount, amount int, wg *sync.WaitGroup) {
	payer.withdraw(amount, wg)
	receiver.deposit(amount, wg)
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

//func pay(payerAccount, receiverAccount *BankAccount, amount int, wg * sync.WaitGroup) {
//	wg.Add(2)
//	go payerAccount.pay(receiverAccount, amount, wg)
//}

//func deposit(account *BankAccount, amount int, wg *sync.WaitGroup) {
//	wg.Add(1)
//	go account.deposit(amount, wg)
//}

//func withdraw(account *BankAccount, amount int, wg *sync.WaitGroup) {
//	wg.Add(1)
//	go account.withdraw(amount, wg)
//}
