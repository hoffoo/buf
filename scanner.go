package buf

import "bufio"

type BufScan struct {
    cursor *Buf
    err    error
}

func (b *Buf) Scan() bool {

    if b.scan == nil {
        b.scan = &BufScan{cursor: b}
        return true
    }

    if b.scan.cursor.bdown != nil {
        b.scan.cursor = b.scan.cursor.bdown
        return true
    }

    // we arent scanning anymore
    b.scan = nil
    return false
}

func (b *Buf) Bytes() []byte {

    if b.scan == nil {
        return []byte{}
    }

    return []byte(b.scan.cursor.data)
}

func (b *Buf) Text() string {
    if b.scan == nil {
        return ""
    }
    return b.scan.cursor.data
}

func (b *Buf) Err() error {
    if b.scan == nil {
        return nil
    }
    return b.scan.err
}

func (b *Buf) Split(sfn bufio.SplitFunc) {
    // we dont split, only scan the buffer
}
