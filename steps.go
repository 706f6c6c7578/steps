package main

import (
    "crypto/sha256"
    "encoding/base64"
    "flag"
    "fmt"
    "math/big"
)

var (
    addition       = flag.Bool("a", false, "Perform addition")
    multiplication = flag.Bool("m", false, "Perform multiplication")
    power          = flag.Bool("p", false, "Perform exponentiation")
    startStr       = flag.String("s", "1", "Start value")
    endStr         = flag.String("e", "100", "End value")
    hash           = flag.Bool("h", false, "Calculate SHA256 hash")
    base64hash     = flag.Bool("b", false, "Calculate base64 encoded SHA256 binary hash")
    showDigits     = flag.Bool("d", false, "Show number of digits")
    onlyResult     = flag.Bool("r", false, "Show only the result")
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
        }

        if *onlyResult {
            fmt.Println(result.String())
            continue
        }

        fmt.Printf("%s", i)

        switch {
        case *addition:
            fmt.Print(" + ")
        case *multiplication:
            fmt.Print(" * ")
        case *power:
            fmt.Print("^")
        }

        fmt.Printf("%s = %s", i, result.String())

        if *showDigits {
            fmt.Printf(" (%d digits)", len(result.String()))
        }

        if *hash || *base64hash {
            h := sha256.New()
            resultStr := result.String()
            h.Write([]byte(resultStr))
            if *hash {
                fmt.Printf(" SHA256: %x", h.Sum(nil))
            }
            if *base64hash {
                fmt.Printf(" Base64 (binary SHA256): %s", base64.StdEncoding.EncodeToString(h.Sum(nil)))
            }
        }

        fmt.Println()
    }
}
