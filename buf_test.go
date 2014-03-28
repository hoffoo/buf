package buf

import (
    "testing"
)

func TestMoving(t *testing.T) {
    a := New("a")
    b := New("b")
    //c := New("c")

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

}
