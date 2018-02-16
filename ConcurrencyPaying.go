package main

import "fmt"



func main() {

	go accountStijn.deposit(50)
	go accountStijn.withdraw(50)
	go accountStijn.deposit(25)
	go accountKoen.withdraw(50)

	fmt.Println(accountStijn, accountKoen)
}
