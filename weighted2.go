package weighted

// weighted2 is a wrapped server with weight that is used to implement LVS weighted round robin algorithm.
type weighted2 struct {
	Server interface{}
	Weight int
}

// W2 is struct that contains weighted servers implement LVS weighted round robin algorithm.
//
// http://kb.linuxvirtualserver.org/wiki/Weighted_Round-Robin_Scheduling
// http://zh.linuxvirtualserver.org/node/37
type W2 struct {
	servers []*weighted2
	n       int
	gcd     int
	maxW    int
	i       int
	cw      int
}

// Add a weighted server.
func (w *W2) Add(server interface{}, weight int) {
	weighted := &weighted2{Server: server, Weight: weight}
	if weight > 0 {
		if w.gcd == 0 {
			w.gcd = weight
			w.maxW = weight
			w.i = -1
			w.cw = 0
		} else {
			w.gcd = gcd(w.gcd, weight)
			if w.maxW < weight {
				w.maxW = weight
			}
		}
	}
	w.servers = append(w.servers, weighted)
	w.n++
}

// RemoveAll removes all weighted servers.
func (w *W2) RemoveAll() {
	w.servers = w.servers[:0]
	w.n = 0
	w.gcd = 0
	w.maxW = 0
	w.i = -1
	w.cw = 0
}

//Reset resets all current weights.
func (w *W2) Reset() {
	w.i = -1
	w.cw = 0
}

// Next returns next selected server.
func (w *W2) Next() interface{} {
	if w.n == 0 {
		return nil
	}

	if w.n == 1 {
		return w.servers[0].Server
	}

	for {
		w.i = (w.i + 1) % w.n
		if w.i == 0 {
			w.cw = w.cw - w.gcd
			if w.cw <= 0 {
				w.cw = w.maxW
				if w.cw == 0 {
					return nil
				}
			}
		}

		if w.servers[w.i].Weight >= w.cw {
			return w.servers[w.i].Server
		}
	}
}

func gcd(x, y int) int {
	var t int
	for {
		t = (x % y)
		if t > 0 {
			x = y
			y = t
		} else {
			return y
		}
	}
}
