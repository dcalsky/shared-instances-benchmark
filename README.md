Shared Instances Benchmark
======
Using echo as the web framework

## Benchmarks

```
goos: darwin
goarch: amd64
pkg: github.com/dcalsky/shared-instances-benchmark
BenchmarkContext-8     	 5811720	       211 ns/op	      52 B/op	       2 allocs/op
BenchmarkSingleton-8   	 7633112	       162 ns/op	      27 B/op	       1 allocs/op
PASS
ok  	github.com/dcalsky/shared-instances-benchmark	2.862s
```

## Running Benchmark

```shell
make bench
```