package fib

import (
    "math"
)

var (
    root_five float64 = math.Sqrt(5)
)

func FibonacciNumberWithIndex(index int) int64 {
    if index == 0 {
        return 0
    }
    modulus := int64(math.Round(math.Pow((root_five + float64(1)) / float64(2), math.Abs(float64(index))) / root_five))
    if index > 0 {
        return modulus
    }
    if index % 2 == 0 {
        return -modulus
    }
    return modulus
}

func FibonacciFloorIndexFor(number int64) int {
    if number == 0 {
        return 0
    }
    indexFloat := math.Log(root_five * math.Abs(float64(number))) / math.Log((float64(1) + root_five) / float64(2))
    indexApprox := math.Round(indexFloat)
    
    if number > 0 {
        if indexFloat >= indexApprox {
            return int(indexApprox)
        }
        if FibonacciNumberWithIndex(int(indexApprox)) > number {
            return int(indexApprox) - 1
        }
        return int(indexApprox)
    }
    if indexFloat <= indexApprox {
        return - int(indexApprox)
    }
    if FibonacciNumberWithIndex(int(indexApprox)) < -number {
        return - int(indexApprox) - 1
    }
    return -int(indexApprox)
}

