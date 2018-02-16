package main

import (
	"fmt"
	"os"
)

var accountStijn = BankAccount{
	001,
	"Stijn",
	100,
	make(map[int]int),
	"Stijn001",
}

var accountKoen = BankAccount{
	002,
	"Koen",
	100,
	make(map[int]int),
	"Koen002",
}

/*func main() {

	var username, password string
	var choice, amount  int

	username, password = login()

	if accountKoen.holder == username && accountKoen.password == password || accountStijn.holder == username && accountStijn.password == password {

		for   {
			fmt.Println("1) Overschrijven")
			fmt.Println("2) Geschiedenis")
			fmt.Println("3) Exit")
			fmt.Print("Maak uw keuze: ")
			fmt.Scan(&choice)

			switch choice {
			case 1:
				fmt.Print("Bedrag: ")
				fmt.Scan(&amount)
				if username == "Koen" {
					accountKoen.pay(&accountStijn, amount)
					fmt.Println(accountStijn, accountKoen)
				}else {
					accountStijn.pay(&accountKoen, amount)
					fmt.Println(accountStijn, accountKoen)
				}
			case 2:
			case 3: os.Exit(0)
			}
			fmt.Println()
		}


	}

}*/

type BankAccount struct {
	accountNumber int
	holder string
	balance int
	history map[int]int
	password string
}

func login()(string,string){

	var username, password string

	fmt.Print("Gebruikersnaam: ")
	fmt.Scan(&username)
	fmt.Print("Wachtwoord: ")
	fmt.Scan(&password)

	return username,password
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