// +build linux

package main

import (
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
	w.Fprintf("tag", "Get ")
	return w
}

func main() {
	log.SetPrefix("Watch:")
	log.SetFlags(0)

	if len(os.Args) < 2 {
		log.Fatal("usage: Watch [args...]")
	}

	win := openWin()

	fd, err := unix.InotifyInit()
	if err != nil {
		log.Fatal(err)
	}
	_, err = unix.InotifyAddWatch(fd, ".", unix.IN_MODIFY)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	for {
		_, err = unix.Read(fd, evbuf[:])
		if err != nil {
			log.Fatal(err)
		}
		// event struct:
		// e := evbuf[:n]
		// ev := (*(*unix.InotifyEvent)(unsafe.Pointer(&e[0])))
		// file name: evbuf[unix.SizeofInotifyEvent:]
		r, w, err := os.Pipe()
		if err != nil {
			log.Fatal(err)
		}
		win.Addr(",")
		win.Write("data", nil)
		win.Ctl("clean")
		win.Fprintf("body", "$ %s\n", strings.Join(os.Args, " "))
		cmd.Stdout = w
		cmd.Stderr = w
		if err := cmd.Start(); err != nil {
			r.Close()
			w.Close()
			win.Fprintf("body", "%s: :%s\n", strings.Join(os.Args, " "), err)
			continue
		}
		w.Close()
		var buf [4096]byte
		_, err = r.Read(buf[:])
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Wait(); err != nil {
			win.Fprintf("body", "%s: %s\n", strings.Join(os.Args, " "), err)
		}
		win.Fprintf("body", "$\n")
		win.Fprintf("addr", "#0")
		win.Ctl("dot=addr")
		win.Ctl("show")
		win.Ctl("clean")
	}
}
