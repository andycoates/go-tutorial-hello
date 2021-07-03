package main

import (
    "fmt"
)

// https://github.com/golang/go/wiki/Iota
type ByteSize float64
const (
    _           = iota             // iota: 0, ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota) // iota: 1, so value is 1 << (10 * 1) = 2^10 = 1024
    MB                             // iota: 2, sp value is 1 << (10 * 2) = 2^20 = 1024 * 1024 = 1048576
    GB                             // etc...
    TB
    PB
    EB
    ZB
    YB
)

type Document struct {
    name string
    size ByteSize // reuse this from package-level definition
}

func main() {
    fmt.Printf("const: KB is %v bytes\n", KB)

    var a string = "abcdefghijklm"
    fmt.Printf("string a: %v sliced as [2:5] is %v\n", a, a[2:5])

    var b = [...]int{1, 2, 3, 4} // array, same as [4]int
    fmt.Printf("array b: %T %v (%d)\n", b, b, len(b))
    // b = append(b, 5) // can't do this - array is fixed length, append only works with slices

    var c = []int{1, 2, 3, 4}                 // slice
    fmt.Printf("slice c: %T %v (%d)\n", c, c, len(c))
    c = append(c, 5)
    fmt.Printf("slice c: %T %v (%d)\n", c, c, len(c))

    c = append(c, 6, 7, 8, 9)                 // append supports any number of parameters
    c = append(c, []int{10, 11, 12, 13}...)   // the ... expands the array out to individual parameters!
    fmt.Printf("slice c: %T %v (%d)\n", c, c, len(c))

    // defining + initialising a map
    d := map[string]int{
        "foo": 1,
    }
    d["bar"] = 2
    fmt.Printf("map d: %T %v\n", d, d)

    // or ..
    var e map[string]int     // define it
    e = make(map[string]int) // initialise it (if you didn't, runtime exception occurs)
    e["foo"] = 1             // assign something
    fmt.Printf("map e: %T %v\n", e, e)

    // f := d["foo"]    // gets value of key "foo", or 0 if the key doesn't exist, but you can't tell if it was actually 0 or if the key doesn't exist..
    f, ok := d["foo"]   // gets value of key "foo", or 0 if the key doesn't exist, but also sets ok as true/false if key exists
    fmt.Printf("map d value foo: %T %v %v\n", f, f, ok)

    g, ok := d["foooooooo"]
    fmt.Printf("map d value foooooooo: %T %v %v\n", g, g, ok)

    h := d
    delete(h, "foo")
    fmt.Printf("map h: %T %v\n", h, h) // h has "foo" removed
    fmt.Printf("map d: %T %v\n", d, d) // so has d, as its by reference

    myDocument := Document{ // from package-level definition
        name: "expenses.xls",
        size: 123456,
    }
    fmt.Printf("myDocument: %T %v\n", myDocument, myDocument)
    fmt.Printf("myDocument.name: %T %v\n", myDocument.name, myDocument.name)

    anonDocument1 := struct{name string; size int}{name: "filename.ext", size: 123456}
    fmt.Printf("anotherDocument: %T %v\n", anonDocument1, anonDocument1)

    var anonDocument2 struct{name string; size int}
    anonDocument2.name = "filename.ext"
    fmt.Printf("anotherDocument: %T %v\n", anonDocument2, anonDocument2)
}