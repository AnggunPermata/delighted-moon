package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println(KulinaFood(15))
	fmt.Println(KulinaFood(12))
	fmt.Println(KulinaFood(1))
	fmt.Println(KulinaFood(3))
}

func KulinaFood(input int)[]string{
	var arr []string
	for i:= 1; i <= input; i++{
		if i % 5 == 0 && i % 3 == 0 {
			arr = append(arr, "Kulina x Food")
		} else if i % 5 == 0{
			arr = append(arr, "Food")
		} else if i % 3 == 0 {
			arr = append(arr, "Kulina")
		} else {
			arr = append(arr, strconv.Itoa(i))
		}
	}
	return arr
}