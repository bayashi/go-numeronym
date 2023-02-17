package main

import (
	"fmt"
	"github.com/bayashi/go-numeronym"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(numeronym.Numeronize(strings.TrimRight(string(in), "\n")))
}
