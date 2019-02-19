package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(scan)
	bo := bufio.NewWriter(os.Stdout)
	defer bo.Flush()
	for s.Scan() {
		_, err := bo.WriteString(s.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "write: %v\n", err)
			continue
		}
	}
}

func scan(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		i++
		if data[0] == '\t' {
			return i, data[1:i], nil
		}
		return i, data[0:i], nil
	}
	if atEOF {
		if data[0] == '\t' {
			return len(data), data[1:], nil
		}
		return len(data), data, nil
	}
	return 0, nil, nil
}
