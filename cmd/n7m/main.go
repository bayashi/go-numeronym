package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/bayashi/go-numeronym"
	"golang.org/x/term"
)

func main() {
	for i, arg := range os.Args {
		if i == 0 || arg == "-" {
			continue
		}
		fmt.Println(numeronym.Numeronize(strings.TrimRight(arg, "\n")))
	}

	if !term.IsTerminal(int(syscall.Stdin)) {
		in, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(numeronym.Numeronize(strings.TrimRight(string(in), "\n")))
	}
}
