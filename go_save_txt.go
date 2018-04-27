package main

import (
    "os"
)

func main() {

	file, err := os.Create("output.txt")
	if err != nil {
		return
	}

	defer file.Close()

	file.WriteString("12345")
	
}
