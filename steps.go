package main

import (
    "flag"
    "fmt"
    "math/big"
)

var (
    addition       = flag.Bool("a", false, "Perform addition")
    multiplication = flag.Bool("m", false, "Perform multiplication")
    power          = flag.Bool("p", false, "Perform exponentiation")
    factorial      = flag.Bool("f", false, "Calculate factorial")
    fibonacci      = flag.Bool("F", false, "Calculate fibonacci sequence")
    startStr       = flag.String("s", "1", "Start value")
    endStr         = flag.String("e", "100", "End value")
    hexResult      = flag.Bool("h", false, "Show hexadecimal result")
    onlyResult     = flag.Bool("r", false, "Show only the result")
    valueChars     = flag.Bool("v", false, "Show amount of characters in decimal and hexadecimal")
)

func main() {
    flag.Parse()

    startVal, ok := new(big.Int).SetString(*startStr, 10)
    if !ok {
        fmt.Println("Error: Invalid start value.")
        return
    }

    endVal, ok := new(big.Int).SetString(*endStr, 10)
    if !ok {
        fmt.Println("Error: Invalid end value.")
        return
    }

    // Check if start value is greater than end value
    if startVal.Cmp(endVal) > 0 {
        fmt.Println("Error: Start value cannot be greater than end value.")
        return
    }

    // Use a reasonable maximum for the end value
    maxEndVal := new(big.Int).Exp(big.NewInt(10), big.NewInt(1000), nil)

    // Set the end value to the maximum if it exceeds the maximum
    if endVal.Cmp(maxEndVal) > 0 {
        endVal.Set(maxEndVal)
        fmt.Printf("Warning: End value exceeds maximum allowed. Setting end value to %s.\n", maxEndVal.String())
    }

    // Calculate from start value to end value
    for i := new(big.Int).Set(startVal); i.Cmp(endVal) <= 0; i.Add(i, big.NewInt(1)) {
        var result big.Int
        switch {
        case *addition:
            result.Add(i, i)
        case *multiplication:
            result.Mul(i, i)
        case *power:
            result.Exp(i, i, nil)
        case *factorial:
            result.MulRange(1, int64(i.Int64()))
        case *fibonacci:
            result.Set(fib(i.Int64()))
        }

        if *onlyResult && *hexResult {
            fmt.Printf("%X\n", &result)
            continue
        }

        if *onlyResult {
            fmt.Println(result.String())
            continue
        }

        if *hexResult {
            hexResult := fmt.Sprintf("%X", &result)
            fmt.Printf("%s", i)

            switch {
            case *addition:
                fmt.Print(" + ")
            case *multiplication:
                fmt.Print(" * ")
            case *power:
                fmt.Print("^")
            case *factorial:
                fmt.Print("!")
            case *fibonacci:
                fmt.Print(" Fibonacci")
            }

            fmt.Printf("%s = %s", i, hexResult)
            if *valueChars {
                fmt.Printf(" (%d characters)", len(hexResult))
            }
        } else {
            fmt.Printf("%s", i)

            switch {
            case *addition:
                fmt.Print(" + ")
            case *multiplication:
                fmt.Print(" * ")
            case *power:
                fmt.Print("^")
            case *factorial:
                fmt.Print("!")
            case *fibonacci:
                fmt.Print(" Fibonacci")
            }

            fmt.Printf("%s = %s", i, result.String())

            if *valueChars {
                fmt.Printf(" (%d digits)", len(result.String()))
            }
        }

        fmt.Println()
    }
}

func fib(n int64) *big.Int {
    fibNumbers := [2]*big.Int{big.NewInt(0), big.NewInt(1)}
    for i := int64(0); i < n; i++ {
        fibNumbers[i%2] = new(big.Int).Add(fibNumbers[0], fibNumbers[1])
    }
    return fibNumbers[n%2]
}

