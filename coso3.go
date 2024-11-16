package main

import "fmt"

func main() {
	var bilangan int
	var hasil bool
	fmt.Scan(&bilangan)
	//if a < 0 && a%2 == 0 {
	//	hasil = true
	//}
	hasil = bilangan%2 == 0 && bilangan < 0
	fmt.Print(hasil)
}
