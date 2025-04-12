# Variable
## const
- const name string = "Junbeom"

## var
### 2 way to use 
- (in func) : 
    name := "Junbeom" 
    < automatically infers the type >
- basic :
    var name string = "Junbeom"

# Function
## Have to declare the type of the variable in the function scope

```go
func multiply(a int, b int) int {
	return a * b
}
```

```go
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
```
### values more than one use ... so that it can accept any number of arguments
```go
func repeatMe(words ...string) {
	fmt.Println(words)
}
```
## naked return
### Don't have to declare the type of the variable in the function scope

```go
func lenAndUpper2(name string) (length int, upper string) {
    defer fmt.Println("I'm done") 
	// defer is used to execute a function after the function has finished executing 
	length = len(name)
	upper = strings.ToUpper(name)
	return
}
```

# loop
## for loop (Have only for)
```go
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
```

# if else 
## variable expression
```go
func canIDrink(age int) bool {
	// variable expression
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}
```

# switch
```go
switch {
	case age < 10:
		return false
	case age == 18:
		return true
	case age > 50:
		return true
	}
	return false
```

## switch with variable expression
```go
switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	case 50:
		return true
	}
	return false
```

# Pointer
```go
// & is used to get the address of a variable
	// * is used to dereference a pointer
	a := 2
	b := &a
	*b = 20
	fmt.Println(a, *b)
```

# Arrays and Slices
```go
	names := []string{"Junbeom", "Jin", "Soo", "Hoon", "Jae"}
	names[3] = "Jay"
	names[4] = "Jane"
	names = append(names, "flynn") // append is used to add an element to the slice
```

# Maps
```go
nico := map[string]string{"name": "Junbeom"}
	for _, value := range nico {
		fmt.Println(value)
	}
	fmt.Println(nico)
```

# Struct
```go
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"Kimchi", "Pizza", "Burger"}
	nico := person{name: "Junbeom", age: 23, favFood: favFood}
	fmt.Println(nico.name)
}
```

# Goroutines
``` 
main func does't wait for goroutines
```
```go
func main() {
    c := make(chan string) // create a channel
	people := [2]string{"John", "Doe"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c) // receive from channel
	fmt.Println(<-c)
    // better way
    for i := 0; i < len(people); i++ {
		fmt.Print("waiting for ", i)
		fmt.Println(<-c)
	}
}
func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy" // send to channel
}
```
