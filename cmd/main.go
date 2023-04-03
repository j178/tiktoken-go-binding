package main

import (
	"fmt"
	"io"
	"os"

	tiktoken_go "github.com/j178/tiktoken-go"
)

func main() {
	in, _ := io.ReadAll(os.Stdin)
	count := tiktoken_go.CountTokens("gpt-3.5-turbo", string(in))
	fmt.Println(count)
}
