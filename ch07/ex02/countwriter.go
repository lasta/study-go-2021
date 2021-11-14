package countwriter

import "io"

type CountWriter struct {
	writer io.Writer
	count  int64
}

func (w *CountWriter) Write(p []byte) (count int, err error) {
	count, err = w.writer.Write(p)
	w.count += int64(count)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	countingWriter := &CountWriter{
		writer: w,
		count:  0,
	}
	return countingWriter, &countingWriter.count
}
