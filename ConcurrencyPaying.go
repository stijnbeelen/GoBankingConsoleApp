package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {

	//waitgroup
	//go pay(&accountStijn, &accountKoen, 50, &wg)
	//go pay(&accountKoen, &accountStijn, 15, &wg)
	//go deposit(&accountStijn, 5, &wg)
	//go withdraw(&accountKoen, 35, &wg)
	//
	//// -> uitkomst: koen = 100
	//// -> uitkomst: stijn = 70
	//
	//wg.Wait()

	// for using Locks instead of waitgroup
	//time.Sleep(time.Second)

	accDepChan := make(chan *BankAccount) //account deposit chan
	accWitChan := make(chan *BankAccount) //account withdraw chan

	amDepChan := make(chan int) //amount deposit chan
	amWitChan := make(chan int) //amount withdraw chan

	go deposit(accDepChan, amDepChan)
	go withdraw(accWitChan, amWitChan)

	var choice, amount int

	for {
		fmt.Println("1) Deposit")
		fmt.Println("2) Withdraw")
		fmt.Println("3) Exit")
		fmt.Print("Maak uw keuze: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Bedrag: ")
			fmt.Scan(&amount)
			accDepChan <- &accountStijn
			amDepChan <- amount
			fmt.Println(accountStijn)
		case 2:
			fmt.Print("Bedrag: ")
			fmt.Scan(&amount)
			accWitChan <- &accountStijn
			amWitChan <- amount
			fmt.Println(accountStijn)
		case 3:
			os.Exit(0)
		}

	}
}

func deposit(accChan chan *BankAccount, amChan chan int) {
	for true {
		ba := <-accChan
		ba.balance += <-amChan
	}
}

func withdraw(accChan chan *BankAccount, amChan chan int) {
	for true {
		ba := <-accChan
		ba.balance -= <-amChan
	}
}

func pay(payerAccount, receiverAccount *BankAccount, amount int, wg * sync.WaitGroup) {
	wg.Add(2)
	go payerAccount.pay(receiverAccount, amount, wg)
}

//func deposit(account *BankAccount, amount int, wg *sync.WaitGroup) {
//	wg.Add(1)
//	go account.deposit(amount, wg)
//}

//func withdraw(account *BankAccount, amount int, wg *sync.WaitGroup) {
//	wg.Add(1)
//	go account.withdraw(amount, wg)
//}
