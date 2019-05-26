package amount

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	regexpJcbAmount = regexp.MustCompile(`(^\d{1,3}(,\d{3})*$)`)
)

func inValidjcbFieldAmount(field string) bool {
	return regexpJcbAmount.MatchString(field)
}

func ParseReqAmount(str string) (int, error) {

	//==========================//
	// Validate  amount has comma
	//==========================//

	// Check Amount comma(,) or not
	if strings.Contains(str, ",") {
		// Validate amount has Prefix 0 so invalid format
		if strings.HasPrefix(str, "0") {
			fmt.Println("The amount is in wrong format")
			return 0, errors.New("9901")
		}

		// Split the amount into whole number and decimal for checking, because we accept 1.00
		parts := strings.Split(str, ".")
		switch {
		case len(parts) == 0, len(parts) > 2:
			fmt.Println("The amount is in wrong format")
			return 0, errors.New("9901")
		case len(parts) == 2:
			if decimal, err := strconv.Atoi(parts[1]); err != nil || decimal != 0 {
				fmt.Println("Numbers with decimal points are not accepted")
				return 0, errors.New("9901")
			}
		}

		if !inValidjcbFieldAmount(parts[0]) {
			fmt.Println("The amount is in wrong format")
			return 0, errors.New("9901")
		}

		// Remove the ',' characters, because some amounts can be in the format 1,000,000
		str := strings.Replace(parts[0], ",", "", -1)

		amount, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("The number is in correct, error: %v", err)
			return 0, errors.New("9901")
		}
		return amount, nil
	}

	//==========================//
	// Validate  amount not comma
	//==========================//

	// Split the amount into whole number and decimal for checking, because we accept 1.00
	parts := strings.Split(str, ".")
	switch {
	case len(parts) == 0, len(parts) > 2:
		fmt.Println("The amount is in wrong format")
		return 0, errors.New("9901")
	case len(parts) == 2:
		if decimal, err := strconv.Atoi(parts[1]); err != nil || decimal != 0 {
			fmt.Println("Numbers with decimal points are not accepted")
			return 0, errors.New("9901")
		}
	}

	// Validate amount case len amount string more than 1 has Prefix 0 so invalid format
	if len(parts[0]) > 1 {
		if strings.HasPrefix(parts[0], "0") {
			fmt.Println("The amount is in wrong format")
			return 0, errors.New("9901")
		}
	}

	amount, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Printf("The number is in correct, error: %v", err)
		return 0, errors.New("9901")
	}

	return amount, nil
}
