package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://google.com",
		"Yahoo":    "https://yahoo.com",
		"Facebook": "https://facebook.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Yahoo")
	fmt.Println(websites)
}
