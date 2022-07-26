package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestPathExists(t *testing.T) {
	path := "/Users/sam/workspace/cmkj/go_web_project"
	b := PathExists(path)
	if !b {
		t.Errorf("result incorrect")
	}
	path = "/Users/haha"
	b = PathExists(path)
	if b {
		t.Errorf("result incorrect")
	}
}

func TestExecPath(t *testing.T) {
	execpath, err := os.Executable() // 获得程序路径
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(execpath)
	fmt.Println(dir)

	s, _ := os.Getwd()
	println(s)

	println(ProjectPath())
}
