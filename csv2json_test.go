package csv2json

import (
    "bytes"
    "testing"
)

var columns = []string{"name", "email", "phone"}
var want = `[
  {
    "email": "ken@google.com",
    "name": "Ken Thompson",
    "phone": "555-5550"
  },
  {
    "email": "rob@google.com",
    "name": "Rob Pike",
    "phone": "555-5551"
  },
  {
    "email": "robert@google.com",
    "name": "Robert Griesemer",
    "phone": "555-5552"
  }
]`

var input = `
Ken Thompson, ken@google.com, 555-5550
Rob Pike, rob@google.com, 555-5551
Robert Griesemer, robert@google.com, 555-5552
`

func TestConvertWithBuffer(t *testing.T) {
    var b bytes.Buffer
    _, err := b.WriteString(input)
    if err != nil {
        t.Error(err)
    }
    got, err := Convert(bytes.NewReader(b.Bytes()), columns)
    if err != nil {
        t.Error(err)
    }
    if string(got) != want {
        t.Errorf("TestConvertWithBuffer(f, %s) = %s; want %s",
            columns, string(got), want)
    }
}
