package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//wg.Add(5)
	//
	//go accountStijn.pay(&accountKoen, 50, &wg)
	//go accountKoen.pay(&accountStijn, 15, &wg)
	//go accountStijn.deposit(5, &wg)

	pay(&accountStijn, &accountKoen, 50, &wg)
	pay(&accountKoen, &accountStijn, 15, &wg)
	deposit(&accountStijn, 5, &wg)

	wg.Wait()

	fmt.Println(accountStijn)
	fmt.Println(accountKoen)
}

func pay (payerAccount, receiverAccount *BankAccount, amount int, wg *sync.WaitGroup) {
	wg.Add(2)
	go payerAccount.pay(receiverAccount, amount, wg)
}

func deposit(account *BankAccount, amount int, wg *sync.WaitGroup) {
	wg.Add(1)
	go account.deposit(amount, wg)
}

func withdraw(account BankAccount, amount int, wg *sync.WaitGroup) {
	wg.Add(1)
	go account.withdraw(amount, wg)
}


