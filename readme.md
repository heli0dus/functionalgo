# Functional go

To run example

```bash
go run src/main.go
```

Benchmarks (during 10 second):

| Name                         | Total ops number | Ns/op       | Bytes/op  | Allocs/op     |
|------------------------------|------------------|-------------|-----------|---------------|
| BenchmarkClassic-8           | 128348689        | 93.30 ns/op | 56 B/op   | 3 allocs/op   |
| BenchmarkStream-8            | 1956348          | 6137 ns/op  | 2608 B/op | 117 allocs/op |
| BenchmarkClassicFilterOnly-8 | 127557206        | 94.35 ns/op | 56 B/op   | 3 allocs/op   |
| BenchmarkStreamFilterOnly-8  | 67234030         | 178.0 ns/op | 144 B/op  | 4 allocs/op   |
| BenchmarkClassicMapOnly-8    | 86066364         | 138.9 ns/op | 248 B/op  | 5 allocs/op   |
| BenchmarkStreamMapOnly-8     | 4937497          | 2433 ns/op  | 1232 B/op | 49 allocs/op  |
| BenchmarkClassicReduceOnly-8 | 374674237        | 32.04 ns/op | 0 B/op    | 0 allocs/op   |
| BenchmarkStreamReduceOnly-8  | 77112294         | 155.2 ns/op | 88 B/op   | 4 allocs/op   |