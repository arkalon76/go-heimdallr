# go-heimdallr
Validates a variety of different personal identification numbers. It will check correct formatting and length, and if possible, checksum. It will not **NOT** check if its a "real" ID number.
![Build Status](https://goreportcard.com/badge/github.com/arkalon76/go-heimdallr)

# Usage
```
heimdallr.ValidateHKID("M812318(2)")
```
# Supported ID's
- Hong Kong National ID (HKID)

# Benchmark
```
goos: linux
goarch: amd64
pkg: github.com/arkalon76/go-heimdallr
BenchmarkValidateHKID-4   	  126553	      9609 ns/op	    6064 B/op	      83 allocs/op
PASS
ok  	github.com/arkalon76/go-heimdallr	1.317s
```