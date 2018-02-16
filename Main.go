package main

import (
	"fmt"
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

func login()(string,string){

	var username, password string

	fmt.Print("Gebruikersnaam: ")
	fmt.Scan(&username)
	fmt.Print("Wachtwoord: ")
	fmt.Scan(&password)

	return username,password
}