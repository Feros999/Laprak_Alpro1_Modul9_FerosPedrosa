package main

import "fmt"

func main() {
	var x, y int
	var isFactorXY, isFactorYX bool

	fmt.Print("Masukkan bilangan bulat positif x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan bulat positif y: ")
	fmt.Scan(&y)

	if y%x == 0 {
		isFactorXY = true
	} else {
		isFactorXY = false
	}

	if x%y == 0 {
		isFactorYX = true
	} else {
		isFactorYX = false
	}

	fmt.Println(isFactorXY)
	fmt.Println(isFactorYX)
}
