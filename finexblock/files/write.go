package files

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

type fileWriter struct {
	logger *log.Logger
}

func (f *fileWriter) write(p []byte) (n int, err error) {
	return f.logger.Writer().Write(p)
}

func initLogger(f *fileWriter, prefix, filename string) {
	f.logger = log.New(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    1,
		MaxAge:     28,
		MaxBackups: 3,
		Compress:   true,
	}, prefix, 0)
}

func newFileWriter(prefix, filename string) *fileWriter {
	f := &fileWriter{}
	initLogger(f, prefix, filename)

	return f
}
