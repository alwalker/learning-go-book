// Package for testing different methods for wrapping/type checking errors
package main

import (
	"errors"
	"fmt"
	"os"
)

// WrappedFileChecker Generates a new error that wraps the OS error
func WrappedFileChecker(name string) error {
    f, err := os.Open(name)
    if err != nil {
        return fmt.Errorf("in fileChecker: %w", err)
    }
    f.Close()
    return nil
}

// NonWrappedFileChecker Generates a new error that embeds the OS error
func NonWrappedFileChecker(name string) error {
    f, err := os.Open(name)
    if err != nil {
        return fmt.Errorf("in fileChecker: %v", err)
    }
    f.Close()
    return nil
}

// main runs both the wrapped and non wrapped error functions as well as
// running the wrapped function a 2nd time to check for
//	error.Is()
func main() {
    fmt.Println("Wrapped Error")
    err := WrappedFileChecker("not_here.txt")
    if err != nil {
        fmt.Println(err)
        if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
            fmt.Println(wrappedErr)
        }
    }

    fmt.Println("\nNon Wrapper Error")
    err = NonWrappedFileChecker("not_here.txt")
    if err != nil {
        fmt.Println(err)
        if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
            fmt.Println(wrappedErr)
        }
    }

    fmt.Println("\nError.is")
    err = WrappedFileChecker("not_here.txt")
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            fmt.Println("That file doesn't exist")
        }
    }
}
