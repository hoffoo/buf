package buf

// Data structure for manupulating line sperated strings.
// useful for managing line sperated data.
//
// not thread safe
type Line struct {
    prev *Line
    next *Line
    data string
}

// Make a new buffer containing the data
func New(data string) *Line {
    return &Line{
        data: data,
    }
}

// Return buf data
func (l *Line) String() string {
    return l.data
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
func (l *Line) Next() *Line {
    return l.next
}

// Prev item in the buf
func (l *Line) Prev() *Line {
    return l.prev
}

// Count of buffers starting at b to Bottom()
func (l *Line) Len() (li int) {
    for l != nil {
        li++
        l = l.Next()
    }

    return
}

// Delete an item, links the items below and above
func (l *Line) Delete() {
    if l.prev != nil {
        l.prev.next = l.next
    }
    if l.next != nil {
        l.next.prev = l.prev
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

// external Join func for convenience
func Join(more ...*Line) *Line {

    if len(more) < 2 {
        return nil
    }

    more[0].Join(more[1:]...)
    return more[0]
}

func (u *Line) JoinString(ss ...string) {

    if len(ss) == 0 {
        return
    }

    ls := make([]*Line, len(ss))
    for i := range ss {
        ls[i] = New(ss[i])
    }
    u.Join(ls...)
}

// external Join func for strings
func JoinString(ss ...string) *Line {

    if len(ss) == 0 {
        return nil
    }

    l := New(ss[0])
    l.JoinString(ss[1:]...)
    return l
}

// Make a marker, empty data buf, to keep track this buf.
// The marker buf element is special because the element below it
// is not linked. This means that you can call Next() on a marker
// to access the data below it, but cannot call Prev() to get to it.
func (l *Line) NewMarker() *Line {
    return &Line{next: l}
}
