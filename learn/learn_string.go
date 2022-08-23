package main

import (
	"fmt"
	"strconv"
	"strings"
)
func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

	str = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))

	str = "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %+v\n", sl)

	for _,val := range sl {
		fmt.Printf("%s - ", val)
	}

	fmt.Println()

	val, _ := strconv.Atoi("10")

	fmt.Printf("string to int is %v\n", val)

}