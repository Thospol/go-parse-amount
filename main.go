package main

import (
	"fmt"
	"go-validate-amount/amount"
	//"github.com/pkg/errors"
)

func main() {
	amount, err := amount.ParseReqAmount(",00.00")
	if err != nil {
		fmt.Println("parseReqAmount error: ", err)
	}
	fmt.Println("amount: ", amount)
}
