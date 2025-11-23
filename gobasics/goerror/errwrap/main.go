package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("1.txt")

	if err != nil {
		fmt.Printf("%T\n", err)

		e := fmt.Errorf("failed to open file: %w", err)
		fmt.Printf("%T\n", e)
		unwrappedErr := errors.Unwrap(e)
		fmt.Printf("%v\n", unwrappedErr)

	}

	defer f.Close()
}
