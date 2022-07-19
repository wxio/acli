
```
go test -benchmem -run=^$ -bench=. github.com/wxio/acli/internal/fib
goos: darwin
goarch: amd64
pkg: github.com/wxio/acli/internal/fib
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkIterativeFib3-8        1000000000               0.5588 ns/op          0 B/op          0 allocs/op
BenchmarkIterativeFib10-8       239097409                4.970 ns/op           0 B/op          0 allocs/op
BenchmarkIterativeFib20-8       100000000               10.11 ns/op            0 B/op          0 allocs/op
BenchmarkIterativeFib40-8       52211992                21.71 ns/op            0 B/op          0 allocs/op
BenchmarkIterativeFib45-8       48067636                23.31 ns/op            0 B/op          0 allocs/op
BenchmarkIterativeFib50-8       46168651                24.72 ns/op            0 B/op          0 allocs/op
BenchmarkChannelFib4-8            611361              2062 ns/op             416 B/op          7 allocs/op
BenchmarkChannelFib10-8           191400              6858 ns/op            1329 B/op         25 allocs/op
BenchmarkChannelFib20-8            93724             15596 ns/op            2852 B/op         55 allocs/op
BenchmarkChannelFib40-8            38293             32167 ns/op            5899 B/op        115 allocs/op
BenchmarkChannelFib45-8            33650             37182 ns/op            6660 B/op        130 allocs/op
BenchmarkChannelFib50-8            30837             40873 ns/op            7424 B/op        145 allocs/op
BenchmarkChannelFib60-8            24663             48500 ns/op            8947 B/op        175 allocs/op
BenchmarkChannelFib80-8            18062             66731 ns/op           11993 B/op        235 allocs/op
BenchmarkRecursiveFib1-8        809165820                1.472 ns/op           0 B/op          0 allocs/op
BenchmarkRecursiveFib2-8        810730071                1.514 ns/op           0 B/op          0 allocs/op
BenchmarkRecursiveFib3-8        288685399                4.102 ns/op           0 B/op          0 allocs/op
BenchmarkRecursiveFib10-8        6957492               168.8 ns/op             0 B/op          0 allocs/op
BenchmarkRecursiveFib20-8          56437             20955 ns/op               0 B/op          0 allocs/op
BenchmarkRecursiveFib40-8              4         322636128 ns/op               0 B/op          0 allocs/op
BenchmarkRecursiveFib45-8              1        3503980108 ns/op               0 B/op          0 allocs/op
BenchmarkRecursiveFib50-8              1        40109093884 ns/op              0 B/op          0 allocs/op
PASS
ok      github.com/wxio/acli/internal/fib       73.087s
```