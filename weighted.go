/*
Package weighted implements two weighted round robin algorithms.
One is the smooth weighted round-robin balancing algorithm used in Nginx, and you can use "w := new W1{}" to use it.
The other is wrr used in LVS and you can use "w := new W2{}" to use it. */
package weighted

// W is a interface that implement a weighted round robin algorithm.
type W interface {
	Next() interface{}
	Add(server interface{}, weight int)
	RemoveAll()
	Reset()
}
