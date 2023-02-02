package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(Valid("059"))
	fmt.Println(Valid("055 444 285"))
}

func Valid(id string) bool {
	fmt.Println(id)
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	fmt.Println(id)

	sumOfNums := 0
	isSecond := false
	for index := len(id) - 1; index >= 0; index-- {
		fmt.Printf(`this number is...%v`, string(id[index]))

		// convert byte to string
		convertedToString := string(id[index])

		// convert string to int
		convertedToInt, err := strconv.Atoi(convertedToString)
		// fmt.Println(err)
		if err != nil {
			return false
		}

		if isSecond {

			// double int
			valueDoubled := convertedToInt * 2
			fmt.Printf(`number is %v and valueDoubled is %v`, convertedToInt, valueDoubled)
			fmt.Println(" ")
			if valueDoubled > 9 {
				sumOfNums += (valueDoubled - 9)
			} else {
				sumOfNums += valueDoubled
			}
			isSecond = !isSecond
		} else {
			sumOfNums += convertedToInt
			isSecond = !isSecond
		}
	}
	// return true if sum is divisible evenly by 10
	fmt.Printf(`the sum is %v\n`, sumOfNums)
	return sumOfNums%10 == 0
}
