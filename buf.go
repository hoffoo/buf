// Data structure for manupulating ordered strings.
// Useful for managing line sperated data.
//
// not thread safe
package buf

// Single item of in the list.
type Line struct {
    prev *Line
    next *Line
    data string
}

// Make a new line containing the data
func New(data string) *Line {
    return &Line{
        data: data,
    }
}

// Return line data
func (l *Line) String() string {
    return l.data
}

// Next line
func (l *Line) Next() *Line {
    return l.next
}

// Prev line
func (l *Line) Prev() *Line {
    return l.prev
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

// Count of lines starting at b to Bottom()
func (l *Line) Len() (li int) {
    for l != nil {
        li++
        l = l.Next()
    }

    return
}

// Delete a line, links the Next() and Prev() of l to remove this
func (l *Line) Delete() {
    if l.prev != nil {
        l.prev.next = l.next
    }
    if l.next != nil {
        l.next.prev = l.prev
    }
}

// Join lines in the given order
func (p *Line) Join(more ...*Line) {

    if (len(more) == 0) || more[0] == nil {
        return
    }

    n := more[0]

    if p.next != nil {
        n.next = p.next.prev
        p.next.prev = n
    } else {
        n.prev = p
    }
    p.next = n

    n.Join(more[1:]...)
}

// Join func for convenience - see *Line.Join
func Join(more ...*Line) *Line {

    if len(more) < 2 {
        return nil
    }

    more[0].Join(more[1:]...)
    return more[0]
}

// Join strings to line - create lines for each ss and join
func (p *Line) JoinString(ss ...string) {

    if len(ss) == 0 {
        return
    }

    ls := make([]*Line, len(ss))
    for i := range ss {
        ls[i] = New(ss[i])
    }
    p.Join(ls...)
}

// Join strings to line - create line for each in ss, join, return
// head line
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
