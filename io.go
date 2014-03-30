package buf

// IO utils to create bufs

import (
    "io"
    "bufio"
)

func ReadIO(in io.Reader) (l *Line, err error) {
    s := bufio.NewScanner(in)

    for s.Scan() == true {

        if l == nil {
            l = &Line{data: s.Text()}
            continue
        }


        if s.Err() != nil {
            return nil, s.Err()
        }

        // TODO more efficient rather than append
        l.Append(&Line{data: s.Text()})
    }

    return
}
