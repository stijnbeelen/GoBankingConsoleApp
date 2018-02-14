package main

func pay(payChan, recChan, returnChan chan BankAccount, amount int) {
	payer := <- payChan
	receiver := <- recChan

	payer.balance = payer.balance-amount
	receiver.balance = receiver.balance+amount

	returnChan <- payer
	returnChan <- receiver
}

func widthdraw() {

}

func deposit() {

}

//git init
//git add README.md
//git commit -m "first commit"
//git remote add origin https://github.com/stijnbeelen/GoBankingConsoleApp.git
//git push -u origin master