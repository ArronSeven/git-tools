package branch

import (
	"fmt"
	"git-tools/internal"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Handler struct {
	Config internal.BranchConfig
}

// 切换分支处理
// switch branch
func (b *Handler) SwitchBranch() {
	// 目录不存在则不处理
	// create directory
	if s, err := os.Stat(b.Config.DirPath); os.IsNotExist(err) {
		fmt.Printf("check file:%s\n", err)
		return
	} else {
		if !s.IsDir() {
			fmt.Printf("The Dir Path:%s,Can only be a directory but found to be a file.\n", b.Config.DirPath)
			return
		}
	}
	//获取当前目录下的所有文件或目录信息
	// index project
	err := filepath.Walk(b.Config.DirPath, func(path string, info os.FileInfo, err error) error {
		if info != nil && info.Name() == ".git" {
			path = strings.ReplaceAll(path, "\\", "/")
			index := strings.LastIndex(path, "/")
			b.exec(path[:index])
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 远程分支名称
// get branch remote name
func (b *Handler) remoteName() string {
	var name string
	index := strings.Index(b.Config.BranchName, "/")
	if index == -1 {
		name = "origin/" + b.Config.BranchName
	} else {
		// 如果是 my/dev不是origin/dev 则截取后面的作为名称
		b.Config.BranchName = b.Config.BranchName[index:]
	}
	return name
}

// 执行切换分支命令行
// execute git branch command
func (b *Handler) exec(execPath string) {
	remoteName := b.remoteName()
	// 执行git命令
	cmd := exec.Command("git", "checkout", "-b", b.Config.BranchName, remoteName)
	// 输出的目录
	cmd.Dir = execPath
	err := internal.Exec(cmd)
	if err != nil {
		fmt.Printf("execute [git checkout -b] command occure err:%v\n branch name:%s\n remote branch name:%s \n project name:%s\n",
			err, b.Config.BranchName, remoteName, execPath)
		fmt.Println()
		return
	}
}
