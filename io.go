package buf

// IO utils to create bufs

import (
    "io"
    "bufio"
)

func ReadIO(in io.Reader) (b *Buf, err error) {
    s := bufio.NewScanner(in)

    for s.Scan() == true {

        if b == nil {
            b = &Buf{data: s.Text()}
            continue
        }


        if s.Err() != nil {
            return nil, s.Err()
        }

        // TODO more efficient rather than append
        b.Append(&Buf{data: s.Text()})
    }

    return
}
