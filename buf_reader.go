package buf

//import "io"

func (b *Buf) Read(p []byte) (n int, err error) {

    for i, c := range b.data[b.offset:] {

        // filled the buffer
        if i == len(p) {
            return
        }

        p[i] = byte(c)
        n++
        b.offset++
    }

    // we need to read more to fill the buffer
    if n < len(p) && b.bdown != nil{
        b.offset = 0
        return b.Read(p[n:])
    }

}
