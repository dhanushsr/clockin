package utils

import "fmt"

func PrintError(err error) {
	fmt.Printf("Something went wrong. %s\n", err)
}
