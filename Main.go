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

	payChan := make(chan BankAccount)
	recChan := make(chan BankAccount)
	returnChan := make(chan BankAccount)

	payChan <- accountKoen
	recChan <- accountStijn

	pay(payChan,recChan, returnChan, 20)

	fmt.Println(<-returnChan)
	fmt.Println(<-returnChan)


}