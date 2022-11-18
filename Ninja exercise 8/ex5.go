package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// func (x StringSlice) Sort()
func print(users []user) {
	for _, v := range users {
		fmt.Println(v.First)
		fmt.Println(v.Last)
		fmt.Println(v.Age)
		fmt.Println(v.Sayings)

	}

}
func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	print(users)

	// func Strings(x []string) { Sort(StringSlice(x))
	// for i_ ; v := range (users){

	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	fmt.Println("Sort by age")
	print(users)

	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Last < users[j].Last
	})
	fmt.Println("Sort by name")
	print(users)
	for _, v := range users {
		sort.Slice(v.Sayings, func(i, j int) bool {
			return v.Sayings[i] < v.Sayings[j]
		})

	}
	fmt.Println("Sort by sayings for each user")
	print(users)
	// fmt.Println(users)
	// your code goes here

}
