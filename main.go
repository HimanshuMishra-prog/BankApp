package main

import (
	"BankApp/bank"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello")
	s := uuid.New()
	fmt.Println(s)
	i := bank.Sum(8, 10)
	fmt.Print(i)
}
