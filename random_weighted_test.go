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

	all := w.All()
	countOK := 0
	for index := range all {
		if (index == "server1" && all[index] == 5) ||
			(index == "server2" && all[index] == 2) ||
			(index == "server3" && all[index] == 3) {
			countOK++
		}
	}
	if countOK != 3 {
		t.Error("the algorithm is wrong")
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

	w.RemoveAll()
	next := w.Next()
	if next != nil {
		t.Error("the algorithm is wrong")
	}
	w.Add("server1", 3)
	next = w.Next()
	if next == nil {
		t.Error("the algorithm is wrong")
	}
}

func checkResults(v, min, max int) bool {
	return v >= min && v <= max
}
