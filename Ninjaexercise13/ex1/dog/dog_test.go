package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(3)
	if y != 21 {
		t.Error("got", y, "expected 21")
	}

}
func TestYearsTwo(t *testing.T) {
	y := YearsTwo(3)
	if y != 21 {
		t.Error("got", y, "expected 21")
	}

}

func ExampleYears() {
	fmt.Println(Years(3))
	// output 21
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(3))
	// output 21
}
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(3)
	}

}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(3)
	}

}
