package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"mango", "banana", "peach"}

	fmt.Println("Slice ", fruits)

	var animals = make([]string, 5)

	animals[0] = "Tiger"
	animals[1] = "Lion"
	animals[2] = "Fox"
	animals[3] = "Bear"
	animals[4] = "Wolf"
	animals = append(animals, "Elephant", "Horse")

	fmt.Println("The animals : ", animals)

	var highScores = []int{100, 400, 200, 600, 500}
	highScores = append(highScores, 900, 700)

	sort.Strings(animals)
	sort.Ints(highScores)
	fmt.Println("IS sorted highscores ", sort.IntsAreSorted(highScores))
	fmt.Println("IS sorted animals ", sort.StringsAreSorted(animals))

	fmt.Println("Highscores ", highScores)
	fmt.Println("Animals ", animals)

	fmt.Println("-----------------")
	fmt.Println("Courses")
	var courses = make([]string, 5)
	courses[0] = "ReactJS"
	courses[1] = "Javascript"
	courses[2] = "Swift"
	courses[3] = "Python"
	courses[4] = "Java"
	fmt.Println("Courses available ", courses)
	// Remove "Swift"
	index := 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println("Opted courses : ", courses)
}
