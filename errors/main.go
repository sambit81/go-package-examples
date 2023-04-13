package main

import (
	"errors"
	"fmt"
)

func process(i int) error {
	if i%2 == 0 {
		return errors.New("only odd numbers allowed")
	}
	return nil
}
func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("Operation Successful")
}
func main() {
	err := process(3)
	checkError(err)
	err = process(2)
	checkError(err)
}
