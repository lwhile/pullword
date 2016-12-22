package main

import (
	"fmt"
	"log"
	"pullword"
)

func main() {
	request := pullword.NewRequest("马化腾是李彦宏最大的威胁", false, false)
	resM, err := request.GetM()
	resS, err := request.GetS()
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range resM {
		fmt.Println(k, v)
	}
	fmt.Println()
	for _, v := range resS {
		fmt.Println(v)
	}
}
