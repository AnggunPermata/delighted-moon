package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findDissimilarity("abc", "abcd"))//d
	fmt.Println(findDissimilarity("bxcn", "abncx")) // a
	fmt.Println(findDissimilarity("", "y")) // y
	fmt.Println(findDissimilarity("annqalff", "fqlnannaf")) // n
}

func findDissimilarity(s, t string) string{
	mapS := mappingChar(s)
	mapT := mappingChar(t)
	start := 97
	end := 122
	result := ""

	for i := start; i <= end; i++{
		char := string(i)
		if valT, okT := mapT[char]; okT {
			if valS, okS := mapS[char]; okS {
				minus := valT - valS
				if minus > 0 {
					for j := 0; j < minus; j++ {
						result += char
					}
				}
			} else if !okS {
				loopsIfExist := valT
				for j := 0; j < loopsIfExist; j++ {
					result += char
				}
			}
		}
	}
	return result
}

func mappingChar(input string) map[string]int{
	arr := strings.Split(strings.ToLower(input), "")
	mapS := make(map[string]int)
	for i:=0;i<len(arr);i++{
		char := arr[i]
		var val int
		var ok bool
		if val, ok = mapS[char]; ok {
			mapS[char] = val + 1
		} else {
			mapS[char] = 1
		}
	}
	return mapS
}