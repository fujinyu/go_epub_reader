package epub

// Opf 表示 EPUB 电子书的 content.opf 文件。
type Opf struct {
	Metadata Metadata   `xml:"metadata" json:"metadata"`      // 电子书的元数据。
	Manifest []Manifest `xml:"manifest>item" json:"manifest"` // 电子书的文件清单。
	Spine    Spine      `xml:"spine" json:"spine"`            // 电子书的阅读顺序。
}

// Metadata 表示 EPUB 电子书的元数据。
type Metadata struct {
	Title       []string     `xml:"title" json:"title"`             // 电子书的标题。
	Language    []string     `xml:"language" json:"language"`       // 电子书的语言。
	Identifier  []Identifier `xml:"identifier" json:"identifier"`   // 电子书的标识符。
	Creator     []Author     `xml:"creator" json:"creator"`         // 电子书的作者。
	Subject     []string     `xml:"subject" json:"subject"`         // 电子书的主题。
	Description []string     `xml:"description" json:"description"` // 电子书的描述。
	Publisher   []string     `xml:"publisher" json:"publisher"`     // 电子书的出版社。
	Contributor []Author     `xml:"contributor" json:"contributor"` // 电子书的贡献者。
	Date        []Date       `xml:"date" json:"date"`               // 电子书的日期。
	Type        []string     `xml:"type" json:"type"`               // 电子书的类型。
	Format      []string     `xml:"format" json:"format"`           // 电子书的格式。
	Source      []string     `xml:"source" json:"source"`           // 电子书的来源。
	Relation    []string     `xml:"relation" json:"relation"`       // 电子书的关系。
	Coverage    []string     `xml:"coverage" json:"coverage"`       // 电子书的覆盖范围。
	Rights      []string     `xml:"rights" json:"rights"`           // 电子书的版权信息。
	Meta        []Metafield  `xml:"meta" json:"meta"`               // 电子书的元数据字段。
}

// Identifier 表示电子书的标识符。
type Identifier struct {
	Data   string `xml:",chardata" json:"data"`     // 标识符的值。
	ID     string `xml:"id,attr" json:"id"`         // 标识符的 ID。
	Scheme string `xml:"scheme,attr" json:"scheme"` // 标识符的方案。
}

// Author 表示电子书的作者。
type Author struct {
	Data   string `xml:",chardata" json:"author"`     // 作者的名称。
	FileAs string `xml:"file-as,attr" json:"file_as"` // 作者的文件名。
	Role   string `xml:"role,attr" json:"role"`       // 作者的角色。
}

// Date 表示电子书的日期。
type Date struct {
	Data  string `xml:",chardata" json:"data"`   // 日期的值。
	Event string `xml:"event,attr" json:"event"` // 日期的事件。
}

// Metafield 表示电子书的元数据字段。
type Metafield struct {
	Name    string `xml:"name,attr" json:"name"`       // 字段的名称。
	Content string `xml:"content,attr" json:"content"` // 字段的内容。
}

// Manifest 表示电子书的文件清单。
type Manifest struct {
	ID           string `xml:"id,attr" json:"id"`                   // 文件的 ID。
	Href         string `xml:"href,attr" json:"href"`               // 文件的路径。
	MediaType    string `xml:"media-type,attr" json:"type"`         // 文件的 MIME 类型。
	Fallback     string `xml:"media-fallback,attr" json:"fallback"` // 文件的备用路径。
	Properties   string `xml:"properties,attr" json:"properties"`   // 文件的属性。
	MediaOverlay string `xml:"media-overlay,attr" json:"overlay"`   // 文件的媒体叠加。
}

// Spine 表示电子书的阅读顺序。
type Spine struct {
	ID              string      `xml:"id,attr" json:"id"`                                  // 阅读顺序的 ID。
	Toc             string      `xml:"toc,attr" json:"toc"`                                // 目录文件的路径。
	PageProgression string      `xml:"page-progression-direction,attr" json:"progression"` // 页面进度方向。
	Items           []SpineItem `xml:"itemref" json:"items"`                               // 阅读顺序的项目。
}

// SpineItem 表示电子书阅读顺序中的一个项目。
type SpineItem struct {
	IDref      string `xml:"idref,attr" json:"id_ref"`          // 项目的 ID。
	Linear     string `xml:"linear,attr" json:"linear"`         // 项目的线性属性。
	ID         string `xml:"id,attr" json:"id"`                 // 项目的 ID。
	Properties string `xml:"properties,attr" json:"properties"` // 项目的属性。
}
