package main

import (
	"fmt"
	"strings"
)

var mapRoman = map[string]int{
	"I":1,
	"V" : 5,
	"X" : 10,
	"L" : 50,
	"C" : 100,
	"D" : 500,
	"M" : 1000,
}

func main() {
	fmt.Println(ChangeRomanIntoInt("XII")) //12
	fmt.Println(ChangeRomanIntoInt("LIV")) //54
	fmt.Println(ChangeRomanIntoInt("MMXXII")) //2022
}

func ChangeRomanIntoInt(input string) int{
	arrString := strings.Split(input, "")
	var result int
	for i := 0; i < len(arrString); i++{
		if arrString[i] == "I" && i+1 < len(arrString) {
			if arrString[i+1] == "V" {
				result += 4
				i++
				continue
			} else if arrString[i+1] == "X" {
				result += 9
				i++
				continue
			}
		}
		if arrString[i] == "I" && i == (len(arrString)-1) {
			result += mapRoman[arrString[i]]
			return result
		}
		if arrString[i] == "X" && i+1 < len(arrString){
			if arrString[i+1] == "L" {
				result += 40
				i++
				continue
			} else if arrString[i+1] == "C" {
				result += 90
				i++
				continue
			}
		}
		if arrString[i] == "X" && i == (len(arrString)-1) {
			result += mapRoman[arrString[i]]
			return result
		}
		if arrString[i] == "C" && i+1 < len(arrString){
			if arrString[i+1] == "D" {
				result += 400
				i++
				continue
			} else if arrString[i+1] == "M" {
				result += 900
				i++
				continue
			}
		}
		if arrString[i] == "C" && i == (len(arrString)-1) {
			result += mapRoman[arrString[i]]
			return result
		}
		result += mapRoman[arrString[i]]
	}
	return result
}