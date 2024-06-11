package main

import (
	"fmt"
	"strings"

	"phpguru.net/go-essentials/helpers"
	"phpguru.net/go-essentials/hero"
	"phpguru.net/go-essentials/persistent"
	"phpguru.net/go-essentials/product"
	"phpguru.net/go-essentials/user"
)

// https://pkg.go.dev/

const FILE_PATH string = "users.json"

type TechStackMap map[string]float64

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
	fmt.Printf("%v + %v = %v\n", e, f, g)

	// playing with array
	var numbers [4]int64 = [4]int64{1, 2, 3, 4}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%v\n", numbers[i])
	}
	// for i := 0; i < len(numbers); i++ {
	// 	numbers[i] *= 2
	// }
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%v\n", numbers[i])
	}
	fmt.Println(numbers[0:3]) // [1,2,3]
	fmt.Println(numbers[1:])  // [2,3,4]
	fmt.Println(numbers[4:])  // []
	fmt.Println(cap(numbers))

	// create dynamic lists with slice
	goldPrices := []float64{1.5, 2.55}
	fmt.Println(goldPrices)
	goldPrices = append(goldPrices, 1)
	goldPrices = append(goldPrices, 2)
	for i := 0; i < len(goldPrices); i++ {
		goldPrices[i] += 0.1
	}
	fmt.Println(goldPrices)

	// create array of Product
	products := []*product.Product{
		{"1", "Mouse 1", 1.51},
	}
	products = append(products,
		product.New("2", "Mouse 2", 1.52),
		product.New("3", "Mouse 3", 1.53),
	)

	for i := 0; i < len(products); i++ {
		fmt.Printf("%v\n", products[i].Title)
	}

	discountProducts := []*product.Product{
		{"4", "Mouse 4", 1.54},
		{"5", "Mouse 5", 1.55},
	}

	products = append(products, discountProducts...)
	for i := 0; i < len(products); i++ {
		fmt.Printf("%v\n", products[i].Title)
	}

	// map
	matrix := map[string]string{}
	matrix["nextjsvietnam"] = "https://nextjsvietnam.com"
	matrix["phpguru"] = "https://phpguru.net"

	websites := []string{"nextjsvietnam", "phpguru"}
	for i := 0; i < len(websites); i++ {
		fmt.Println(matrix[websites[i]])
	}
	delete(matrix, "nextjsvietnam")

	// make
	students := make([]string, 5, 6)
	students[0] = "A"
	students[1] = "B"
	students[2] = "C"
	students[3] = "D"
	students[4] = "E"
	students = append(students, "F")
	fmt.Println(students)

	// optimize memory for map with make
	ratings := make(TechStackMap, 3)

	ratings["nextjs"] = 8.5
	ratings["wordpress"] = 9.75

	fmt.Println(ratings)

	for index, value := range students {
		fmt.Printf("%v:%v\n", index, value)
	}

	fmt.Println(strings.Repeat("*", 10))

	for key, value := range ratings {
		fmt.Printf("%v:%v\n", key, value)
	}

	/*
	* Functions:
	* - Using functions as values
	* - Anonymous Functions
	* - Recursion
	* - Variadic
	 */
	nums := make([]int64, len(numbers))
	for i := 0; i < len(numbers); i++ {
		nums[i] = numbers[i]
	}
	transform2 := getTransformerFn("double")
	transform3 := getTransformerFn("tripple")

	dnumbers := transformNumber(&nums, transform2)
	tnumbers := transformNumber(&nums, transform3)
	quadrupleNumbers := transformNumber(&nums, func(n int64) int64 {
		return n * 4
	})
	quintupleNumbers := transformNumber(&nums, createTransformerFn(5))
	fmt.Println(*dnumbers)
	fmt.Println(*tnumbers)
	fmt.Println(*transformNumber(&nums, createTransformerFn(6)))
	fmt.Println(*quadrupleNumbers)
	fmt.Println(*quintupleNumbers)
	fmt.Println(strings.Repeat("*", 20))
	fmt.Println(*transformNumber(&nums, func(n int64) int64 {
		return factorial(n)
	}))
	var cond Condition = func(n int64) bool {
		return n%2 == 0
	}
	fmt.Println(sumif(cond, 1, 2, 3, 4, 5, 6))
	fmt.Println(sumif(cond, *dnumbers...))

	staticArray := [9]any{1, 2, 3, 4, 5, 6, "a", "b", "c"}
	slicesArray := staticArray[:]
	slicesArray = append(slicesArray, "d", "e", "f")
	fmt.Println(slicesArray)
}

type transformFn func(int64) int64

func getTransformerFn(transformType string) transformFn {
	if transformType == "tripple" {
		return tripple
	}
	return double
}

func createTransformerFn(factor int64) transformFn {
	return func(n int64) int64 {
		return n * factor
	}
}

func transformNumber(numbers *[]int64, transform transformFn) *[]int64 {
	dNumbers := make([]int64, len(*numbers))
	for index, value := range *numbers {
		dNumbers[index] = transform(value)
	}
	return &dNumbers
}

func double(n int64) int64 {
	return n * 2
}

func tripple(n int64) int64 {
	return n * 3
}

// recursion

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

type Condition func(n int64) bool

func sumif(condition Condition, numbers ...int64) int64 {
	var total int64
	for _, value := range numbers {
		if condition(value) {
			total += value
		}
	}
	return int64(total)
}
