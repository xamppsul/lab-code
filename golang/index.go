package main

import (
	"fmt"
	"golang/basic"
)

func main() {
	var num int
	//array dulu
	basic.Array()
	//condition
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&num) //input dari user

	result := basic.ValidateNumberAvailable(num)
	message := fmt.Sprintf("Angka %d, %s", basic.ExecNumber(num), result)
	fmt.Println(message)
}
