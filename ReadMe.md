# GoFlag
このモジュールは大規模なフラグの管理をできるようにするための機能です。

フラグの管理個数は
- 64
- 4096
- 262144
となっている。

特徴としてはフラグを立てるもしくは下すまでの時間が短いことと

まだフラグの立っていないものを素早く見つけられるという特徴を持っています

非同期アクセスへの対応をしているものとそうでないものを用意する予定です

`
    # cmd sample 
    go test -benchmem -bench BenchmarkGoFlagFindAndUp262144 github.com/handball811/goflag -o prof/prof.bin -cpuprofile='prof/cpu.prof' -memprofile='prof/mem.prof'
    go tool pprof --svg prof/prof.bin prof/cpu.prof > prof/cpuprof.svg

    # generate
    go run cmd/find/main.go -out add.s -stubs stub.go
`

### Benchmark
BenchmarkGoFlagFindAndUp4096_0-8      	61394263	        18.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceFindAndUp4096_0-8       	1000000000	         0.897 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoFlagFindAndUp4096_100-8    	64623487	        18.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceFindAndUp4096_100-8     	40983606	        29.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoFlagFindAndUp4096_1000-8   	64712353	        18.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceFindAndUp4096_1000-8    	 5182806	       231 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoFlagFindAndUp4096_4096-8   	588282298	         2.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceFindAndUp4096_4096-8    	 1284268	       935 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoFlagFindAndUp262144-8      	55902877	        21.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoFlagFindAndUp64-8          	87825870	        13.3 ns/op	       0 B/op	       0 allocs/op

#### Benchmarkの考察
- 安定性を除いてあまり良い結果は得られていない

