package main

import (
	"fmt"
	"strings"
)

func main() {

	//Contains(Original string, String to be searched)
	str := "Golang is developed by Google"
	ans1 := strings.Contains(str, "Golang") // true
	ans2 := strings.Contains(str, "golang") // false (The string should be exactly equal)
	fmt.Println(ans1)
	fmt.Println(ans2)

	//Count(Original string, string whose presence is to be counted)
	str1 := "Golang Golang Golang"
	ans3 := strings.Count(str1, "Golang")
	fmt.Println(ans3) // 3, (Count the number of the occurances of string 'Golang')

	//Replace(Original string, old string, new string)
	str2 := "Folang is developed by Foogle"
	ans4 := strings.ReplaceAll(str2, "Fo", "Go") // Golang is developed by Google (The term Fo will be replaced by Go)
	fmt.Println(ans4)

}
