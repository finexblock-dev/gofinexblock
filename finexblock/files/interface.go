package files

type Writer interface {
	Logger
	Write(p []byte) (n int, err error)
}

type Reader interface {
}

type Logger interface {
	Println(v ...any)
}

func NewWriter(prefix, filename string) Writer {
	return newFileWriter(prefix, filename)
}
