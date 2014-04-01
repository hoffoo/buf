package buf

func (b *Line) debug() (s string) {

    scan := b.NewScanner()
    for scan.Scan() {
        s += "|" + scan.Text() + "|\n"
    }

    return
}
