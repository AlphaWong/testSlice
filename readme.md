# Objective
Test the diffidence of 

1. slice[]
1. slice[:]
1. copy(newSlice,len(slice[]))

in string.Join.

# Actual result

```
func doJoinBySlice(input []string) string {
	return strings.Join(input[:], ",")
}

func doJoin(input []string) string {
	return strings.Join(input, ",")
}

func doJoinByCopy(input []string) string {
	tmpCopy := make([]string, len(input))
	copy(tmpCopy, input)
	return strings.Join(tmpCopy, ",")
}
```


```
BenchmarkDoJoinBySlice-4   	       1	9613103626 ns/op

BenchmarkDoJoin-4          	       1	7911466316 ns/op

BenchmarkDoJoinByCopy-4    	       1	26701439306 ns/op
```

```
BenchmarkDoJoin-4          	       1	9691130814 ns/op

BenchmarkDoJoinBySlice-4   	       1	7854835957 ns/op

BenchmarkDoJoinByCopy-4    	       1	27255503465 ns/op

```

# Result 
- Bench of slice[:] === slice
- Copy slice is the worst case.

# Run
<code>go test -bench=.</code>

# Requirement
- Go 1.7+

# Reference 
1. [https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)

