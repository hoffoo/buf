package buf

func (b *Line) Debug() (s string) {

    scan := b.NewScanner()
    for scan.Scan() {
        s += "|" + scan.Text() + "|\n"
    }

    return
}
