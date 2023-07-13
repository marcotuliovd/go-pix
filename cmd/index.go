package cmd

import (
	"fmt"
	"pix/internal"
)

func Testando() int {
	return 2
}

func Index(amount int) int {
	tere := internal.AddAmountInBalance(amount)
	fmt.Println(tere)
	return tere
}