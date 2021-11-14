package count

import "io"

type CountingWriter struct {
	writer io.Writer
	count int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {

}
