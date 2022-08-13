package main

import "fmt"

func main() {
	var kilolar [4]int = [4]int{80, 65, 129, 135}
	tartiIslemi(kilolar)
}

func tartiIslemi(kilolar [4]int) {
	var tartiMaks int = 120
	uyan := 0
	uymayan := 0
	for x := 0; x < len(kilolar); x++ {

		if kilolar[x] <= tartiMaks {
			fmt.Println(x, ". kişi", kilolar[x], "kilo")
			uyan++
		} else {
			fazlalik := kilolar[x] - tartiMaks
			fmt.Println(x, ". kişi", kilolar[x], "kilo", fazlalik, "kilo fazlan ile listeye eklenmedin")
			uymayan++
		}
	}
	fmt.Println("120 kilo altı tartılan kişi", uyan)
	fmt.Println("120 kilo fazlazı tartılan kişi", uymayan)
}
