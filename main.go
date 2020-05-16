package main

import (
	"fmt"
)

type Setting struct {
	ContainerName string   `json:"container_name"`
	RebuildCmds   []string `json:"rebuild_cmds"`
	TestLists     []string `json:"test_lists"`
}

func main() {
	fmt.Println("Hello World!")
}
