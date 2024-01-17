package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-swift/internal/runner"
	"github.com/wasilibs/go-protoc-gen-swift/internal/wasm"
)

func main() {
	os.Exit(runner.Run("protoc-gen-swift", os.Args[1:], wasm.ProtocGenSwift, os.Stdin, os.Stdout, os.Stderr, "."))
}
