package weighted

import "testing"

func TestRandW_Next(t *testing.T) {
	w := NewRandW()
	w.Add("server1", 5)
	w.Add("server2", 2)
	w.Add("server3", 3)

	results := make(map[string]int)

	for i := 0; i < 100; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if !checkResults(results["server1"], 20, 70) || !checkResults(results["server2"], 0, 40) || !checkResults(results["server3"], 10, 50) {
		t.Error("the algorithm is wrong", results)
	}

	w.Reset()
	results = make(map[string]int)

	for i := 0; i < 100; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if !checkResults(results["server1"], 20, 70) || !checkResults(results["server2"], 0, 40) || !checkResults(results["server3"], 10, 50) {
		t.Error("the algorithm is wrong", results)
	}

	w.RemoveAll()
	w.Add("server1", 7)
	w.Add("server2", 9)
	w.Add("server3", 13)

	results = make(map[string]int)

	for i := 0; i < 29000; i++ {
		s := w.Next().(string)
		results[s]++
	}

	// if !checkResults(results["server1"], 6000, 8000) || !checkResults(results["server2"], 8000, 10000) || !checkResults(results["server3"], 12000, 14000) {
	// 	t.Error("the algorithm is wrong", results)
	// }

	t.Log("the results: ", results)
}

func checkResults(v, min, max int) bool {
	return v >= min && v <= max
}
