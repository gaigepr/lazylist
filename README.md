# LazyList
A Lazy Concurrent List-Based Set Algorithm in Go with lock and wait free membership testing. 

This implementation is based on the following white paper: http://people.csail.mit.edu/shanir/publications/Lazy_Concurrent.pdf

If you have serious doubts about the safety or purpose of this data structure please refer to te above paper. I would prefer feedback to be related to *my* implemenatation and its merits or lack thereof. If you see something that, after reading the paper, seems wrong let me know!

Examples
-------
```
package main

import (
        "github.com/gaigepr/lazylist"

        "fmt"
)

func main() {
        fmt.Println("We add 10 things to the list 1 at a time.")
        list := lazylist.NewLazyList()
        for i := 0; i < 10; i++ {
                list.Add(uint64(i))
        }
        list.PrintLazyList()
        
        // We add 10 things to the list 1 at a time.
        // 0
        // 1
        // 2
        // 3
        // 4
        // 5
        // 6
        // 7
        // 8
        // 9

        fmt.Println("We remove even numbers from the list.")
        for i := 0; i < 10; i++ {
                if i%2 == 0 {
                        list.Remove(uint64(i))
                }
        }
        list.PrintLazyList()
        
        // We remove even numbers from the list.
        // 1
        // 3
        // 5
        // 7
        // 9

        fmt.Println("Is 9 in there?")
        fmt.Println(list.Contains(uint64(9)))
        
        // Is 9 in there?
        // true

        // Concurrent examples coming soon!
        // You can just put "go" in front of any of the functions and it'll work!
}
```

Goals
-----
Eventually I want to use this library to make a [lockfree skiplist](http://www.cs.tau.ac.il/~shanir/nir-pubs-web/Papers/OPODIS2006-BA.pdf) but for now this is all.  I will add benchmarks of various types over time and would be curious to have people test speed on varyious hardware configurations.

Benchmarking
------------
Mine is tested on the following:
* AMD FX-8120 Zambezi 8-Core 3.1GHz
* 16 GB DDR3 RAM @ 1600Mhz

This test only test 1000 'simeltaneous' inserts. Nothing more so far.
```
$ GOMAXPROCS=8 go test -bench=.

BenchmarkAdd-8       100          19447525 ns/op
ok      github.com/gaigepr/lazylist     2.009s
```
With up to 4 Processes
```
BenchmarkAdd-4       100          34776539 ns/op
ok      github.com/gaigepr/lazylist     3.544s
```
With up to 1 Processes
```
$ GOMAXPROCS=1 go test -bench=.

BenchmarkAdd         100         185902241 ns/op
ok      github.com/gaigepr/lazylist     18.689s
```
