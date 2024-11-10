package main

import "fmt"

//
//func main() {
//	var productNames = [4]string{"A book"}
//	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
//
//	productNames[2] = "A carpet"
//	fmt.Println(prices)
//	fmt.Println(productNames)
//	fmt.Println(prices[2])
//
//	featuredPrices := prices[1:]
//	featuredPrices[0] = 199.99
//	highlightedPrices := featuredPrices[:1]
//	fmt.Println(featuredPrices)
//	fmt.Println(highlightedPrices)
//	fmt.Println(prices)
//
//	fmt.Println(len(prices), cap(prices))
//	fmt.Println(len(featuredPrices), cap(featuredPrices))
//	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
//
//	highlightedPrices = highlightedPrices[:3]
//	fmt.Println(highlightedPrices)
//}

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices)
	fmt.Println(prices[0:1])

	prices[1] = 9.99
	//prices[2] = 2.99 raises error because of cap limit

	updatedPrices := append(prices, 5.99)
	fmt.Println(prices)
	fmt.Println(updatedPrices)

	updatedPrices = updatedPrices[1:]
	fmt.Println(updatedPrices)
}
