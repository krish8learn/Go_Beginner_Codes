package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	x := Greet()
	if x != "welcome machine" {
		t.Error("Expected: welcome machine ,got:", x)
	}
}

func ExampleGreet() {
	fmt.Println(Greet())
	//output:
	//welcome machine
}

func BenchmarkGreet(b *testing.B) { // go test -bench .
	for i := 0; i < b.N; i++ {
		Greet()
	}
}
