package main

import (
	"fmt"
	"strings"

	handler "qubecinemas.com/handler"
)

//ASSIGNMENT LINK - https://github.com/RealImage/challenge2015
//Input 1: martin-scorsese and robert-de-niro
//Output 1: Degrees of Seperation is 1

func main() {
	var firstPerson string
	var secondPerson string

	fmt.Print("Enter the First Person - ")
	fmt.Scanf("%s", &firstPerson)
	fmt.Print("Enter the Second Person - ")
	fmt.Scanf("%s", &secondPerson)

	fmt.Println()
	fmt.Println("Sit back for a while until we come back with Minimum Degree of Seperation")
	fmt.Println()

	if strings.ToLower(firstPerson) == strings.ToLower(secondPerson) {
		fmt.Println("No Degree of Seperation!")
	} else {
		_, err := handler.GetProcessedData(strings.ToLower(firstPerson), strings.ToLower(secondPerson), 0)
		if err != nil {
			fmt.Println(err)
		} else {
			min := handler.MinCount[0]
			for i := 1; i < len(handler.MinCount); i++ {
				if min > handler.MinCount[i] {
					min = handler.MinCount[i]
				}
			}
			fmt.Println()
			fmt.Println("Minimum Degree of Seperation is ", min)
		}
	}
}
