package main

import (
	"fmt"
	"sort"
)

func main() {
	// There is a different sorting function for each type"
	vars := []int{5, 2, 0, 3, 4, 9, 6}
	sort.Ints(vars)
	fmt.Println(vars)

	vars2 := []string{"Golang", "was", "developed", "by", "Google"}
	sort.Strings(vars2)
	fmt.Println(vars2)
}
