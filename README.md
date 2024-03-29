# CRNDM - Generate Random Numbers with Hardware

[![Go Reference](https://pkg.go.dev/badge/github.com/SeanTolstoyevski/crndm.svg)](https://pkg.go.dev/github.com/SeanTolstoyevski/crndm)

## Benchmark Results:

### Note

`Rdrand` and `Rdseed` instructions  are **not designed for speed**. Moreover, the process is even slower because Golang cannot inline assembly codes.

***

```bash
...>go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/SeanTolstoyevski/crndm
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkCrmdnAsUint64-12       10217383               115.6 ns/op
BenchmarkCrmdnAsUint32-12       10326488               116.5 ns/op
BenchmarkCrmdnAsUint16-12       10796784               114.0 ns/op
PASS
ok      github.com/SeanTolstoyevski/crndm       4.756s
```

## Sources

[Rdrand - Wikipedia](https://en.wikipedia.org/wiki/RDRAND)

[A Quick Guide to Go's Assembler - The Go Programming Language](https://go.dev/doc/asm)

