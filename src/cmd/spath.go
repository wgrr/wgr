package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func usage() {
	log.Print("path")
	flag.PrintDefaults()
}

func main() {
	log.SetPrefix("spath: ")
	log.SetFlags(0)
	flag.Parse()

	if len(flag.Args()) < 1 {
		usage()
		os.Exit(1)
	}
	higher, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Stat(higher)
	if err != nil {
		log.Fatal(err)
	}
	if !f.IsDir() {
		log.Fatalf("%s is not a directory", f.Name())
	}
	path := os.Getenv("PATH")
	if !strings.Contains(path, higher) {
		fmt.Printf("%s:%s", higher, path)
		return
	}
	re, err := regexp.Compile(":?" + higher + ":?")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s:%s", higher, re.ReplaceAllLiteralString(path, ""))
}
