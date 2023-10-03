# EPUB 电子书解析器

这是一个用 Go 语言编写的 EPUB 电子书解析器，可以读取 EPUB 文件并提取出其中的元数据、文件清单和阅读顺序等信息。

## 安装

要使用这个库，你需要先安装 Go 语言的开发环境。然后，使用以下命令进行安装：

```sh
go get github.com/will-nb/go_epub-reader
```

## 使用

这是一个golang库，你可以在你的 Go 项目中引入这个库，然后使用其中的函数和结构体来解析 EPUB 文件。以下是一个简单的示例代码：

```go
package main

import (
	"fmt"
	"github.com/will-nb/go_epub-reader"
)

func main() {
	book, err := epub.Open("/path/to/book.epub")
	if err != nil {
		panic(err)
	}

	fmt.Println("Title:", book.Opf.Metadata.Title)
	fmt.Println("Author:", book.Opf.Metadata.Creator)
	fmt.Println("Language:", book.Opf.Metadata.Language)
	fmt.Println("Identifier:", book.Opf.Metadata.Identifier)
}
```

这将输出 EPUB 文件的元数据信息。

## 示例

以下是一个示例 EPUB 文件的解析结果：

```
Title: Example Book
Author: John Doe
Language: en
Identifier: urn:uuid:12345678-1234-5678-1234-567812345678
```

## 贡献

如果你发现了 bug 或者有改进的建议，欢迎提交 issue 或者 pull request。

## 许可证

这个项目使用 MIT 许可证。详情请参阅 LICENSE 文件。



## Epub格式介绍
- <http://idpf.org/epub>
- <http://www.cnblogs.com/Alex80/p/5127104.html>
- <http://www.cnblogs.com/diligenceday/p/4999315.html>

## 备注
- CHANGELOG.md 通过以下命令导出：
```
git log --pretty=format:"- %s (%ad)" --date=iso > CHANGELOG.md
```
- 本项目fork自<github.com/kapmahc/epub>，感谢原作者的贡献。但是原作者从2016年以来没有再更新，所以我将不考虑向原作者提交pull request。
- 当前本项目主要考虑为gutenberg.org提供epub文件的兼容性，来源于其他制作者的epub文件有可能出错。