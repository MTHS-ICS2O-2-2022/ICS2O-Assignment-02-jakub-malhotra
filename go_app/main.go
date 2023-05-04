// Created by: Jakub Malhotra
// Created on: April 2023
//
// This program calculates the area of a parallelogram

package main

import "fmt"

func main() {
	// This function calculates the area of a parallelogram
	var base int
	var height int
	var area int

	fmt.Println("This program finds the area of a parallelogram.")
	fmt.Println()
	fmt.Print("Enter the base (cm): ")
	fmt.Scanln(&base)
	fmt.Print("Enter the height (cm): ")
	fmt.Scanln(&height)

	area = (base * height)

	fmt.Println()
	fmt.Println("The area is:", area, "cm2.")
	fmt.Println("\nDone.")
}
