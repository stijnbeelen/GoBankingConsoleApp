package main

type BankAccount struct {
	accountNumber int
	holder string
	balance int
	history map[int]int
	password string
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
