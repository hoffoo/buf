package buf

func (b *Buf) Debug() (s string) {

    scan := NewScanner(b)
    for scan.Scan() {
        s += "|" + scan.Text() + "|\n"
    }

    return
}
