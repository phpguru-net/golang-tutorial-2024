package main

import (
	"fmt"
	"strings"

	"phpguru.net/go-essentials/user"
)

// https://pkg.go.dev/

func main() {
	fmt.Println("This is the main entry point of your go application!")

	// I need to test something related to pointer and address
	expectedIncome := 115
	var expectedIncomePointer *int = &expectedIncome

	// when you want to get the address of a variable,
	// you must you an ampersand symbol before your variable
	fmt.Println(expectedIncomePointer)
	// when you want to get the value of your pointer,
	// you must you an asterisk
	fmt.Println(*expectedIncomePointer)

	// I need to test something related to struct
	u, err := user.New(1, "I love", "Golang", "")
	if err != nil {
		fmt.Println(err)
		//panic(err)
		return
	}
	fmt.Println(u.DisplayUser())
	u.SetPrize("Gold")

	fmt.Println(strings.Repeat("*", 20))

	winner := user.NewWinner("phpguru.net@sonnm.com", u)
	fmt.Println(winner.DisplayUser())
}
