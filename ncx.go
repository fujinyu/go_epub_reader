package epub

// Ncx 表示 EPUB 电子书的 OPS/toc.ncx 文件。
type Ncx struct {
	Points []NavPoint `xml:"navMap>navPoint" json:"points"` // 电子书的目录结构。
}

// NavPoint 表示电子书目录中的一个导航点。
type NavPoint struct {
	Text    string     `xml:"navLabel>text" json:"text"` // 导航点的文本内容。
	Content Content    `xml:"content" json:"content"`    // 导航点的内容。
	Points  []NavPoint `xml:"navPoint" json:"points"`    // 导航点的子导航点。
}

// Content 表示导航点的内容。
type Content struct {
	Src string `xml:"src,attr" json:"src"` // 内容的路径。
}
