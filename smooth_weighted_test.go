package weighted

import "testing"

func TestSW_Next(t *testing.T) {
	w := &SW{}
	w.Add("server1", 5)
	w.Add("server2", 2)
	w.Add("server3", 3)

	results := make(map[string]int)

	for i := 0; i < 100; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if results["server1"] != 50 || results["server2"] != 20 || results["server3"] != 30 {
		t.Error("the algorithm is wrong")
	}

	w.Reset()
	results = make(map[string]int)

	for i := 0; i < 100; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if results["server1"] != 50 || results["server2"] != 20 || results["server3"] != 30 {
		t.Error("the algorithm is wrong")
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

	if results["server1"] != 7000 || results["server2"] != 9000 || results["server3"] != 13000 {
		t.Error("the algorithm is wrong")
	}

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
