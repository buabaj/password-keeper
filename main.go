package main

import (
	"fmt"
)

type PasswordKeeper struct {
	website  string
	password string
}

var Database = map[string]string{}

func greet() {
	fmt.Println("Hello, welcome to your password keeper app!")
}

func getAllStoredPasswords() {
	for website, password := range Database {
		fmt.Printf("Your password for %s is: %s \n", website, password)
	}
}

func main() {
	greet()

	var website, password string
	fmt.Println("Please enter the website you want to save the password for:")
	fmt.Scan(&website)
	fmt.Println("Please enter the password for the website:")
	fmt.Scan(&password)

	keeper := PasswordKeeper{website, password}
	fmt.Printf("Your password for %s is: %s", keeper.website, keeper.password)
	Database[keeper.website] = keeper.password

	var choice string
	fmt.Println("\nDo you want to see all your stored passwords? (y/n)")
	fmt.Scan(&choice)
	if choice == "y" {
		getAllStoredPasswords()
		main()
	} else {
		fmt.Println("Thank you for using the password keeper app!")
	}

}
