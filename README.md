[![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/smallnest/weighted?status.png)](http://godoc.org/github.com/smallnest/weighted)  [![travis](https://travis-ci.org/smallnest/weighted.svg?branch=master)](https://travis-ci.org/smallnest/weighted) [![coveralls](https://coveralls.io/repos/smallnest/weighted/badge.svg?branch=master&service=github)](https://coveralls.io/github/smallnest/weighted?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/smallnest/weighted)](https://goreportcard.com/report/github.com/smallnest/weighted)


**rust version**: [weighted-rs](https://github.com/smallnest/weighted-rs)

Package **weighted** implements the smooth weighted round-robin balancing algorithm. This algorithm is implemented in Nginx:
https://github.com/phusion/nginx/commit/27e94984486058d73157038f7950a0a36ecc6e35.

**Notice**: The weighted is NOT goroutine-safe so you MUST use the synchronization primitive to protect it (the Next method) in concurrent cases.


Algorithm is as follows: on each peer selection we increase current_weight
of each eligible peer by its weight, select peer with greatest current_weight
and reduce its current_weight by total number of weight points distributed
among peers.

In case of { 5, 1, 1 } weights this gives the following sequence of
current_weight's: (a, a, b, a, c, a, a)

This is an example to use it:

```go
package main

import "fmt"

func ExampleW1_Next() {
	w := &SW{}
	w.Add("a", 5)
	w.Add("b", 2)
	w.Add("c", 3)

	for i := 0; i < 10; i++ {
		fmt.Printf("%s ", w.Next())
	}
}
```

And this lib has provides another weighted round robin algorithm. This algorithm is used in [LVS](http://kb.linuxvirtualserver.org/wiki/Weighted_Round-Robin_Scheduling).
It has better performance but it is not so more smooth than the first algorithm, so you can select one algorithm according to your case. It is used like the first:

```go
package main

import "fmt"

func ExampleW2_Next() {
	w := &W2{}
	w.Add("a", 5)
	w.Add("b", 2)
	w.Add("c", 3)

	for i := 0; i < 10; i++ {
		fmt.Printf("%s ", w.Next())
	}
}
```