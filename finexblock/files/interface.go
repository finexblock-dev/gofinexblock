package files

type Writer interface {
	write(p []byte) (n int, err error)
}

type Reader interface {
}

func NewWriter(prefix, filename string) Writer {
	return newFileWriter(prefix, filename)
}
