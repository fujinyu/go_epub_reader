package epub

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"path"
)

// Book 表示一个 EPUB 电子书。
type Book struct {
	Ncx       Ncx             `json:"ncx"` // 电子书的目录文件。
	Opf       Opf             `json:"opf"` // 电子书的 OPF 文件。
	Container Container       `json:"-"`   // 电子书的容器文件。
	Mimetype  string          `json:"-"`   // 电子书的 MIME 类型。
	fd        *zip.ReadCloser // 电子书的 ZIP 文件读取器。
}

// Open 打开电子书中的指定文件，并返回一个 io.ReadCloser 接口。
func (p *Book) Open(n string) (io.ReadCloser, error) {
	return p.open(p.filename(n))
}

// Files 返回电子书中所有文件的文件名列表。
func (p *Book) Files() []string {
	var fns []string
	for _, f := range p.fd.File {
		fns = append(fns, f.Name)
	}
	return fns
}

// Close 关闭电子书的 ZIP 文件读取器。
func (p *Book) Close() {
	p.fd.Close()
}

// filename 返回指定文件名在电子书中的完整路径。
func (p *Book) filename(n string) string {
	return path.Join(path.Dir(p.Container.Rootfile.Path), n)
}

// readXML 读取指定文件名的 XML 文件，并将其解码为指定类型的值。
func (p *Book) readXML(n string, v interface{}) error {
	fd, err := p.open(n)
	if err != nil {
		return nil
	}
	defer fd.Close()
	dec := xml.NewDecoder(fd)
	return dec.Decode(v)
}

// readBytes 读取指定文件名的字节数据。
func (p *Book) readBytes(n string) ([]byte, error) {
	fd, err := p.open(n)
	if err != nil {
		return nil, nil
	}
	defer fd.Close()

	return ioutil.ReadAll(fd)

}

// open 打开指定文件名的文件，并返回一个 io.ReadCloser 接口。
func (p *Book) open(n string) (io.ReadCloser, error) {
	for _, f := range p.fd.File {
		if f.Name == n {
			return f.Open()
		}
	}
	return nil, fmt.Errorf("file %s not exist", n)
}
