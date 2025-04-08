package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/kremlsa/digrecon/internal/simple"
)

var menu = map[string]func(){
	"1": simple.Recon,
}

func main() {
	counter := menuCounter

Menu:
	for {
		counter()
		operation := outputData(
			"Choose recognition operation:",
			"1. Simple recognition",
			"0. Exit\n",
			"Your choice",
		)
		menuFunc := menu[operation]
		if menuFunc == nil {
			break Menu
		}
		menuFunc()
	}
}

func outputData(data ...any) string {
	for i, line := range data {
		if i == len(data)-1 {
			fmt.Printf("%v ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

func menuCounter() func() {
	i := 0
	return func() {
		i++
		color.Green("%d", i)
	}
}
