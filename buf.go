package buf

// Data structure for manupulating line sperated strings.
// useful for managing line sperated data
//
// not thread safe
type Buf struct {
    bup   *Buf
    bdown *Buf
    data  string
}

// Make a new buffer containing the data
func New(data string) *Buf {
    return &Buf{
        data: data,
    }
}

// Return buf data
func (b *Buf) String() string {
    return b.data
}

// Next item in the buf
func (b *Buf) Next() *Buf {
    return b.bdown
}

// Prev item in the buf
func (b *Buf) Prev() *Buf {
    return b.bup
}

// Count of buffers starting at b to Bottom()
func (b *Buf) Len() (l int) {
    for b != nil {
        l++
        b = b.Next()
    }

    return
}

// Find the bottom of this buffer
func (b *Buf) Bottom() (*Buf) {
    if bb := b.Next(); bb != nil {
        return bb.Bottom()
    }

    return b
}

// Find the start of this buffer
func (b *Buf) Top() (*Buf) {
    if tb := b.Prev(); tb != nil {
        return tb.Top()
    }

    return b
}

// Delete an item, links the items below and above
func (b *Buf) Delete() {
    if b.bup != nil {
        b.bup = b.bdown
    }
    if b.bdown != nil {
        b.bdown.bup = b.bup
    }
}

// Join bufs
func (u *Buf) Join(more ...*Buf) {

    if (len(more) == 0) {
        return
    }

    d := more[0]

    if u.bdown != nil {
        d.bdown = u.bdown.bup
        u.bdown.bup = d
    } else {
        d.bup = u
    }
    u.bdown = d

    d.Join(more[1:]...)
}


// Appends to the end of the buffer
func (b *Buf) Append(more ...*Buf) {
    b.Bottom().Join(more...)
}
