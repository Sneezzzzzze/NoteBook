package main

import (
	"fmt"
	"strings"
)

// Have to declare the type of the variable in the function scope
func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// Don't have to declare the type of the variable in the function scope
func lenAndUpper2(name string) (length int, upper string) {
	defer fmt.Println("I'm done") // defer is used to execute a function after the function has finished executing
	length = len(name)
	upper = strings.ToUpper(name)
	return
}

// values more than one use ... so that it can accept any number of arguments
func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	//for i := 0; i < len(numbers); i++ {
	//	fmt.Println(numbers[i])
	//}
	return total
}

func canIDrink(age int) bool {
	// variable expression
	//if koreanAge := age + 2; koreanAge < 18 {
	//	return false
	//}
	//return true
	//switch {
	//case age < 10:
	//	return false
	//case age == 18:
	//	return true
	//case age > 50:
	//	return true
	//}
	//return false
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	case 50:
		return true
	}
	return false
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"Kimchi", "Pizza", "Burger"}
	nico := person{name: "Junbeom", age: 23, favFood: favFood}
	fmt.Println(nico.name)

	//nico := map[string]string{"name": "Junbeom"}
	//for _, value := range nico {
	//	fmt.Println(value)
	//}
	//fmt.Println(nico)

	//names := []string{"Junbeom", "Jin", "Soo", "Hoon", "Jae"}
	//names[3] = "Jay"
	//names[4] = "Jane"
	//names = append(names, "flynn") // append is used to add an element to the slice
	//fmt.Println(names)
	// & is used to get the address of a variable
	// * is used to dereference a pointer
	//a := 2
	//b := &a
	//*b = 20
	//fmt.Println(a, *b)

	//fmt.Println(canIDrink(20))
	//result := superAdd(1, 2, 3, 4, 5, 6)
	//fmt.Println(result)
	//const name string = "Junbeom"
	//// const in Go is a keyword used to declare a constant
	//
	////var name2 string = "Junbeom"
	//// same as
	////name2 := "Junbeom" // Go automatically infers the type of the variable and only work in the function scope
	//totalLenght, _ := lenAndUpper(name)
	//fmt.Println(totalLenght)
	//fmt.Println(lenAndUpper(name))
	//repeatMe("Hello", "World", "Junbeom")
}
