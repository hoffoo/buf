package buf

// Data structure for manupulating line sperated strings.
// useful for managing line sperated data
//
// not thread safe
type Line struct {
    prev   *Line
    next *Line
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
    l.next = to
    to.prev = l
}

// Set Prev to *Line
func (l *Line) SetPrev(to *Line) {
    l.prev = to
    to.next = l
}

// Next item in the buf
func (b *Line) Next() *Line {
    return b.next
}

// Prev item in the buf
func (b *Line) Prev() *Line {
    return b.prev
}

// Count of buffers starting at b to Bottom()
func (b *Line) Len() (l int) {
    for b != nil {
        l++
        b = b.Next()
    }

    return
}

// Delete an item, links the items below and above
func (b *Line) Delete() {
    if b.prev != nil {
        b.prev = b.next
    }
    if b.next != nil {
        b.next.prev = b.prev
    }
}

// Join bufs
func (u *Line) Join(more ...*Line) {

    if (len(more) == 0) || more[0] == nil {
        return
    }

    d := more[0]

    if u.next != nil {
        d.next = u.next.prev
        u.next.prev = d
    } else {
        d.prev = u
    }
    u.next = d

    d.Join(more[1:]...)
}

// Make a marker, empty data buf, to keep track this buf.
// The marker buf element is special because the element below it
// is not linked. This means that you can call Next() on a marker
// to access the data below it, but cannot call Prev() to get to it.
// It will of unlink a buf from anything above it since iteration
// relies on the Prev() method
func (b *Line) NewMarker() *Line {
    return &Line{next: b}
}

