package main

import (
	"fmt"
	"strings"

	handler "qubecinemas.com/handler"
)

//ASSIGNMENT LINK - https://github.com/RealImage/challenge2015
// Input: amitabh-bachchan and robert-de-niro
//Input 2: amitabh-bachchan and carey-mulligan as second Input
// Output: Degrees of Seperation is 3

func main() {
	var firstPerson string = "amitabh-bachchan"
	var secondPerson string = "robert-de-niro"

	// fmt.Print("Enter the First Person - ")
	// fmt.Scanf("%s", &firstPerson)
	// fmt.Print("Enter the Second Person - ")
	// fmt.Scanf("%s", &secondPerson)

	fmt.Println()
	fmt.Println("Sit back for a while until we come back with Minimum Degree of Seperation")
	fmt.Println()
	nameset := make(map[string]struct{})
	movieset := make(map[string]struct{})
	if strings.EqualFold(firstPerson, secondPerson) {
		fmt.Println("No Degree of Seperation!")
	} else {
		result, err := handler.GetProcessedData(strings.ToLower(firstPerson), strings.ToLower(secondPerson), nameset, movieset, 0)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Degree of Seperation is", result)
		}
	}
}
