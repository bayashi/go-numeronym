# go-numeronym

`go-numeronym` provides the function `Numeronize` to contract a word as Numeronym.

## Usage

    package main

    import (
        "fmt"
        "github.com/bayashi/go-numeronym"
    )

    func main() {
        fmt.Println(numeronym.Numeronize("internationalization")) // i18n
    }

## For CLI

You can use `n7m` command to make numeronym in several ways.

    $ n7m internationalization
    i18n

    $ echo internationalization | n7m
    i18n

    cat file.txt
    internationalization
    $ n7m < file.txt
    i18n

## Installation

    go get github.com/bayashi/go-numeronym

## License

MIT

## Author

Dai Okabayashi: @bayashi
