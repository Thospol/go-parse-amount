package main

import (
	"fmt"
	"go-validate-amount/amount"
	//"github.com/pkg/errors"
)

func main() {
	amount, err := amount.ParseReqAmountWithoutDecimal("123,456.00")
	if err != nil {
		fmt.Println("parseReqAmount error: ", err)
	}
	fmt.Println("amount: ", amount)
}
