> go test -bench .

goos: windows

goarch: amd64

cpu: 13th Gen Intel(R) Core(TM) i7-1370P

BenchmarkContains-20                     1420240               777.4 ns/op

BenchmarkStringsContains-20             25506679                55.94 ns/op

BenchmarkWgoContains-20                  1000000              1059 ns/op

BenchmarkWgoStringsContains-20          21460915                48.68 ns/op

BenchmarkWgSwitchContains-20             1205704               918.1 ns/op

BenchmarkSliceContains-20               18048748                62.82 ns/op

PASS
