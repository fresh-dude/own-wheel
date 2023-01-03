package main

/*
通过系统调度相关函数，实现一个简单版本的shell。有几点需要注意：
1. 命令跟命令相关参数是分开传参的
2. cd 命令需要更换目录，需特殊处理

待完善点：
1. 通过键盘⬆️⬇️键查看历史输入
2. 完善输入框（用户名/host名/目录）
*/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")

		// 读取命令行的内容
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			// 当命令执行出错时，将错误输出到控制台
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// 定义没有地址时错误提示 （全局定义方便修改）
var ErrNoPath = errors.New("path required")

func execInput(input string) error {
	// 去除命令行末尾的换行符
	input = strings.TrimSuffix(input, "\n")

	// 通过空格区分命令及其参数
	args := strings.Split(input, " ")

	// cd 命令会更改目录所在位置，需要单独处理
	// exit 命令便于退出自定义的shell
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return ErrNoPath
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}
	// 区分命令和参数
	cmd := exec.Command(args[0], args[1:]...)

	// 指定输出的设备（原控制台）
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// 返回运行结果
	return cmd.Run()
}
