package handler

import (
	"fmt"
	"strings"

	srvError "qubecinemas.com/error"
	model "qubecinemas.com/internal/models"
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

// var count int

func GetProcessedData(firstPerson, secondPerson string, nameset map[string]struct{}, movieset map[string]struct{}, count int) (int, error) {
	fmt.Println("nameset", nameset)
	// fmt.Println("moviset", movieset)
	if strings.EqualFold(strings.ToLower(firstPerson), strings.ToLower(secondPerson)) {
		fmt.Println("---------------------------------------------------------------")
		return count, nil
	}
	//First Person to get movie names
	err, person := utility.PersonUtility(firstPerson)
	if err != nil {
		return 0, fmt.Errorf(srvError.PackError(&srvError.UnableToBindResourceBody, err))
	}
	var people = []model.People{}
	var movieSub *model.Movie
	for _, value := range person.Movies {
		fmt.Println("Movie", value.Url)
		_, ok := movieset[value.Url]
		if ok {
			fmt.Println("movie exists")
			continue
		} else {
			movieset[value.Url] = struct{}{}
			err, movieSub = utility.MovieUtility(value.Url)
			if err != nil {
				return 0, fmt.Errorf(srvError.PackError(&srvError.UnableToBindResourceBody, err))
			}
			people = append(people, movieSub.Cast...)
			people = append(people, movieSub.Crew...)
		}
		// fmt.Println(people)
	}
	fmt.Println("-----------", people)
	count += 1
	for _, inval := range people {
		fmt.Println(inval.Url, secondPerson)
		_, ok := nameset[inval.Url]
		if ok {
			fmt.Println("name exists")
			continue
		} else {
			nameset[inval.Url] = struct{}{}
			GetProcessedData(inval.Url, secondPerson, nameset, movieset, count)
		}
	}

	return count, nil
}
