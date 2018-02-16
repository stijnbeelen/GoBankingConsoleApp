package main

type BankAccount struct {
	accountNumber int
	holder string
	balance int
	history map[int]int
	password string
}

func (bankAccount *BankAccount) withdraw(amount int) {
	mux.Lock()
	bankAccount.balance -= amount
	mux.Unlock()
}

func (bankAccount *BankAccount) deposit(amount int) {
	mux.Lock()
	bankAccount.balance +=  amount
	mux.Unlock()
}

func (payer *BankAccount) pay (receiver *BankAccount, amount int) {
	payer.withdraw(amount)
	receiver.deposit(amount)
}
