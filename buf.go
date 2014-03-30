package buf

// data structure for manupulating chunks of strings
// useful for managing line sperated strings
//
// not thread safe

type Buf struct {
    bup   *Buf
    bdown *Buf
    data  string
    scan  *BufScan
    read  *BufRead
}

func New(data string) *Buf {
    return &Buf{
        data: data,
    }
}

func Next(b *Buf) *Buf {
    return b.bdown
}

func Prev(b *Buf) *Buf {
    return b.bup
}

func Delete(b *Buf) {
    if b.bup != nil {
        b.bup = b.bdown
    }
    if b.bdown != nil {
        b.bdown.bup = b.bup
    }
}

func Join(u *Buf, more ...*Buf) {

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

    Join(d, more[1:]...)
}
