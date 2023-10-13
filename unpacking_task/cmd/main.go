package main

import (
	"fmt"
	"log"

	"unpacking_task/pkg/unpacking"
)

func main() {
	str := ""
	fmt.Scan(&str)
	str, err := unpacking.Unpacking(str)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(str)
}
