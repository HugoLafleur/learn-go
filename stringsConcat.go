package main

import (
    "bytes"
    "fmt"
)

func main() {
    var buffer bytes.Buffer

    buffer.WriteString("http://duproprio.com/search/?hash=/g-re=6/s-pmin=200000/s-pmax=500000/s-bmin=2/s-build=1/s-parent=1/s-filter=forsale/s-days=0/m-pack=house-condo-multiplex/p-con=main/p-ord=date/p-dir=DESC/pa-ge=")
    buffer.WriteString("1")

    var newstring = buffer.String()
    buffer.Reset()

    fmt.Println(buffer.String())
    fmt.Println(newstring)
}