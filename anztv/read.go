package main

import (
	"fmt"
	"strings"

	"github.com/will-nb/go_epub_reader/pkg/epub"
)

func (r *Repository) Open() {
	filepath := r.args[2]
	if len(filepath) == 0 {
		fmt.Println("请指定要打开的文件")
		return
	}
	bk, err := epub.Open(filepath)
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	defer bk.Close()
	printMetaData(bk)
}

func printMetaData(bk *epub.Book) {
	metaData := bk.Opf.Metadata
	fmt.Printf("书名: %s\n", strings.Join(metaData.Title, ""))
	fmt.Printf("作者: %s\n", joinAuthors(metaData.Creator))
	fmt.Printf("语言: %s\n", strings.Join(metaData.Language, ""))
	fmt.Printf("出版社: %s\n", strings.Join(metaData.Publisher, ""))
	fmt.Printf("出版时间: %s\n", joinDates(metaData.Date))
	fmt.Printf("简介: %s\n", strings.Join(metaData.Description, ""))
	fmt.Println("目录:")
	for _, v := range bk.Ncx.Points {
		fmt.Println(v.Text)
	}
}

func joinAuthors(authors []epub.Author) string {
	var authorNames []string
	for _, author := range authors {
		authorNames = append(authorNames, author.Data)
	}
	return strings.Join(authorNames, ", ")
}

func joinDates(dates []epub.Date) string {
	var dateStrings []string
	for _, date := range dates {
		dateStrings = append(dateStrings, date.Data)
	}
	return strings.Join(dateStrings, ", ")
}

func (r *Repository) Read() {
	//读取一个指定章节的内容
	filepath := r.args[2]
	if len(filepath) == 0 {
		fmt.Println("请指定要打开的文件")
		return
	}
	bk, err := epub.Open(filepath)
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	defer bk.Close()
	//读取一个指定章节的内容
	chapterName := r.args[3]
	if len(chapterName) == 0 {
		fmt.Println("请指定要打开的章节")
		return
	}
	chapterContent, err := bk.ReadChapter(chapterName)
	if err != nil {
		fmt.Println("打开章节失败", err)
	}
	fmt.Println(chapterContent)
}

func (r *Repository) ReadAll() {
	//读取所有章节的内容
	filepath := r.args[2]
	if len(filepath) == 0 {
		fmt.Println("请指定要打开的文件")
		return
	}
	bk, err := epub.Open(filepath)
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	defer bk.Close()
	content, err := bk.ReadAll()
	if err != nil {
		fmt.Println("打开章节失败", err)
	}
	fmt.Println(content)
}
