package handler

import (
	"fmt"
	"strings"

	srvError "qubecinemas.com/error"
	utility "qubecinemas.com/internal/utility"
)

//Processing of data to find Minimum degree of Seperation between two people
/*
FLOW: Possibility
1. Actor -> Movie -> Second Person (if found) - Minimum Steps
2. Actor -> Movie -> Actor
3. Actor -> Movie -> Actor -> Movie -> Actor
4. *****

NOTE: Actor to Movie is ONE Single Step
*/

var flag bool = false
var MinCount []int

func GetProcessedData(firstPerson, secondPerson string, count int) (int, error) {
	secondPerson = strings.Join(strings.Split(secondPerson, "-"), " ")
	if strings.EqualFold(strings.ToLower(firstPerson), strings.ToLower(secondPerson)) {
		fmt.Println("----------------Minimum Degree of Seperation from the recursion is--------------------", count)
		flag = true
		return count, nil
	}

	//First Person to get movie names
	err, person := utility.PersonUtility(firstPerson)
	if err != nil {
		return 0, fmt.Errorf(srvError.PackError(&srvError.UnableToBindResourceBody, err))
	}
	for _, value := range person.Movies {
		err, movieSub := utility.MovieUtility(value.Url)
		if err != nil {
			return 0, fmt.Errorf(srvError.PackError(&srvError.UnableToBindResourceBody, err))
		}
		people := append(movieSub.Cast, movieSub.Crew...)
		count += 1
		for _, inval := range people {
			if flag == true {
				flag = false
				MinCount = append(MinCount, count)
				count = 0
			}
			GetProcessedData(strings.ToLower(inval.Name), secondPerson, count)
		}

	}
	return count, nil
}
