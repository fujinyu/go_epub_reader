package epub

import "archive/zip"

// Open 打开指定路径的 EPUB 电子书文件，并返回一个 Book 结构体指针。
func Open(fn string) (*Book, error) {
	// 打开 ZIP 文件读取器。
	fd, err := zip.OpenReader(fn)
	if err != nil {
		return nil, err
	}

	// 创建 Book 结构体。
	bk := Book{fd: fd}

	// 读取电子书的 MIME 类型。
	mt, err := bk.readBytes("mimetype")
	if err == nil {
		bk.Mimetype = string(mt)
		// 读取电子书的容器文件。
		err = bk.readXML("META-INF/container.xml", &bk.Container)
	}
	if err == nil {
		// 读取电子书的 OPF 文件。
		err = bk.readXML(bk.Container.Rootfile.Path, &bk.Opf)
	}

	// 查找电子书的目录文件。
	for _, mf := range bk.Opf.Manifest {
		if mf.ID == bk.Opf.Spine.Toc {
			err = bk.readXML(bk.filename(mf.Href), &bk.Ncx)
			break
		}
	}

	// 如果出现错误，则关闭 ZIP 文件读取器并返回错误信息。
	if err != nil {
		fd.Close()
		return nil, err
	}

	// 返回 Book 结构体指针。
	return &bk, nil
}
