package main
// SumInts adds together the values of m.

import "fmt"

func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}

type Number interface {
    int64 | float64
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}


type Vector[T any] struct{
	start *T
	size int32
	capacity int32
}
func Add[T int64 | float64 | int] (a, b T) T{
	return a + b
}

func main() {
	fmt.Println(Add(1,2))
	fmt.Println(Add(1.1,2.2))
    // Initialize a map for the integer values
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumIntsOrFloats(ints),
        SumIntsOrFloats(floats))
}



