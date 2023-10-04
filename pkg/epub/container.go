package epub

// Container 表示 EPUB 电子书的 META-INF/container.xml 文件。
type Container struct {
	Rootfile Rootfile `xml:"rootfiles>rootfile" json:"rootfile"` // 电子书的根文件。
}

// Rootfile 表示 EPUB 电子书的根文件。
type Rootfile struct {
	Path string `xml:"full-path,attr" json:"path"`  // 根文件的路径。
	Type string `xml:"media-type,attr" json:"type"` // 根文件的 MIME 类型。
}
