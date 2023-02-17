# go-numeronym

Provide the function `Numeronize` to contract a word as Numeronym.

## Usage

    package main

    import (
        "fmt"
        "github.com/bayashi/go-numeronym"
    )

    func main() {
        fmt.Println(Numeronize("internationalization")) // i18n
    }

## Installation

    go get github.com/bayashi/go-numeronym

## For CLI

You can use `n7m` command.

    $ echo 'internationalization' | n7m
    i18n

## License

MIT

## Author

Dai Okabayashi
