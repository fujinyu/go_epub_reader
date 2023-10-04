package main

import (
	"fmt"
	"time"
)

// help 函数实现了 help 命令的功能
func (r *Repository) Help() {
	fmt.Println("用法: anztv <命令>")
	fmt.Println()
	fmt.Println("可用命令:")
	fmt.Println("  help    显示此帮助消息")
	fmt.Println("  version  显示版本信息")
	fmt.Println("  open <file path>    打开epub文件")
	fmt.Println("  read <file path> <chapter name>    读取指定章节的内容")
	fmt.Println("  readall <file path>    读取所有章节的内容")
}

func (r *Repository) Version() {
	year := time.Now().Year()
	fmt.Printf("anztv version 0.0.1 © %d www.anz.tv\n", year)
}
