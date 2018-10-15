// TODO:
// 		document, instead of os.File, io.Reader.
// 		maybe spread grerep into a new package so others
// 		can import
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

func grerep(r io.Reader, from, delim *regexp.Regexp) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("grerep: readall %v: %s", r, err)
	}
	loc1, loc2 := from.FindIndex(body), delim.FindIndex(body)
	switch {
	case loc1 == nil || loc2 == nil:
		return nil
	case loc1[0] > loc2[1]:
		return nil
	}
	_, err = os.Stdout.Write(body[loc1[0]:loc2[1]])
	if err != nil {
		return fmt.Errorf("grerep: write: %s", err)
	}
	return nil
}

func newRegexps(from, delim string) (*regexp.Regexp, *regexp.Regexp, error) {
	i, err := regexp.Compile(from)
	if err != nil {
		return nil, nil, fmt.Errorf("grerep: regexp %s: %s", from, err)
	}
	v, err := regexp.Compile(delim)
	if err != nil {
		return nil, nil, fmt.Errorf("grerep: regexp %s: %s", delim, err)
	}
	return i, v, nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "grerep: [fromre delimre] [file ...]\n")
	os.Exit(1)
}

func main() {
	args := os.Args[1:]
	l := len(args)
	switch {
	case l == 2:
		run(os.Stdin, args[0], args[1])
	case l > 2:
		for _, v := range args[2:] {
			f, err := os.Open(v)
			if err != nil {
				fmt.Fprintf(os.Stderr, "grerep: can't open: %s: %s\n", v, err)
				continue
			}
			run(f, args[0], args[1])
		}
	default:
		usage()
	}
}

func run(r io.Reader, fromRe, delimRe string) {
	from, delim, err := newRegexps(fromRe, delimRe)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	err = grerep(r, from, delim)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	os.Exit(0)
}