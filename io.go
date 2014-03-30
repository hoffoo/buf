package buf

// IO utils to create lines

import (
    "io"
    "bufio"
)

// Read from in, first line is returned, the rest are joined under it.
// returns error in case of scanner/io err.
func ReadIO(in io.Reader) (l *Line, err error) {

    s := bufio.NewScanner(in)

    // scan once to get the first line
    if s.Scan() == false || s.Err() != nil {
        return nil, s.Err()
    }

    l = &Line{data: s.Text()}
    scanIO(l, s)

    return l, s.Err()
}

func scanIO(l *Line, s *bufio.Scanner) {

    if s.Scan() == false || s.Err() != nil {
        return
    }

    lnext := &Line{data: s.Text()}
    l.Join(lnext)
    scanIO(lnext, s)
}
