package main

import (
	"fmt"
)

func main() {
	a := 15;

	var res = make([]int, 0)
	numericalSystem := 2
	for {
		bit := a % numericalSystem
		res = append(res, bit)
		a /= numericalSystem
		if 0 == a {
			break
		}
	}
	reverse(res)

	fmt.Printf("%v\n", res)
	fmt.Printf("%v\n", toInt(res))


}

func reverse(a []int) {
	for i := len(a)/2-1; i >= 0; i-- {
		opp := len(a)-1-i
		a[i], a[opp] = a[opp], a[i]
	}
}

func toInt(a []int) (result int){

	for _, b := range a {
		result |= b

		result <<= 1
	}

	return result >> 1
}