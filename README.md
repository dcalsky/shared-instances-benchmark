Shared Instances Benchmark
======
Using echo as the web framework

## Benchmarks

```
goos: darwin
goarch: amd64
pkg: github.com/dcalsky/shared-instances-benchmark
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkContext-8     	 5960125	       194.8 ns/op	      36 B/op	       2 allocs/op
BenchmarkSingleton-8   	 8152830	       150.1 ns/op	      26 B/op	       1 allocs/op
PASS
ok  	github.com/dcalsky/shared-instances-benchmark	2.772s
```

## Running Benchmark

```shell
make bench
```