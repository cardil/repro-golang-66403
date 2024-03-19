package main

import (
	"fmt"

	"github.com/buildkite/agent/v3/env"
)

func main() {
	n, v, ok := env.Split("a=b")
	fmt.Println(n, v, ok)
}
