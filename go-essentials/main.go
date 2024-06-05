package main

import (
	"fmt"
	"strings"

	"phpguru.net/go-essentials/helpers"
	"phpguru.net/go-essentials/hero"
	"phpguru.net/go-essentials/persistent"
	"phpguru.net/go-essentials/user"
)

// https://pkg.go.dev/

const FILE_PATH string = "users.json"

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
	var id user.ID
	u, err := user.New(user.ID(id.GenerateUUID()), "I love", "Golang", "")
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

	// test write file
	fileStorage := persistent.New(FILE_PATH)
	data, err := helpers.ParseJSON(u)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = fileStorage.SaveData(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	// test with interface
	var k *hero.Knight = hero.NewKnight()
	fmt.Println(k.GetInformation())
	hero.HeroAttack(k)
	fmt.Println(k.GetInformation())

	var archer *hero.Archer = hero.NewArcher()
	fmt.Println(archer.GetInformation())
	hero.HeroAttack(archer)
	fmt.Println(archer.GetInformation())

	var spearMan *hero.SpearMan = hero.NewSpearMan()
	fmt.Println(spearMan.GetInformation())
	hero.HeroAttack(spearMan)
	fmt.Println(spearMan.GetInformation())

	var swordMan *hero.SwordMan = hero.NewSwordMan()
	fmt.Println(swordMan.GetInformation())
	hero.HeroAttack(swordMan)
	fmt.Println(swordMan.GetInformation())

	// generic example
	var a, b int64
	a = 10
	b = 99
	c := helpers.Add(a, b)
	fmt.Printf("%v + %v = %v", a, b, c)

	var e, f float64
	e = 0.3
	f = 0.2
	g := e - f
	fmt.Printf("%v + %v = %v", e, f, g)
}
