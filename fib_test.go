package fib

import (
    "testing"
)

func TestFibonacciNumberWithIndex(t *testing.T) {
    index := 0
    var number1, number2 int64 = 1, 0
    for index < 76 {
        actualResult := FibonacciNumberWithIndex(index)
        if actualResult != number2 {
            t.Fatalf("FibonacciNumberWithIndex(%d) = %d, but should be %d.", index, actualResult, number2)
        }
        number1, number2 = number2, number1 + number2
        index += 1
    }
    index = -1
    number1, number2 = 1, 0
    for index > -76 {
        actualResult := FibonacciNumberWithIndex(index)
        if actualResult != number1 {
            t.Fatalf("FibonacciNumberWithIndex(%d) = %d, but should be %d.", index, actualResult, number1)
        }
        number1, number2 = number2 - number1, number1
        index -= 1
    }
}

func TestFibonacciFloorIndexFor(t *testing.T) {
    index := 3
    var number1, number2 int64 = 1, 2
    for index < 76 {
        actualResult := FibonacciFloorIndexFor(number2)
        if actualResult != index {
            t.Fatalf("FibonacciFloorIndexFor(%d) = %d, but should be %d.", number2, actualResult, index)
        }
        number1, number2 = number2, number1 + number2
        index += 1
    }
    index = -3
    number1, number2 = 2, -1
    for index > -76 {
        if index %2 == 0 {
            actualResult := FibonacciFloorIndexFor(number1)
            if actualResult != index {
                t.Fatalf("FibonacciFloorIndexFor(%d) = %d, but should be %d.", number1, actualResult, index)
            }
        }
        number1, number2 = number2 - number1, number1
        index -= 1
    }
}

