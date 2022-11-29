package word

import (
	"bharathi/Ninjaexercise13/ex2/quote"
	"testing"
)

func TestCount(t *testing.T) {

	if x := Count("a b c b b"); x != 5 {
		t.Error("got", x, "expected 5")
	}
}

func TestUseCount(t *testing.T) {

	x := UseCount("a b b b")
	for k, v := range x {
		switch k {
		case "a":
			if x["a"] != 1 {
				t.Error("got ", v, "expected", 1)
			}
		case "b":
			if x["b"] != 3 {
				t.Error("got ", v, "expected", 3)
			}
		}
	}

}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
