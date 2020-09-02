package interfaces

import (
	"io"
	"os"
	"time"
)

func reader(r *io.PipeReader) {
	io.Copy(os.Stdout, r)
}

func write(w *io.PipeWriter) {
	for {
		<-time.Tick(time.Second * 1)
		w.Write([]byte("test\n"))
	}
}

func main() {
	r, w := io.Pipe()
	go reader(r)
	write(w)

}
