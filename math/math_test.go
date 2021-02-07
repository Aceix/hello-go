package math

import (
	"fmt"
	"testing"
)

/// Commented out ti test "go test -coverprofile=coverage.out"
/// View report with "go tool cover -html=coverage.out"
// func TestAdd(t *testing.T) {

// 	got := Add(2, 3)
// 	want := 5

// 	if got != want {
// 		t.Errorf("Wanted: %v, got: %v", want, got)
// 	}
// }

func TestSubtract(t *testing.T) {

	type arg struct {
		a int
		b int
	}

	args := []struct {
		name string
		args arg
		want int
	}{
		{"", arg{1, 2}, -1},
		{"", arg{10, 2}, 8},
		{"", arg{20, 2}, 18},
		{"", arg{-1, -9}, 8},
	}

	for _, x := range args {
		if got := Subtract(x.args.a, x.args.b); got != x.want {
			t.Errorf("got Subtract() = %v; want = %v", got, x.want)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

func ExampleAdd() {
	fmt.Println(Add(1, 4))
	// Output: 5
}
