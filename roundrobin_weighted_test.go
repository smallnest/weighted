package weighted

import "testing"

func TestRRW_Next(t *testing.T) {
	w := &RRW{}
	w.Add("server1", 5)
	w.Add("server2", 2)
	w.Add("server3", 3)

	results := make(map[string]int)

	for i := 0; i < 100; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if results["server1"] != 50 || results["server2"] != 20 || results["server3"] != 30 {
		t.Error("the algorithm is wrong", results)
	}

	w.Reset()
	results = make(map[string]int)

	for i := 0; i < 100; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if results["server1"] != 50 || results["server2"] != 20 || results["server3"] != 30 {
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

	if results["server1"] != 7000 || results["server2"] != 9000 || results["server3"] != 13000 {
		t.Error("the algorithm is wrong", results)
	}

	w.RemoveAll()
	w.Add("server1", 0)
	w.Add("server2", 0)

	results = make(map[string]int)

	for i := 0; i < 20000; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if results["server1"] != 10000 || results["server2"] != 10000 {
		t.Error("the algorithm is wrong", results)
	}
}
