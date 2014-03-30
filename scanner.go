package buf

import "bufio"

type BufScanner struct {
    start  *Buf // first element of buf
    cursor *Buf // cursor while we are scanning
    err    error
}

func NewScanner(b *Buf) *BufScanner {
    return &BufScanner{start: b}
}

// Scanner interface for a Buf. TODO split func not implemented.
// It would add a lot of code if we wanted to split words or \n
// within a Buf
func (b *BufScanner) Scan() bool {

    if b.cursor != nil && b.cursor.bdown != nil {
        b.cursor = b.cursor.bdown
        return true
    } else if b.cursor == nil {
        b.cursor = b.start
        return true
    }

    // we arent scanning anymore
    return false
}

// Current scan data as bytes
func (b *BufScanner) Bytes() []byte {
    return []byte(b.cursor.data)
}

// Current scan data string
func (b *BufScanner) Text() string {
    return b.cursor.data
}

// Scan err
func (b *BufScanner) Err() error {
    return b.err
}

// TODO not implemented
func (b *BufScanner) Split(sfn bufio.SplitFunc) {
    // we dont split, only scan the buffer
}
