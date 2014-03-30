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

    Join(a, b, c, d ,e ,f, g)

    expect := []byte{'a'}
    for a.Scan() == true {
        if a.Text() != string(expect) {
            t.Errorf("expected string %s: got %s", string(expect), a.Text())
        }
        if bytes.Equal(a.Bytes(), expect) == false {
            t.Errorf("expected []byte %s: got %s", string(expect), a.Text())
        }
        expect[0]++
    }
}
