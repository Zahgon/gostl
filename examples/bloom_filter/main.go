package main

import (
	"fmt"
)

func main() {
	filter := bloom.New(100, 4, bloom.WithGoroutineSafe())
	filter.Add("hhhh")
	filter.Add("gggg")

	fmt.Printf("%v\n", filter.Contains("aaaa"))
	fmt.Printf("%v\n", filter.Contains("gggg"))
}
