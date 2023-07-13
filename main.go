package main

import (
	"pix/cmd"
)

func ReturnMeu(amount int) int {
	r := cmd.Index(amount)
	return r
}

func main() {
	ReturnMeu(10)
}