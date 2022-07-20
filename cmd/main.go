package main

import (
	"fmt"
	"os"

	"github.com/moderncircuits/paket/cmd/runtime"
)

// These variables get set during link time. See Makefile.
var (
	hostOS string
	commit string
	date   string
)

func init() {
	runtime.BuildCommit = commit
	runtime.BuildDate = date
	runtime.BuildOS = hostOS
}

func main() {
	if err := execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
