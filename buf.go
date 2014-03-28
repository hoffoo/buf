package buf

// data structure for manupulating chunks of strings
// useful for managing line sperated strings
//
// not thread safe

type Buf struct {
    bup    *Buf
    bdown  *Buf
    cursor *Buf  // used when scanning
    serr   error // scanner error
    data   string
    offset int   // offset of string during last read
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

func Join(u *Buf, d *Buf) {
    if u.bdown != nil {
        d.bup = u.bdown
    }
    u.bdown = d
    d.bup = u
}
