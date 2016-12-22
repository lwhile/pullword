package main

import (
	"fmt"
	"log"
	"pullword"
)

func main() {
	request := pullword.NewRequest("感谢pullword的中文分词服务", false, true)
	resM, err := request.GetM()
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range resM {
		fmt.Println(k, v)
	}
}
