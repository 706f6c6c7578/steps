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
	start          = flag.Int("s", 0, "Start value")
	end            = flag.Int("e", 0, "End value")
	hash           = flag.Bool("h", false, "Calculate SHA256 hash")
	base64hash     = flag.Bool("b", false, "Calculate base64 encoded SHA256 binary hash")
)

func main() {
	flag.Parse()

	for i := *start; i <= *end; i++ {
		var result big.Int
		switch {
		case *addition:
			result.Add(big.NewInt(int64(i)), big.NewInt(int64(i)))
			fmt.Printf("%d + %d = ", i, i)
		case *multiplication:
			result.Mul(big.NewInt(int64(i)), big.NewInt(int64(i)))
			fmt.Printf("%d * %d = ", i, i)
		case *power:
			result.Exp(big.NewInt(int64(i)), big.NewInt(int64(i)), nil)
			fmt.Printf("%d^%d = ", i, i)
		}

		fmt.Printf("%s", &result)

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

