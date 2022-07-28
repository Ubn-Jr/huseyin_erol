package main

import "fmt"

func main() {
	underage(2020)

}
func underage(birth int) {
	if 2022-birth >= 18 {
		fmt.Println("reşit")
	} else {
		fmt.Println("reşit değil")
	}

}
