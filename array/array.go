package main

import "fmt"

func main() {

	// var bahasa [5]string
	// bahasa[0] = "ARAB"
	// bahasa[1] = "INDONESIA"
	// bahasa[2] = "JEPANG"
	// bahasa[3] = "BELGIA"
	// bahasa[4] = "BELANDA"

	bahasa := [...]string{"arab", "belanda", "indonesia"}

	fmt.Println(bahasa)
	panjang := len(bahasa)
	fmt.Println(panjang)
}
