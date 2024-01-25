package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"xit/utils"
)

// Init 初始化仓库，创建 .xit/objects 目录，.xit/refs/heads 目录，.xit/HEAD 文件
// 设置 .xit 为隐藏文件夹
// 无法重复初始化
func Init(cmd *cobra.Command, args []string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	xitDir := filepath.Join(dir, utils.RootDir)
	if _, err := os.Stat(xitDir); err == nil {
		fmt.Printf("!Already a xit repo - %s\n", xitDir)
		return
	}

	// 定义并创建 .xit 仓库的子目录
	dirs := []string{filepath.Join(xitDir, "objects"), filepath.Join(xitDir, "refs", "heads")}
	for _, d := range dirs {
		if err := os.MkdirAll(d, os.ModePerm); err != nil {
			fmt.Println(err)
			return
		}
	}

	// 在 .xit 仓库的根目录写入 HEAD 文件内容
	head := filepath.Join(xitDir, "HEAD")
	if err := os.WriteFile(head, []byte("ref: refs/heads/master\n"), os.ModePerm); err != nil {
		fmt.Println(err)
		return
	}

	// 设置 .xit 目录的隐藏(跨平台)
	setDirHidden(xitDir)
	fmt.Printf("Initialized empty xit repository in %s\n", dir)
}

// setDirHidden 设置目录为隐藏(在 Windows 系统下)
func setDirHidden(dir string) {
	cmd := exec.Command("attrib", "+H", dir)
	cmd.Run()
}
