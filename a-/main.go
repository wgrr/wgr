package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bio := bufio.NewScanner(os.Stdin)
	for bio.Scan() {
		l := bio.Text()
		// instead of this logic, i could just implement
		// a SpliFunc to scanner, just don't bother for now.
		if len(l) > 0 && l[0] == '\t' {
			fmt.Printf("%s\n", l[1:])
		} else {
			fmt.Printf("%s\n", l)
		}
	}
}
