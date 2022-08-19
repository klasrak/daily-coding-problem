# daily-coding-problem
https://www.dailycodingproblem.com/

### Day 1

Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
Bonus: Can you do this in one pass?

```sh
Running tool: /Users/danaugusto/.go/bin/go test -timeout 30s -run ^TestTwoSum$ github.com/klasrak/daily-coding-problem/d1

=== RUN   TestTwoSum
=== RUN   TestTwoSum/twoSum([10_15_3_7],_17)_should_return_true)
=== RUN   TestTwoSum/twoSum([10_15_3_7],_25)_should_return_true)
=== RUN   TestTwoSum/twoSum([10_15_3_7],_26)_should_return_false)
=== RUN   TestTwoSum/twoSum([10_15_3_7],_10)_should_return_true)
=== RUN   TestTwoSum/twoSum([10_15_3_7],_11)_should_return_false)
--- PASS: TestTwoSum (0.00s)
    --- PASS: TestTwoSum/twoSum([10_15_3_7],_17)_should_return_true) (0.00s)
    --- PASS: TestTwoSum/twoSum([10_15_3_7],_25)_should_return_true) (0.00s)
    --- PASS: TestTwoSum/twoSum([10_15_3_7],_26)_should_return_false) (0.00s)
    --- PASS: TestTwoSum/twoSum([10_15_3_7],_10)_should_return_true) (0.00s)
    --- PASS: TestTwoSum/twoSum([10_15_3_7],_11)_should_return_false) (0.00s)
PASS
ok      github.com/klasrak/daily-coding-problem/d1      (cached)


> Test run finished at 19/08/2022 13:19:01 <

Running tool: /Users/danaugusto/.go/bin/go test -benchmem -run=^$ -bench ^BenchmarkTwoSum$ github.com/klasrak/daily-coding-problem/d1

goos: darwin
goarch: amd64
pkg: github.com/klasrak/daily-coding-problem/d1
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkTwoSum
BenchmarkTwoSum-12      20174986                51.85 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/klasrak/daily-coding-problem/d1      1.686s


> Test run finished at 19/08/2022 13:19:11 <
```
