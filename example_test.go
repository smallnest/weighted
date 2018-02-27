package weighted

import "fmt"

func ExampleSW_Next() {
	w := &SW{}
	w.Add("a", 5)
	w.Add("b", 2)
	w.Add("c", 3)

	for i := 0; i < 10; i++ {
		fmt.Printf("%s ", w.Next())
	}

	// Output: a c b a a c a b c a
}

func ExampleRRW_Next() {
	w := &RRW{}
	w.Add("a", 5)
	w.Add("b", 2)
	w.Add("c", 3)

	for i := 0; i < 10; i++ {
		fmt.Printf("%s ", w.Next())
	}

	// Output: a a a c a b c a b c
}
