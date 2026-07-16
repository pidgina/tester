package main

import "fmt"

func JoinStr(text1, text2 string) string {
	return text1 + text2
}

func SummInteger(num1, num2 int) int {
	return num1 + num2
}

func main() {
	text1 := "Первое"
	text2 := "Второе"
	num1 := 123
	num2 := 321

	fmt.Printf("Склеивая строку %s и %s получится: %s\n", text1, text2, JoinStr(text1, text2))
	fmt.Printf("Складывая число %d и %d получится: %d\n", num1, num2, SummInteger(num1, num2))

}
