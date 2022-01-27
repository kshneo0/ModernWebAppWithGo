package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {

	myMap := make(map[string]User)

	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	myMap["me"] = me

	log.Println((myMap["me"].FirstName))

	var mySlice []string
	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Mary")

	log.Println(mySlice)

	var mySliceInt []int

	mySliceInt = append(mySliceInt, 2)
	mySliceInt = append(mySliceInt, 1)
	mySliceInt = append(mySliceInt, 3)

	log.Println(mySliceInt)

	sort.Ints(mySliceInt)

	log.Println(mySliceInt)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)
	log.Println(numbers[0:2])
	log.Println(numbers[6:9])

	names := []string{"onr", "seven", "fish", "cat"}
	log.Println(names)

}
