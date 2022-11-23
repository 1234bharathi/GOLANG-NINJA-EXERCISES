package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		//  the purpose is to marshal the code thus if it does not exit
		log.Fatalln(err)
		// or else println and return
		// log.Println(err)
		// return

	}
	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {

		// with errorf

		return []byte{}, fmt.Errorf("was not able to marshal %v", err)
		// with return new
		// return []byte{}, errors.New(fmt.Sprintf("was not able to marshal%v", err))
		// return []byte{}, errors.New("was not able to marshal")
	}

	return bs, nil

}
