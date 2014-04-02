package buf

import (
    "testing"
    "strings"
    "bytes"
)

func TestNew(t *testing.T) {
    data := strings.Split("a b c d e f g", " ")

    b := New(data...)

    if b.Len() != 7 {
        t.Error("New broken")
    }

    // test that the right data is in there with a scanner
    expect := []byte{'a'}
    scan := b.NewScanner()
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

func TestJoin(t *testing.T) {
    a := New("a")
    b := New("b")

    a.Join(b)

    if a.Next() != b {
        t.Errorf("Next broken")
    }

    if b.Prev() != a {
        t.Error("Prev broken")
    }

    if b.Next() != nil {
        t.Error("Next broken")
    }

    if a.Prev() != nil {
        t.Error("Prev broken")
    }

    c := New("c")
    d := New("d")
    e := New("e")
    f := New("f")

    b.Join(c, d, e, f)

    if a.Next() != b || b.Next() != c || c.Next() != d || d.Next() != e || e.Next() != f || f.Next() != nil {
        t.Error("Join broken")
    }

    if f.Prev() != e || e.Prev() != d || d.Prev() != c || c.Prev() != b || b.Prev() != a || a.Prev() != nil {
        t.Error("Join broken")
    }

    if a.Len() != 6 || f.Len() != 1 {
        t.Error("Len broken")
    }

    g := New("g")

    f.SetNext(g)
    if g.Next() != nil || f.Next() != g || g.Prev() != f {
        t.Error("SetNext broken")
    }

    h := New("h")

    h.SetPrev(g)
    if h.Next() != nil || h.Prev() != g || g.Next() != h {
        t.Error("SetPrev broken")
    }
}

func TestDelete(t *testing.T) {
    b1 := JoinString(strings.Split("a b c d e f", " ")...)

    if b1.Len() != 6 {
        t.Error("JoinString broken")
    }
}

func TestGet(t *testing.T) {
    b1 := JoinString(strings.Split("a b c d e f", " ")...)

    if b1.Get(2).String() != "c" {
        t.Error("Get failed")
    }

    if b1.Get(100) != nil {
        t.Error("Get failed")
    }
}

func TestFirstLast(t *testing.T) {
    b1 := JoinString(strings.Split("a b c d e f", " ")...)

    if b1.First().String() != "a" {
        t.Error("First failed")
    }

    if b1.Last().String() != "f" {
        t.Error("Last failed")
    }
}
