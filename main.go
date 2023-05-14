package main

import (
	"fmt"
	"strings"

	handler "qubecinemas.com/handler"
)

//ASSIGNMENT LINK - https://github.com/RealImage/challenge2015
// Input: amitabh-bachchan and robert-de-niro
// Output: Degrees of Seperation

func main() {
	var firstPerson string
	var secondPerson string

	fmt.Print("Enter the First Person - ")
	fmt.Scanf("%s", &firstPerson)
	fmt.Println()
	fmt.Println("Enter the Second Person - ")
	fmt.Scanf("%s", &secondPerson)

	fmt.Println("Sit back for a while until we come back with Minimum Degree of Seperation")

	if strings.ToLower(firstPerson) == strings.ToLower(secondPerson) {
		fmt.Println("No Degree of Seperation")
	} else {
		result, err := handler.GetProcessedData(firstPerson, secondPerson)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}
