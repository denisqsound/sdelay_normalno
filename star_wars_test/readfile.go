package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	file.Close()

}
