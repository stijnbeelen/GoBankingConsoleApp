package main

import (
	"fmt"
	"os"
)

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

func main() {

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

	accountDepositChannel := make(chan *BankAccount)  	//account deposit chan
	accountWithdrawChannel := make(chan *BankAccount) 	//account withdraw chan

	amountDepositChannel := make(chan int)  			//amount deposit chan
	amountWithdrawChannel := make(chan int) 			//amount withdraw chan

	go deposit(accountDepositChannel, amountDepositChannel)
	go withdraw(accountWithdrawChannel, amountWithdrawChannel)

	var username, password string
	username, password = login()

	var choice, amount int

	if accountKoen.holder == username && accountKoen.password == password || accountStijn.holder == username && accountStijn.password == password {

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
				os.Exit(0)
			}

		}
	}

}

func login() (username, password string) {

	fmt.Print("Gebruikersnaam: ")
	fmt.Scan(&username)
	fmt.Print("Wachtwoord: ")
	fmt.Scan(&password)

	return
	}
