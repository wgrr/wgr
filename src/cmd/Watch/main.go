// +build linux

package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"9fans.net/go/acme"
	"golang.org/x/sys/unix"
)

var evbuf [8192]byte

func openWin() *acme.Win {
	w, err := acme.New()
	if err != nil {
		log.Fatal(err)
	}
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	w.Name(filepath.Join(wd, "+watch"))
	w.Ctl("clean")
	return w
}

func main() {
	log.SetPrefix("Watch: ")
	log.SetFlags(0)
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("usage: args...")
	}

	fd, err := unix.InotifyInit()
	if err != nil {
		log.Fatal(err)
	}
	_, err = unix.InotifyAddWatch(fd, ".", unix.IN_MODIFY)
	if err != nil {
		log.Fatal(err)
	}

	win := openWin()
	go func() {
		del := []byte("Del")
		for e := range win.EventChan() {
			if e.C2 == 'x' || e.C2 == 'X' {
				if bytes.Equal(e.Text, del) {
					win.Ctl("delete")
				}
			}
			win.WriteEvent(e)
		}
		os.Exit(0)
	}()
	for {
		_, err = unix.Read(fd, evbuf[:])
		if err != nil {
			log.Fatal(err)
		}
		// event struct:
		// e := evbuf[:n]
		// ev := (*(*unix.InotifyEvent)(unsafe.Pointer(&e[0])))
		// file name: evbuf[unix.SizeofInotifyEvent:]
		win.Addr(",")
		win.Write("data", nil)
		win.Ctl("clean")
		win.Fprintf("body", "$ %s\n", strings.Join(args, " "))
		var w bytes.Buffer
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = &w
		cmd.Stderr = &w
		if err := cmd.Start(); err != nil {
			win.Fprintf("body", "%s: :%s\n", strings.Join(args, " "), err)
			continue
		}
		if err := cmd.Wait(); err != nil {
			win.Fprintf("body", "%s: %s\n", strings.Join(args, " "), err)
		}
		win.Write("body", w.Bytes())
		win.Fprintf("body", "$\n")
		win.Fprintf("addr", "#0")
		win.Ctl("dot=addr")
		win.Ctl("show")
		win.Ctl("clean")
	}
}
