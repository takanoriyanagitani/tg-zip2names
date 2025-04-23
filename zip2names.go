package main

import (
	"archive/zip"
	"fmt"
	"io"
	"iter"
	"log"
	"os"
)

type ZipReader struct{ *zip.Reader }

func (z ZipReader) ToNames() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, file := range z.Reader.File {
			var hdr zip.FileHeader = file.FileHeader
			if !yield(hdr.Name) {
				return
			}
		}
	}
}

func (z ZipReader) PrintNames() error {
	for name := range z.ToNames() {
		_, e := fmt.Println(name)
		if nil != e {
			return e
		}
	}

	return nil
}

func envValByKey(key string) string { return os.Getenv(key) }

func zipFilename() string { return envValByKey("ENV_ZIP_FILENAME") }

type ZipReaderLike struct {
	io.ReaderAt
	Size int64
}

func (l ZipReaderLike) ToReader() (ZipReader, error) {
	rdr, e := zip.NewReader(l.ReaderAt, l.Size)
	return ZipReader{rdr}, e
}

func (l ZipReaderLike) PrintItemNames() error {
	zr, e := l.ToReader()
	if nil != e {
		return e
	}
	return zr.PrintNames()
}

func printItemNamesOfZipFile(zipfile *os.File) error {
	stat, e := zipfile.Stat()
	if nil != e {
		return e
	}
	return ZipReaderLike{
		ReaderAt: zipfile,
		Size:     stat.Size(),
	}.PrintItemNames()
}

func zipfilename2itemnames2stdout(zipfilename string) error {
	f, e := os.Open(zipfilename)
	if nil != e {
		return e
	}
	defer f.Close()
	return printItemNamesOfZipFile(f)
}

func main() {
	e := zipfilename2itemnames2stdout(zipFilename())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
