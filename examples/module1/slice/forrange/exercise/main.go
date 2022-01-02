package main

import "fmt"

func main() {
	mySlice := []string{"I", "am", "stupid", "and", "weak"}
	for index, value := range mySlice {
		if value == "stupid" {
			mySlice[index] = "smart"
		} else if value == "weak" {
			mySlice[index] = "strong"
		}
	}
	fmt.Printf("mySlice %+v\n", mySlice)
}
