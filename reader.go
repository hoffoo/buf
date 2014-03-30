package buf

import "io"

type BufRead struct {
    offset int // offset of the buf data
}

const NEWLINE = byte('\n')

// Reader read method. Fills in with bytes from B until reaching
// the end of the data. If we are at the end of the data we read
// again from the next buffer. By design this means that Read will
// always return data from a single a buffer, no matter if the
// array is full or not.
// Always appends \n when reaching the end of the buf data
func (b *Buf) Read(in []byte) (n int, err error) {

    if b.read != nil {
        b.read = &BufRead{0}
    } else if len(b.data) == b.read.offset {
        // we read to the last buffer
        if b.bdown == nil {
            return 0, io.EOF
        } else {
            // we have more to read
            return b.bdown.Read(in)
        }
    }

    for i, c := range b.data[b.read.offset:] {

        n++             // we read a char, increment n
        b.read.offset++ // we read into the in buffer, offset it

        // filled the buffer
        if b.read.offset == len(in)-1 {
            in[i] = NEWLINE
            b.read.offset = 0
            return
        }

        // reached end of the data
        if b.read.offset == len(b.data) {
            in[i] = NEWLINE
            return
        }

        // put the char into in
        in[i] = byte(c)

    }

    return
}
