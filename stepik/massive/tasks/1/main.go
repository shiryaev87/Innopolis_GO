package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	firstNumber := n / 100
	secondNumber := n % 100 / 10
	thirdNumber := n % 10

	fmt.Print(thirdNumber*100 + secondNumber*10 + firstNumber)
}
