package main

import (
	"encoding/json"
	"fmt"
)

type Setting struct {
	ContainerName string   `json:"container_name"`
	RebuildCmds   []string `json:"rebuild_cmds"`
	TestLists     []string `json:"test_lists"`
}

func main() {
	loadSettingFile()
}

func loadSettingFile() {
	// TODO(shida): 設定ファイルから読み込むようにする。
	b := []byte(`{"container_name": "web", "rebuild_cmds": ["down", "build", "up -d"], "test_lists":["rspec", "rails test", "rubocop"]}`)
	var s Setting
	if err := json.Unmarshal(b, &s); err != nil {
		fmt.Println(err)
	}

	fmt.Println("container_name:", s.ContainerName)
	fmt.Println("rebuild_cmds:", s.RebuildCmds)
	fmt.Println("test_lists:", s.TestLists)
}
