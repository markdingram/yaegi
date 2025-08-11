package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []string{"Bat", "Fox", "Owl", "Fox"}
	slices.Sort(s)
	fmt.Println(s)
}

// Output:
// [Bat Fox Fox Owl]
