package main

import (
	"errors"
	"fmt"
)

type customerr struct {
	errdes string
	err    error
}

func (cerr customerr) Error() string {
	return fmt.Sprintf("DES: %v and error in case: %v", cerr.errdes, cerr.err)

}
func foo(e error) {
	fmt.Printf("calling foo %v", e)
	// assertion
	fmt.Printf("using assertion in foo  %v", e.(customerr).errdes)

}
func main() {
	fmt.Println("Enter a value")
	var no int
	_, err := fmt.Scan(&no)
	if err != nil {
		cr := customerr{
			errdes: fmt.Sprint("It is not number"),
			err:    err,
		}
		foo(cr)
	} else {
		cr := customerr{
			errdes: fmt.Sprintf("the value is %v and the no divides by 2 gives %v", no, no/2),
			err:    errors.New("no error"),
		}
		foo(cr)
	}
}
