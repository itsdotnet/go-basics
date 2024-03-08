package main

import (
	"fmt"
	"time"
)

func main() {

	var nigga = "Iskandar Qodirov"

	var age int = 18

	fmt.Println("FIO: "+nigga+"\nMy Age: ", age)

	var now = time.Now()
	birth_str := "2006-10-13"

	birth, _ := time.Parse(birth_str, "2006-10-13 00:00:00")
	all_days := now.Sub(birth)

	print("How many days I waste : ", all_days)
}
