package main

import (
	"fmt"
	"go-workspace/src/pattern/factory"
)

func main() {
	fmt.Println("HI")
	ak47, _ := factory.GetGun("ak47")
	factory.PrintDetails(ak47)
}
