# gocodegen - Go Code Generator

This package provides Go structures for constructs used in Go source code and a helper to write those constructs to `.go` file with valid syntax.

### Usage

1. Get the package

    ```bash
    go get github.com/xapi-tools/gocodegen@latest
    ```

2. Create Go structures for intended source code and generate file

    ```go
    package main

    import gfg "github.com/xapi-tools/gocodegen"

    func main() {
        gw := gfg.NewGoFileWriter()

        gw.ToFile("./test.go")
    }
    ```


    Check following for more comprehensive examples:
    - [Unit Test](gen_test.go)
    - [Output](testdata/example.go)
