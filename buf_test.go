package buf

import (
    "testing"
)

func TestJoin(t *testing.T) {
    a := New("a")
    b := New("b")

    Join(a, b)

    if Next(a) != b {
        t.Error("Next broken")
    }

    if Prev(b) != a {
        t.Error("Prev broken")
    }

    if Next(b) != nil {
        t.Error("Next broken")
    }

    if Prev(a) != nil {
        t.Error("Prev broken")
    }

    c := New("c")
    d := New("d")
    e := New("e")
    f := New("f")

    Join(b, c, d, e, f)

    if Next(a) != b || Next(b) != c || Next(c) != d || Next(d) != e || Next(e) != f || Next(f) != nil {
        t.Error("Join broken")
    }

    if Prev(f) != e || Prev(e) != d || Prev(d) != c || Prev(c) != b || Prev(b) != a || Prev(a) != nil {
        t.Error("Join broken")
    }

}
