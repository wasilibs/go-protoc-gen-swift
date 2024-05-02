package main

import (
	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	tasks.Define(tasks.Params{
		LibraryName: "protoc-gen-swift",
		LibraryRepo: "apple/swift-protobuf",
		GoReleaser:  true,
	})
	boot.Main()
}
