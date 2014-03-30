package buf

import (
    "testing"
)

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

    if a.Bottom() != f || b.Bottom() != f || f.Bottom() != f {
        t.Error("Bottom broken")
    }

    if f.Top() != a || c.Top() != a || a.Top() !=a {
        t.Error("Top broken")
    }

    if a.Len() != 6 || f.Len() != 1 {
        t.Error("Len broken")
    }

    g := New("g")
    a.Append(g)

    if a.Bottom() != g || b.Bottom() != g || g.Bottom() != g {
        t.Error("Bottom broken")
    }
}
