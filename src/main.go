package main

import (
	"fmt"
	"go-workspace/src/pattern/factory"
	"go-workspace/src/pattern/singleton"
)

func main() {
	fmt.Println("HI")
	ak47, _ := factory.GetGun("ak47")
	factory.PrintDetails(ak47)

	for i := 0; i < 30; i++ {
		go singleton.GetInstance()
	}
}
