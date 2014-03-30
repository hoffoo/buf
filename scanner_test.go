package buf

import (
    "testing"
    "bytes"
)

func TestScan(t *testing.T) {

    a := New("a")
    b := New("b")
    c := New("c")
    d := New("d")
    e := New("e")
    f := New("f")
    g := New("g")

    a.Join(b, c, d ,e ,f, g)

    expect := []byte{'a'}

    scan := NewScanner(a)
    for scan.Scan() == true {
        if scan.Text() != string(expect) {
            t.Errorf("expected string %s: got %s", string(expect), scan.Text())
        }
        if bytes.Equal(scan.Bytes(), expect) == false {
            t.Errorf("expected []byte %s: got %s", string(expect), scan.Text())
        }
        expect[0]++
    }
}
