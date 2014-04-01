buf
===

go linked list of strings 

[![GoDoc](https://godoc.org/github.com/hoffoo/buf?status.png)](https://godoc.org/github.com/hoffoo/buf)

usage
===

```go

a := buf.New("a")
b := buf.New("b")

a.Join(b)

c := buf.New("c")
d := buf.New("d")
e := buf.New("e")
f := buf.New("f")

b.Join(c, d, e, f)

if a.Next() != b || b.Next() != c || c.Next() != d || d.Next() != e || e.Next() != f || f.Next() != nil {
    t.Error("Join broken")
}

if f.Prev() != e || e.Prev() != d || d.Prev() != c || c.Prev() != b || b.Prev() != a || a.Prev() != nil {
    t.Error("Join broken")
}

```
