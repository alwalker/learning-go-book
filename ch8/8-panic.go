package main

import "os"

func doPanic(msg string) {
    panic("BALLS" + msg)
}

func main() {
    doPanic(os.Args[0])
}
