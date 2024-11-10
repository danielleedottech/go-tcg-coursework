package main

import "fmt"

func main() {
	hobbies := [3]string{"biking", "reading", "coding"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// step 3
	hobbiesSlice := hobbies[0:2]
	hobbiesSlice2 := []string{}
	hobbiesSlice2 = append(hobbiesSlice2, hobbies[0])
	hobbiesSlice2 = append(hobbiesSlice2, hobbies[1])

	fmt.Println(hobbiesSlice, hobbiesSlice2)

	// step 4
	hobbiesReSlice := hobbies[1:]
	fmt.Println(hobbiesReSlice)

	// step 5
	courseGoals := []string{"learn go", "learn rest", "learn protobuff"}
	courseGoals[1] = "learn rust"
	courseGoals = append(courseGoals, "learn algo")
	fmt.Println(courseGoals)

	type Product struct {
		title string
		price float64
		id    int
	}

	products := []Product{
		Product{
			title: "book",
			price: 12.43,
			id:    1,
		},
		Product{
			title: "game",
			price: 60.00,
			id:    2,
		},
	}

	fmt.Println(products)

	products = append(products, Product{title: "video", price: 3.44, id: 3})
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
