package main

import (
    "fmt"
)

func main() {
    a := 1
    b := a
    fmt.Printf("%v %v (%T)\n", a, b, b)  // b gets a copy of a's value
    a = 2
    fmt.Printf("%v %v (%T)\n", a, b, b)  // b is unchanged

    c := &a                              // now we use c to point to address of a
    fmt.Printf("%v %v (%T)\n", a, c, c)  // c's value is the address of a, not the value - and its type is now a pointer (to an int)
    fmt.Printf("%v %v (%T)\n", &a, c, c) // we can compare &a (address of a) to c (address of a as well) to show they are the same address
    fmt.Printf("%v %v (%T)\n", a, *c, c) // we can also access the underlying value of the pointer by dereferencing it (which points to a, so is 2)

    *c = 3                               // we can also change the underlying value itself when dereferencing it
    fmt.Printf("%v %v (%T)\n", a, *c, c) // now a also changed because we dereferenced c (the address of a) and assigned a value

    var d *int
    fmt.Printf("%v (%T)\n", d, d)        // if a pointer isn't initialised, it's value is nil ()
    // fmt.Printf("%v (%T)\n", *d, d)    // would crash because we're trying to dereference a nil pointer (it's not pointing to anything)
    d = &a
    fmt.Printf("%v (%T)\n", d, d)        // its a pointer to the address of..
    fmt.Printf("%v (%T)\n", *d, d)       // and again we can dereference to get the value

    // basic struct like from last tutorial
    type Document struct {
        name string
    }

    doc1 := new(Document)                // we can use new() to create a pointer to a new Document struct
    fmt.Printf("%v (%T)\n", *doc1, doc1) // its empty (string defaul value is "")
    (*doc1).name = "filename1.txt"       // deference to access the value, and assign the name field a value
    fmt.Printf("%v (%T)\n", *doc1, doc1) // its empty (string defaul value is "")
    doc1.name = "filename2.txt"          // shortcut... go knows you want to access the value of the struct field
    fmt.Printf("%v (%T)\n", *doc1, doc1)
    fmt.Printf("%v (%T)\n", (*doc1).name, doc1) // same thing - accessing the value of doc1, specifically its name field
    fmt.Printf("%v (%T)\n", doc1.name, doc1)    // same thing - go knows you just want to access the value

    doc2 := &Document{name: "filename2.txt"}    // if we wanted to declare+point in one statement
    fmt.Printf("%v (%T)\n", doc2.name, doc2)
}