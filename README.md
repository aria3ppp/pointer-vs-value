### Golang pointer vs. value performance benchmark

Should you pass arguments by value or pointer to get the last bit of performance out of cpu?

clone the repo and run benchmarks on your system:
```bash
go test -bench .
```
result of running benchmark on my own local system:

```
goos: linux
goarch: amd64
pkg: github.com/aria3ppp/pointer-vs-value
cpu: Intel(R) Core(TM) i3-1005G1 CPU @ 1.20GHz
BenchmarkSingleFieldIntStruct/value-4   839958211                1.361 ns/op
BenchmarkSingleFieldIntStruct/pointer-4                 1000000000               1.168 ns/op
BenchmarkTwoFieldIntStruct/value-4                      954404379                1.199 ns/op
BenchmarkTwoFieldIntStruct/pointer-4                    963290631                1.225 ns/op
BenchmarkFourFieldIntStruct/value-4                     936266030                1.196 ns/op
BenchmarkFourFieldIntStruct/pointer-4                   947216920                1.203 ns/op
BenchmarkEightFieldIntStruct/value-4                    420129532                2.809 ns/op
BenchmarkEightFieldIntStruct/pointer-4                  812859490                1.390 ns/op
BenchmarkSixteenFieldIntStruct/value-4                  202114904                5.929 ns/op
BenchmarkSixteenFieldIntStruct/pointer-4                901845739                1.284 ns/op
BenchmarkSingleFieldStringStruct/value-4                1000000000               0.9808 ns/op
BenchmarkSingleFieldStringStruct/pointer-4              960792352                1.197 ns/op
BenchmarkTwoFieldStringStruct/value-4                   810558349                1.343 ns/op
BenchmarkTwoFieldStringStruct/pointer-4                 1000000000               1.333 ns/op
BenchmarkFourFieldStringStruct/value-4                  417920527                2.826 ns/op
BenchmarkFourFieldStringStruct/pointer-4                930331584                1.195 ns/op
BenchmarkEightFieldStringStruct/value-4                 205161762                5.789 ns/op
BenchmarkEightFieldStringStruct/pointer-4               953009877                1.191 ns/op
BenchmarkSixteenFieldStringStruct/value-4               129371098                9.273 ns/op
BenchmarkSixteenFieldStringStruct/pointer-4             1000000000               1.065 ns/op
```

### TLDR
Golang is a fast enough programming language and 8 nano seconds is not that big a deal!

So feel free to [write simple and readable code](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=14m35s) and try not to be smart!
