package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

func Exec(cmd *exec.Cmd) error {
	//显示运行的命令
	fmt.Println(cmd.Args, cmd.Dir)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		s := string(tmp)
		if strings.Contains(s, "fatal") {
			fmt.Print(s)
		}
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
