package main

import (
	"fmt"
	"github.com/laughing_egg/learngo/account"
	"github.com/laughing_egg/learngo/mydict"
)

func main() {
	account := account.NewAccount("Junbeom")
	account.Deposit(100)
	fmt.Println(account)
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
