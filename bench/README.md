## Performance comparison of home grown algorithm vs. the built in fnv1a function of go (fnv.New64a)

For small strings like in our use case, the fnv1a function of @cybercdh is 4 times faster as the one from go lib:

Example hashing of string "xwforensics64":
```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: main.go/bench
BenchmarkFnv1aGolib-8   	16810881	        64.2 ns/op
BenchmarkFnv1aOwn-8     	65636450	        15.5 ns/op
PASS
ok  	main.go/bench	2.194s
```

Example hashing of 10kb characters, about the same speed:

```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: main.go/bench
BenchmarkFnv1aGolib-8   	   74402	     14522 ns/op
BenchmarkFnv1aOwn-8     	   78336	     13752 ns/op
PASS
ok  	main.go/bench	2.480s
```
Test code in this directory.