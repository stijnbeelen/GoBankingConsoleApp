package main

import (
	"fmt"
	"time"
)

func main() {

	go accountStijn.pay(&accountKoen, 50)
	go accountKoen.pay(&accountStijn, 15)
	go accountStijn.deposit(5)

	time.Sleep(time.Second)

	fmt.Println(accountStijn)
	fmt.Println(accountKoen)
}
