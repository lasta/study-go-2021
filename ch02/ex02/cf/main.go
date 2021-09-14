package main

import (
	"bufio"
	"fmt"
	"me.lasta/study-go-2021/ch02/ex01/tempconv"
	"me.lasta/study-go-2021/ch02/ex02/astronomicalunit"
	"os"
	"strconv"
)

func main() {
	var inputs []string
	if len(os.Args) > 1 {
		inputs = os.Args[1:]
	} else {
		inputs = readLines(os.Stdin)
	}

	// TODO: impl logics
	for _, input := range inputs {
		convert(input)
	}
}

func readLines(f *os.File) []string {
	input := bufio.NewScanner(f)
	var lines []string
	for input.Scan() {
		lines = append(lines, input.Text())
	}
	return lines
}

func convert(input string) {
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		return
	}

	fmt.Printf("---------- input: %g ----------", value)
	fmt.Printf("%s = %s\n", input, tempconv.KToC(tempconv.Kelvin(value)))
	fmt.Printf("%s = %s\n", input, tempconv.FToC(tempconv.Fahrenheit(value)))
	fmt.Printf("%s = %s\n", input, tempconv.CToF(tempconv.Celsius(value)))
	fmt.Printf("%s = %s\n", input, tempconv.KToF(tempconv.Kelvin(value)))
	fmt.Printf("%s = %s\n", input, tempconv.CToK(tempconv.Celsius(value)))
	fmt.Printf("%s = %s\n", input, tempconv.FToK(tempconv.Fahrenheit(value)))
	fmt.Println()
	fmt.Printf("%s = %s\n", input, astronomicalunit.Meter(value).ToAU())
	fmt.Printf("%s = %s\n", input, astronomicalunit.Meter(value).ToLY())
	fmt.Printf("%s = %s\n", input, astronomicalunit.Meter(value).ToPC())
	fmt.Printf("%s = %s\n", input, astronomicalunit.AstronomicalUnit(value).ToMeter())
	fmt.Printf("%s = %s\n", input, astronomicalunit.LightYear(value).ToMeter())
	fmt.Printf("%s = %s\n", input, astronomicalunit.Parsec(value).ToMeter())
	fmt.Println()
}
