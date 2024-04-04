# Functional go

To run example

```bash
go run src/main.go
```

Benchmarks (during 10 second):

| Name                         | Total ops number | Ns/op       | Bytes/op  | Allocs/op     |
|------------------------------|------------------|-------------|-----------|---------------|
| BenchmarkClassic-8           | 127980086        | 93.61 ns/op | 56 B/op   | 3 allocs/op   |
| BenchmarkStream-8            | 1949524          | 6158 ns/op  | 2608 B/op | 117 allocs/op |
| BenchmarkClassicFilterOnly-8 | 127329626        | 94.33 ns/op | 56 B/op   | 3 allocs/op   |
| BenchmarkStreamFilterOnly-8  | 11897248         | 1007 ns/op  | 376 B/op  | 20 allocs/op  |
| BenchmarkClassicMapOnly-8    | 85813686         | 138.9 ns/op | 248 B/op  | 5 allocs/op   |
| BenchmarkStreamMapOnly-8     | 4890870          | 2442 ns/op  | 1232 B/op | 49 allocs/op  |
| BenchmarkClassicReduceOnly-8 | 374714445        | 32.06 ns/op | 0 B/op    | 0 allocs/op   |
| BenchmarkStreamReduceOnly-8  | 37207575         | 323.0 ns/op | 136 B/op  | 6 allocs/op   |
