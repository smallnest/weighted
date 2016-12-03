/*
Package weighted implements the smooth weighted round-robin balancing algorithm. This algorithm is implemented in Nginx:
https://github.com/phusion/nginx/commit/27e94984486058d73157038f7950a0a36ecc6e35.

Algorithm is as follows: on each peer selection we increase current_weight
of each eligible peer by its weight, select peer with greatest current_weight
and reduce its current_weight by total number of weight points distributed
among peers.

In case of { 5, 1, 1 } weights this gives the following sequence of
current_weight's: (a, a, b, a, c, a, a)

*/
package weighted

// Weighted is a wrapped server with weight
type Weighted struct {
	Server          interface{}
	Weight          int
	CurrentWeight   int
	EffectiveWeight int
}

// W is struct that contains weighted servers and provides methods to select a weighted server.
type W struct {
	servers []*Weighted
	n       int
}

func (w *Weighted) fail() {
	w.EffectiveWeight -= w.Weight
	if w.EffectiveWeight < 0 {
		w.EffectiveWeight = 0
	}
}

// Add a weighted server.
func (w *W) Add(server interface{}, weight int) {
	weighted := &Weighted{Server: server, Weight: weight, EffectiveWeight: weight}
	w.servers = append(w.servers, weighted)
	w.n++
}

// RemoveAll removes all weighted servers.
func (w *W) RemoveAll() {
	w.servers = w.servers[:0]
	w.n = 0
}

//Reset resets all current weights.
func (w *W) Reset() {
	for _, s := range w.servers {
		s.EffectiveWeight = s.Weight
		s.CurrentWeight = 0
	}
}

// Next returns next selected server.
func (w *W) Next() interface{} {
	if w.n == 0 {
		return nil
	}
	if w.n == 1 {
		return w.servers[0].Server
	}

	return nextWeighted(w.servers).Server
}

// NextWeighted returns next selected weighted object.
func (w *W) NextWeighted() *Weighted {
	if w.n == 0 {
		return nil
	}
	if w.n == 1 {
		return w.servers[0]
	}

	return nextWeighted(w.servers)
}

//https://github.com/phusion/nginx/commit/27e94984486058d73157038f7950a0a36ecc6e35
func nextWeighted(servers []*Weighted) (best *Weighted) {
	total := 0

	for i := 0; i < len(servers); i++ {
		w := servers[i]

		if w == nil {
			continue
		}

		w.CurrentWeight += w.EffectiveWeight
		total += w.EffectiveWeight
		if w.EffectiveWeight < w.Weight {
			w.EffectiveWeight++
		}

		if best == nil || w.CurrentWeight > best.CurrentWeight {
			best = w
		}

	}

	if best == nil {
		return nil
	}

	best.CurrentWeight -= total
	return best
}
