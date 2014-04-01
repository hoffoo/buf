package buf

import (
    "testing"
    "strings"
)

func TestReadIO(t *testing.T) {
    data := `buf1
buf2
buf3`

    rd := strings.NewReader(data)
    b, err := ReadIO(rd)

    if err != nil {
        t.Error(err)
    }

    if b.String() != "buf1" || b.Next().String() != "buf2" || b.Next().Next().String() != "buf3" || b.Next().Next().Next() != nil {
        t.Errorf("ReadIO broken, output was: %s", b.debug())
    }
}
