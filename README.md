# go-numeronym

`go-numeronym` provides the function `Numeronize` to contract a word as Numeronym.

## Usage

```go
package main

import (
    "fmt"
    "github.com/bayashi/go-numeronym"
)

func main() {
    fmt.Println(numeronym.Numeronize("internationalization")) // i18n
}
```

## For CLI

You can use `n7m` command to make numeronym in several ways.

    $ n7m internationalization
    i18n

    $ echo internationalization | n7m
    i18n

    $ cat file.txt
    internationalization
    $ n7m < file.txt
    i18n

### Installation of `n7m`

    go install github.com/bayashi/go-numeronym/cmd/n7m@latest

Binaries are here: https://github.com/bayashi/go-numeronym/releases

## License

MIT

## Author

Dai Okabayashi: @bayashi
