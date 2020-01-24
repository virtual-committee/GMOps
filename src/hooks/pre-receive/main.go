package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	os.Exit(1)
}
