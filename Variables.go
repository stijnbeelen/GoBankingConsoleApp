package main

import "sync"

var accountStijn = BankAccount{
	001,
	"Stijn",
	100,
	"Stijn001",
}

var accountKoen = BankAccount{
	002,
	"Koen",
	100,
	"Koen002",
}

var mux = sync.Mutex{}
var wg = sync.WaitGroup{}
