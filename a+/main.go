package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bio := bufio.NewScanner(os.Stdin)
	for bio.Scan() {
		fmt.Printf("\t%s\n", bio.Text())
	}
}
