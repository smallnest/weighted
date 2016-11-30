package weighted

import "fmt"

func ExampleW_() {
	w := &W{}
	w.Add("a", 5)
	w.Add("b", 2)
	w.Add("c", 3)

	for i := 0; i < 10; i++ {
		fmt.Printf("%s ", w.Next())
	}

	// Output: a c b a a c a b c a
}
