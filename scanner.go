package buf

import "bufio"

// Scanner for Lines
type LineScanner struct {
    start  *Line // first element of buf
    cursor *Line // cursor while we are scanning
    err    error
}

// Scanner interface for a Line. TODO split func not implemented.
// It would add a lot of code if we wanted to split words or \n
// within a Line
func (l *Line) NewScanner() *LineScanner {
    return &LineScanner{start: l}
}

// Advance scanner
func (ls *LineScanner) Scan() bool {

    if ls.cursor != nil && ls.cursor.next != nil {
        ls.cursor = ls.cursor.next
        return true
    } else if ls.cursor == nil {
        ls.cursor = ls.start
        return true
    }

    // we arent scanning anymore
    return false
}

// Current scan data as bytes
func (ls *LineScanner) Bytes() []byte {
    return []byte(ls.cursor.data)
}

// Current scan data string
func (ls *LineScanner) Text() string {
    return ls.cursor.data
}

// Scan err
func (ls *LineScanner) Err() error {
    return ls.err
}

// TODO not implemented
func (ls *LineScanner) Split(sfn bufio.SplitFunc) {
    // we dont split, only scan the buffer
}
