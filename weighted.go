/*
Package weighted implements two weighted round robin algorithms.
One is the smooth weighted round-robin balancing algorithm used in Nginx, and you can use "w := new SW{}" to use it.
The other is wrr used in LVS and you can use "w := new RRW{}" to use it.

For Nginx smooth weighted round-robin balancing algorithm, you can check https://github.com/phusion/nginx/commit/27e94984486058d73157038f7950a0a36ecc6e35.
For LVS round-robin balancing algorithm, you can check http://kb.linuxvirtualitem.org/wiki/Weighted_Round-Robin_Scheduling.
*/
package weighted

// W is a interface that implement a weighted round robin algorithm.
type W interface {
	// Next gets next selected item.
	Next() (item interface{})
	// Add adds a weighted item for selection.
	Add(item interface{}, weight int)

	// All returns all items.
	All() map[interface{}]int

	// RemoveAll removes all weighted items.
	RemoveAll()
	// Reset resets the balancing algorithm.
	Reset()
}
