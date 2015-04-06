# LazyList
A Lazy Concurrent List-Based Set Algorithm in Go

This implementation is based on the following white paper: http://people.csail.mit.edu/shanir/publications/Lazy_Concurrent.pdf

Goals
-----
Eventually I want to use this library to make a [lockfree skiplist](http://www.cs.tau.ac.il/~shanir/nir-pubs-web/Papers/OPODIS2006-BA.pdf) but for now this is all.  I will add benchmarks of various types over time and would be curious to have people test speed on varyious hardware configurations. 
Mine is tested on the following:
* AMD FX-8120 Zambezi 8-Core 3.1GHz
* 16 GB DDR3 RAM @ 1600Mhz
```
BenchmarkAdd-8       100          19447525 ns/op
ok      github.com/gaigepr/lazylist     2.009s

```
