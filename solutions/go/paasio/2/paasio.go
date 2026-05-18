package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type writeCounter struct {
	w     io.Writer
	bytes int64
	calls int
	mu    sync.Mutex
}

type readCounter struct {
	r     io.Reader
	bytes int64
	calls int
	mu    sync.Mutex
}

type readWriteCounter struct {
	writeCounter
	readCounter
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{w: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{r: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		writeCounter: writeCounter{w: readwriter},
		readCounter:  readCounter{r: readwriter},
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	n, err := rc.r.Read(p)
	if n != 0 {
		rc.bytes += int64(n)
		rc.calls++
	}
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	return rc.bytes, rc.calls
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	n, err := wc.w.Write(p)
	if n != 0 {
		wc.bytes += int64(n)
		wc.calls++
	}
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	return wc.bytes, wc.calls
}
