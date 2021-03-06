package main

import (
	"fmt"
	"os"
)

func main() {
	accountDepositChannel := make(chan *BankAccount)  //account deposit chan
	accountWithdrawChannel := make(chan *BankAccount) //account withdraw chan

	amountDepositChannel := make(chan int)  //amount deposit chan
	amountWithdrawChannel := make(chan int) //amount withdraw chan

	go deposit(accountDepositChannel, amountDepositChannel)
	go withdraw(accountWithdrawChannel, amountWithdrawChannel)

	var username, password string
	username, password = login()

	var choice, amount int

	if accountKoen.holder == username && accountKoen.password == password || accountStijn.holder == username && accountStijn.password == password {

		for {
			fmt.Println("1) Zet geld op je rekening")
			fmt.Println("2) Haal geld af")
			fmt.Println("3) Schrijf over")
			fmt.Println("4) Sluit af")
			fmt.Print("Maak uw keuze: ")
			fmt.Scan(&choice)

			switch choice {
			case 1:
				fmt.Print("Bedrag: ")
				fmt.Scan(&amount)

				switch username {
				case "Stijn":
					accountDepositChannel <- &accountStijn
				case "Koen":
					accountDepositChannel <- &accountKoen
				}

				amountDepositChannel <- amount
				fmt.Println(accountStijn, accountKoen)
			case 2:
				fmt.Print("Bedrag: ")
				fmt.Scan(&amount)

				switch username {
				case "Stijn":
					accountWithdrawChannel <- &accountStijn
				case "Koen":
					accountWithdrawChannel <- &accountKoen
				}

				amountWithdrawChannel <- amount
				fmt.Println(accountStijn, accountKoen)
			case 3:
				fmt.Print("Bedrag: ")
				fmt.Scan(&amount)

				switch username {
				case "Stijn":
					accountWithdrawChannel <- &accountStijn
					accountDepositChannel <- &accountKoen
				case "Koen":
					accountWithdrawChannel <- &accountKoen
					accountDepositChannel <- &accountStijn
				}

				amountDepositChannel <- amount
				amountWithdrawChannel <- amount
				fmt.Println(accountStijn,accountKoen)
			case 4:
				os.Exit(0)
			}

		}
	}

}
