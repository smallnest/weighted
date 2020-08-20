package weighted

import (
	"time"

	"golang.org/x/exp/rand"
)

// randWeighted is a wrapped weighted item that is used to implement weighted random algorithm.
type randWeighted struct {
	Item   interface{}
	Weight int
}

// RandW is a struct that contains weighted items implement weighted random algorithm.
type RandW struct {
	items        []*randWeighted
	n            int
	sumOfWeights int
	r            *rand.Rand
}

// NewRandW creates a new RandW with a random object.
func NewRandW() *RandW {
	return &RandW{r: rand.New(rand.NewSource(uint64(time.Now().UnixNano())))}
}

// Next returns next selected item.
func (rw *RandW) Next() (item interface{}) {
	if rw.n == 0 {
		return nil
	}
	randomWeight := rw.r.Intn(rw.sumOfWeights) + 1
	for _, item := range rw.items {
		randomWeight = randomWeight - item.Weight
		if randomWeight <= 0 {
			return item.Item
		}
	}

	return rw.items[len(rw.items)-1].Item
}

// Add adds a weighted item for selection.
func (rw *RandW) Add(item interface{}, weight int) {
	rItem := &randWeighted{Item: item, Weight: weight}
	rw.items = append(rw.items, rItem)
	rw.sumOfWeights += weight
	rw.n++
}

// All returns all items.
func (rw *RandW) All() map[interface{}]int {
	m := make(map[interface{}]int)
	for _, i := range rw.items {
		m[i.Item] = i.Weight
	}
	return m
}

// RemoveAll removes all weighted items.
func (rw *RandW) RemoveAll() {
	rw.items = make([]*randWeighted, 0)
	rw.r = rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
}

// Reset resets the balancing algorithm.
func (rw *RandW) Reset() {
	rw.r = rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
}
