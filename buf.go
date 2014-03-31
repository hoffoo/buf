package buf

// Data structure for manupulating line sperated strings.
// useful for managing line sperated data
//
// not thread safe
type Line struct {
    lup   *Line
    ldown *Line
    data  string
}

// Make a new buffer containing the data
func New(data string) *Line {
    return &Line{
        data: data,
    }
}

// Return buf data
func (b *Line) String() string {
    return b.data
}

// Set Next to *Line
func (l *Line) SetNext(to *Line) {
    l.ldown = to
    to.lup = l
}

// Set Prev to *Line
func (l *Line) SetPrev(to *Line) {
    l.lup = to
    to.ldown = l
}

// Next item in the buf
func (b *Line) Next() *Line {
    return b.ldown
}

// Prev item in the buf
func (b *Line) Prev() *Line {
    return b.lup
}

// Count of buffers starting at b to Bottom()
func (b *Line) Len() (l int) {
    for b != nil {
        l++
        b = b.Next()
    }

    return
}

// Find the bottom of this buffer
func (b *Line) Bottom() (*Line) {
    if bb := b.Next(); bb != nil {
        return bb.Bottom()
    }

    return b
}

// Find the start of this buffer
func (b *Line) Top() (*Line) {
    if tb := b.Prev(); tb != nil {
        return tb.Top()
    }

    return b
}

// Delete an item, links the items below and above
func (b *Line) Delete() {
    if b.lup != nil {
        b.lup = b.ldown
    }
    if b.ldown != nil {
        b.ldown.lup = b.lup
    }
}

// Join bufs
func (u *Line) Join(more ...*Line) {

    if (len(more) == 0) || more[0] == nil {
        return
    }

    d := more[0]

    if u.ldown != nil {
        d.ldown = u.ldown.lup
        u.ldown.lup = d
    } else {
        d.lup = u
    }
    u.ldown = d

    d.Join(more[1:]...)
}

// Make a marker, empty data buf, to keep track this buf.
// The marker buf element is special because the element below it
// is not linked. This means that you can call Next() on a marker
// to access the data below it, but cannot call Prev() to get to it.
// It will of unlink a buf from anything above it since iteration
// relies on the Prev() method
func (b *Line) NewMarker() *Line {
    return &Line{ldown: b}
}

