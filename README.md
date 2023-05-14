# fib
This is a simple package to expose high speed functions related to Fibonacci numbers using methods that do not use the recurrence relation F<Sub>n+1</sub> = F<Sub>n</sub> + F<Sub>n-1</sub>.

## `FibonacciNumberWithIndex(index int) int64`
This function finds the Fibinacci number having the given index. The indices may be negative. Note this function is only accurate for indices between -75 and 75 inclusive.

## `FibonacciFloorIndexFor(number int64) int`
This function finds the index of the greatest Fibonacci number that does not exceed the given number. Note that positive numbers may be ambiguous, as some indices give the same number as their opposites.
In this case the positive index is returned. For instance indices 3 and -3 both give Fibonacci number 2. Also number 1 is ambiguous, as indices 1 and 2 both give Fibonacci number 1. This function may be inaccurate for numbers less than -2111485077978050
or greater than 2111485077978050.
