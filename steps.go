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
	startStr       = flag.String("s", "0", "Start value")
	endStr         = flag.String("e", "0", "End value")
	hash           = flag.Bool("h", false, "Calculate SHA256 hash")
	base64hash     = flag.Bool("b", false, "Calculate base64 encoded SHA256 binary hash")
	showDigits     = flag.Bool("d", false, "Show number of digits")
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

	for i := new(big.Int).Set(startVal); i.Cmp(endVal) <= 0; i.Add(i, big.NewInt(1)) {
		var result big.Int
		switch {
		case *addition:
			result.Add(i, i)
			fmt.Printf("%s + %s = ", i, i)
		case *multiplication:
			result.Mul(i, i)
			fmt.Printf("%s * %s = ", i, i)
		case *power:
			result.Exp(i, startVal, nil)
			fmt.Printf("%s^%s = ", i, startVal)
		}

		resultStr := result.String()
		fmt.Printf("%s", resultStr)

		if *showDigits {
			fmt.Printf(" (%d digits)", len(resultStr))
		}

		if *hash || *base64hash {
			h := sha256.New()
			h.Write(result.Bytes())
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
