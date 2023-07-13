package cmd

import (
	"fmt"
	"pix/cmd/conta"
)

func CMD(){
	r := Testando()
	cont := conta.Testasas(1, 2)
	fmt.Println(r)
	fmt.Println(cont)
}