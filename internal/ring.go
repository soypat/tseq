package internal

import (
	"bytes"
	"errors"
	"io"
	"math"
	"unsafe"
)

var errRingBufferFull = errors.New("tseq/ring: buffer full")

// NewRing returns a new ring buffer ready for use.
func NewRing(buf []byte) *Ring {
	return &Ring{buf: buf}
}

// Ring implements basic Ring buffer functionality.
type Ring struct {
	buf []byte
	off int
	end int
}

// WriteString is a wrapper around [Ring.Write] that avoids allocation of converting byte slice to string.
func (r *Ring) WriteString(s string) (int, error) {
	return r.Write(unsafe.Slice(unsafe.StringData(s), len(s)))
}

// Write appends data to the ring buffer that can then be read back in order with [Ring.Read] methods. An error is returned if length of data too large for buffer.
func (r *Ring) Write(b []byte) (int, error) {
	free := r.Free()
	if len(b) > free {
		return 0, errRingBufferFull
	}
	midFree := r.midFree()
	if midFree > 0 {
		// start     end       off    len(buf)
		//   |  used  |  mfree  |  used  |
		n := copy(r.buf[r.end:r.off], b)
		r.end += n
		return n, nil
	}
	// start       off       end      len(buf)
	//   |  sfree   |  used   |  efree   |
	n := copy(r.buf[r.end:], b)
	r.end += n
	if n < len(b) {
		n2 := copy(r.buf, b[n:])
		r.end = n2
		n += n2
	}
	return n, nil
}

// ReadDiscard is a performance auxiliary method that performs a dummy read or no-op read
// for advancing the read pointer n bytes without actually copying data.
// This method panics if amount of bytes is more than buffered (see [Ring.Buffered]).
func (r *Ring) ReadDiscard(n int) {
	if n < 0 {
		panic("negative discard amount")
	}
	buffered := r.Buffered()
	switch {
	case n > buffered:
		panic("discard exceeds length")
	case n == buffered:
		r.Reset()
	case n+r.off > len(r.buf):
		r.off = n - (len(r.buf) - r.off)
	default:
		r.off += n
	}
}

// ReadAt reads data at an offset from start of readable data but does not advance read pointer. [io.EOF] returned when no data available.
func (r *Ring) ReadAt(p []byte, off64 int64) (int, error) {
	if math.MaxInt != math.MaxInt64 && off64+int64(len(p)) > math.MaxInt32 {
		return 0, errors.New("offset too large (32 bit overflow)") // Check only compiles for 32-bit platforms.
	}
	off := int(off64)
	if off+len(p) > r.Buffered() {
		return 0, io.ErrUnexpectedEOF
	}
	r2 := *r
	r2.off = (r2.off + off) % (r.Size())
	return r2.ReadPeek(p)
}

// ReadPeek reads up to len(b) bytes from the ring buffer but does not advance the read pointer. [io.EOF] returned when no data available.
func (r *Ring) ReadPeek(b []byte) (int, error) {
	n, _, err := r.read(b)
	return n, err
}

// Read reads up to len(b) bytes from the ring buffer and advances the read pointer. [io.EOF] returned when no data available.
func (r *Ring) Read(b []byte) (int, error) {
	n, newOff, err := r.read(b)
	if err != nil {
		return n, err
	}
	r.off = newOff
	r.onReadEnd()
	return n, nil
}

func (r *Ring) read(b []byte) (n, newOff int, err error) {
	newOff = r.off
	if r.Buffered() == 0 {
		return 0, newOff, io.EOF
	}
	if r.end > r.off {
		// start       off       end      len(buf)
		//   |  sfree   |  used   |  efree   |
		n = copy(b, r.buf[r.off:r.end])
		newOff += n
		return n, newOff, nil
	}
	// start     end       off     len(buf)
	//   |  used  |  mfree  |  used  |
	n = copy(b, r.buf[r.off:])
	newOff += n
	if n < len(b) {
		n2 := copy(b[n:], r.buf[:r.end])
		newOff = n2
		n += n2
	}
	return n, newOff, nil
}

// Reset flushes all data from ring buffer so that no data can be further read.
func (r *Ring) Reset() {
	r.off = 0
	r.end = 0
}

// Size returns the capacity of the ring buffer.
func (r *Ring) Size() int {
	return len(r.buf)
}

// Buffered returns amount of bytes ready to read from ring buffer. Always less than [ring.Size].
func (r *Ring) Buffered() int {
	return r.Size() - r.Free()
}

// Free returns amount of bytes that can be read into ring buffer before reaching maximum capacity given by [ring.Size]. Always less than [ring.Size].
func (r *Ring) Free() int {
	if r.off == 0 {
		return len(r.buf) - r.end
	}

	if r.off < r.end {
		// start       off       end      len(buf)
		//   |  sfree   |  used   |  efree   |
		startFree := r.off
		endFree := len(r.buf) - r.end
		return startFree + endFree
	}
	// start     end       off     len(buf)
	//   |  used  |  mfree  |  used  |
	return r.off - r.end
}

func (r *Ring) midFree() int {
	if r.end >= r.off {
		return 0
	}
	return r.off - r.end
}

// onReadEnd does some cleanup of [ring.off] and [ring.end] fields if possible for contiguous read performance benefits.
func (r *Ring) onReadEnd() {
	if r.end == len(r.buf) {
		r.end = 0 // Wrap around.
	}
	if r.off == len(r.buf) {
		r.off = 0 // Wrap around.
	}
	if r.off == r.end {
		r.Reset() // We read everything, reset.
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (r *Ring) string() string {
	var b bytes.Buffer
	r2 := *r
	b.ReadFrom(&r2)
	return b.String()
}

func (r *Ring) _string(off int64) string {
	s := r.string()
	return s[off:]
}
