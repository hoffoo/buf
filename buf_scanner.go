package buf

import "bufio"

func (b *Buf) Scan() bool {
	if b.bdown != nil {
		b.cursor = b.bdown
		return true
	}

	return false
}

func (b *Buf) Bytes() []byte {
	return []byte(b.cursor.data)
}

func (b *Buf) Text() string {
	return b.cursor.data
}

func (b *Buf) Split(sfn bufio.SplitFunc) {
	// we dont split, only buf iter
}

func (b *Buf) Err() error {
	return b.serr
}
