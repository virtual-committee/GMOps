package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	user := os.Getenv("GMOPS_USER")
	repoId := os.Getenv("GMOPS_REPO_ID")

	input, _ := ioutil.ReadAll(os.Stdin)
	param := strings.Fields(string(input))
	if len(param) != 3 {
		os.Exit(1)
	}

	fmt.Println(user)
	fmt.Println(repoId)

	os.Exit(1)
}
