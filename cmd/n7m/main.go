package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/bayashi/go-numeronym"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	for _, arg := range os.Args {
		fmt.Println(numeronym.Numeronize(strings.TrimRight(arg, "\n")))
	}

	if terminal.IsTerminal(syscall.Stdin) && !terminal.IsTerminal(0) {
		in, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(numeronym.Numeronize(strings.TrimRight(string(in), "\n")))
	} else {
		fmt.Println("Error: Could not convert from a file.")
	}
}
