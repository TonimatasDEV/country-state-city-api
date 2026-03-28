package util

import (
	"compress/gzip"
	"os"
)

func fileClose(file *os.File) {
	_ = file.Close()
}

func gzipReaderClose(gzReader *gzip.Reader) {
	_ = gzReader.Close()
}
