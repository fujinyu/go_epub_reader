package epub

import (
	"errors"
	"io/ioutil"
	"strings"
)

// 返回一个章节的内容
func (p *Book) ReadChapter(n string) (string, error) {
	if p == nil {
		return "", errors.New("nil pointer receiver")
	}

	// 如果 n 是一个标题或文件链接，则打开对应的章节内容
	for _, point := range p.Ncx.Points {
		if point.Text == n || strings.Contains(point.Content.Src, n) {
			src := point.Content.Src
			if strings.Contains(src, "#") {
				parts := strings.Split(src, "#")
				if len(parts) != 2 {
					return "", errors.New("路径不止一个锚点," + src)
				}
				src = parts[0]
			}
			content, _ := p.readFile(src)
			return content, nil
		}
	}

	// 如果 n 既不是一个文件名也不是一个标题，则返回错误
	return "", errors.New("找不到该文件," + n)
}

// 返回一个文件的内容
func (p *Book) readFile(n string) (string, error) {
	src := p.filename(n)
	fd, err := p.open(src)
	if err != nil {
		return "", err
	}
	defer fd.Close()
	b, err := ioutil.ReadAll(fd)
	if err != nil {
		return "", err
	}

	// 返回整个文件内容
	return string(b), nil
}

// 返回所有章节的内容
func (p *Book) ReadAll() (string, error) {
	if p == nil {
		return "", errors.New("nil pointer receiver")
	}

	var content string
	readFiles := make(map[string]bool)
	var readAll func(points []NavPoint) error
	readAll = func(points []NavPoint) error {
		for _, point := range points {
			if point.Content.Src != "" {
				src := point.Content.Src
				if strings.Contains(src, "#") {
					parts := strings.Split(src, "#")
					if len(parts) != 2 {
						return errors.New("路径不止一个锚点:" + src)
					}
					src = parts[0]
				}
				if readFiles[src] {
					continue
				}
				readFiles[src] = true
				chapter, err := p.readFile(src)
				if err != nil {
					return err
				}
				content += chapter
			}
			if len(point.Points) > 0 {
				if err := readAll(point.Points); err != nil {
					return err
				}
			}
		}
		return nil
	}

	if err := readAll(p.Ncx.Points); err != nil {
		return "", err
	}

	return content, nil
}
