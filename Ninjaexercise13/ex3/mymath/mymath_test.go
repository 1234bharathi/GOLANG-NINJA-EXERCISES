package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	x := []int{1, 2, 3, 4, 9}
	val := CenteredAvg(x)
	if val != 3 {
		t.Error("got", val, "expected", 3)
	}

}

func ExampleCenteredAvg() {
	x := []int{1, 2, 3, 4, 9}

	fmt.Println(CenteredAvg(x))
	// output
	// 3
}
func BenchmarkCenteredAvg(b *testing.B) {
	x := []int{1, 2, 3, 4, 9}
	for i := 0; i < b.N; i++ {

		CenteredAvg(x)
	}
}
