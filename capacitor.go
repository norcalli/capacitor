package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

var prefixes []string

func init() {
	prefixes = []string{"p", "n", "u", "m", ""}
}

// return rounded version of x with prec precision.
func roundFloat(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	intermed += .5
	x = .5
	if frac < 0.0 {
		x = -.5
		intermed -= 1
	}
	if frac >= x {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

func getPrefix(value float64) string {
	fmt.Println(math.Log10(value))
	return prefixes[int(math.Log10(value))/3+1]
}

func simplify(value float64) float64 {
	if roundFloat(value, 14) == value {
		return roundFloat(value, 14)
	}
	return value
}

func Usage() {
	flag.PrintDefaults()
	fmt.Printf("Usage: %s <code>\n", os.Args[0])
	os.Exit(0)
}

func main() {
	flag.Parse()
	code := flag.Arg(0)
	if code == "" {
		Usage()
	}

	if _, err := strconv.ParseUint(code, 10, 64); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		Usage()
	}
	base, _ := strconv.ParseInt(code[:2], 10, 16)
	exponent := 0
	if code[2:] != "" {
		tmp, _ := strconv.ParseInt(code[2:], 10, 16)
		exponent = int(tmp)
	}
	exponent++

	prefix := prefixes[exponent/3]
	value := simplify(float64(base) * math.Pow10(exponent%3))
	fmt.Printf("%g %sF\n", value, prefix)
	if value < 100 {
		return
	}
	value /= 1000
	prefix = prefixes[exponent/3+1]
	fmt.Printf("%g %sF\n", value, prefix)
}
