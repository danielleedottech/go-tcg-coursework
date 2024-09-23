package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	fmt.Println("Age:", age)
	fmt.Println("Age Pointer:", agePointer)

	getAdultYears(&age)
	fmt.Println("Adult years:", age)

	getAdultYears(agePointer)
	fmt.Println("Adult years:", age)

}

func getAdultYears(age *int) {
	*age = *age - 18
}
