package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
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

	if flag.NFlag() == 0 {
		fmt.Println("Usage: steps [flags]")
		fmt.Println("Flags:")
		flag.PrintDefaults()
		os.Exit(1)
	}

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

		// Print only the result if the flag is set
		if *onlyResult {
			printResult(i, &result)
			fmt.Println()
			continue
		}

		// Otherwise, print the full expression
		printExpression(i, &result)
	}
}

func printResult(input, result *big.Int) {
	if *hexResult {
		fmt.Printf("%X", result)
	} else {
		fmt.Print(result.String())
	}
}

func printExpression(input, result *big.Int) {
	if *onlyResult {
		fmt.Printf("%s = ", input)
	}

	switch {
	case *addition:
		fmt.Printf("%s + %s", input, input)
	case *multiplication:
		fmt.Printf("%s * %s", input, input)
	case *power:
		fmt.Printf("%s ^ %s", input, input)
	case *factorial:
		fmt.Printf("%s!", input)
	case *fibonacci:
		fmt.Printf("F%s", input)
	}

	fmt.Print(" = ")
	printResult(input, result)

	if *valueChars {
		if *hexResult {
			fmt.Printf(" (%d characters)\n", len(fmt.Sprintf("%X", result)))
		} else {
			fmt.Printf(" (%d digits)\n", len(result.String()))
		}
	} else {
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

