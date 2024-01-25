package main

import (
	"github.com/spf13/cobra"
	"xit/commands"
)

func Run() {
	root := &cobra.Command{
		Use: "xit",
	}

	init := &cobra.Command{
		Use:   "init",
		Short: "初始化仓库",
		Long:  "初始化仓库",
		Args:  cobra.MinimumNArgs(0),
		Run:   commands.Init,
	}

	add := &cobra.Command{
		Use:   "add",
		Short: "添加文件到暂存区",
		Long:  "添加文件到暂存区",
		Args:  cobra.MinimumNArgs(1),
	}

	rm := &cobra.Command{
		Use:   "rm",
		Short: "删除文件",
		Long:  "删除文件",
		Args:  cobra.MinimumNArgs(1),
	}

	commit := &cobra.Command{
		Use:   "commit",
		Short: "提交暂存区的文件",
		Long:  "提交暂存区的文件",
		Args:  cobra.MinimumNArgs(1),
	}

	status := &cobra.Command{
		Use:   "status",
		Short: "查看当前状态",
		Long:  "查看当前状态",
		Args:  cobra.MinimumNArgs(1),
	}

	log := &cobra.Command{
		Use:   "log",
		Short: "log",
		Long:  "log",
		Args:  cobra.MinimumNArgs(1),
	}

	branch := &cobra.Command{
		Use:   "branch",
		Short: "branch",
		Long:  "branch",
		Args:  cobra.MinimumNArgs(1),
	}

	switchs := &cobra.Command{
		Use:   "switch",
		Short: "切换分支",
		Long:  "切换分支",
		Args:  cobra.MinimumNArgs(1),
	}

	restore := &cobra.Command{
		Use:   "restore",
		Short: "restore",
		Long:  "restore",
		Args:  cobra.MinimumNArgs(1),
	}

	merge := &cobra.Command{
		Use:   "merge",
		Short: "merge",
		Long:  "merge",
		Args:  cobra.MinimumNArgs(1),
	}

	root.AddCommand(init, add, rm, commit, status, log, branch, switchs, restore, merge)
	root.Execute()
}
