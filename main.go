package main

import (
	"fmt"
	"time"
)

func main() {
	calculator()
}

func calculator() {
	//simple calculator
	fmt.Print("\n****** CALCULATOR APP ******\n Add (1)\n Substract (2)\n Multiply (3)\n Divide (4)\n Please enter the option:")
	var s int
	fmt.Scanln(&s)

	switch s {
	case 1:
		fmt.Println("Perfoming Addition")
		fmt.Println("Enter value 1:")
		var n1 int
		fmt.Scanln(&n1)

		fmt.Println("Enter value 2:")
		var n2 int
		fmt.Scanln(&n2)

		sum := n1 + n2
		fmt.Printf("The sum of %d and %d is %d \n", n1, n2, sum)

		time.Sleep(3 * time.Second)
		calculator()

	case 2:
		fmt.Println("Perfoming Substraction")
		fmt.Println("Enter value 1:")
		var n1 int
		fmt.Scanln(&n1)

		fmt.Println("Enter value 2:")
		var n2 int
		fmt.Scanln(&n2)

		sub := n1 - n2
		fmt.Printf("The substract of %d and %d is %d \n", n1, n2, sub)

		time.Sleep(3 * time.Second)
		calculator()

	case 3:
		fmt.Println("Perfoming Multiplication")
		fmt.Println("Enter value 1:")
		var n1 int
		fmt.Scanln(&n1)

		fmt.Println("Enter value 2:")
		var n2 int
		fmt.Scanln(&n2)

		mul := n1 * n2
		fmt.Printf("The multiplication of %d and %d is %d \n", n1, n2, mul)

		time.Sleep(3 * time.Second)
		calculator()

	case 4:
		fmt.Println("Perfoming Division")
		fmt.Println("Enter value 1:")
		var n1 int
		fmt.Scanln(&n1)

		fmt.Println("Enter value 2:")
		var n2 int
		fmt.Scanln(&n2)

		div := n1 / n2
		fmt.Printf("The division of %d and %d is %d \n", n1, n2, div)

		time.Sleep(3 * time.Second)
		calculator()

	default:
		fmt.Println("Invalid selection. Please try again")
		calculator()
	}
}
