package main

import "fmt"

func main() {

	accountStijn := BankAccount{
		001,
		"Stijn",
		100,
		make(map[int]int),
	}

	accountKoen := BankAccount{
		002,
		"Koen",
		100,
		make(map[int]int),
	}

/*	payChan := make(chan BankAccount)
	recChan := make(chan BankAccount)
	returnChan := make(chan BankAccount)

	payChan <- accountKoen
	recChan <- accountStijn

	pay(payChan,recChan, returnChan, 20)

	fmt.Println(<-returnChan)
	fmt.Println(<-returnChan)*/

	accountKoen.pay(&accountStijn, 50)

	fmt.Println(accountStijn, accountKoen)


}

type BankAccount struct {
	accountNumber int
	holder string
	balance int
	history map[int]int
}

func (bankAccount *BankAccount) withdraw(amount int) {
	bankAccount.balance -= amount
}

func (bankAccount *BankAccount) deposit(amount int) {
	bankAccount.balance +=  amount
}

func (payer *BankAccount) pay (receiver *BankAccount, amount int) {
	payer.withdraw(amount)
	receiver.deposit(amount)
}

func pay(payChan, recChan, returnChan chan BankAccount, amount int) {
	payer := <- payChan
	receiver := <- recChan

	payer.balance = payer.balance-amount
	receiver.balance = receiver.balance+amount

	returnChan <- payer
	returnChan <- receiver
}