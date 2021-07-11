# 04-pointers

Pointers (now I'm really having to think back to my CompSci studies...)

## What did we learn?

- https://gobyexample.com/pointers
- most variable type assignments are copied by value (`b := a`, `b` is a copy of `a` and independant of changes to `a`), except as in previous tutorial, things like `maps` and `slices` are by reference.
- think of `&` as "address of", so if we did `b := &a`, `b` is now a pointer to the address of `a` - if you printed the value of `b` you'd get the memory address
- to access the actual value a pointer is pointing to, dereference it with `*`, i.e. `fmt.Println(*b)`
- side note to avoid confusion: the `*` before a type is just declaring that variable as a pointer, it does not dereference when used in variable type definitions, e.g. from above, `b := &a` is basically `var b *int = &a` (`b` is a pointer to an `int` which is the address of `int` `a`), it does not dereference the value at this point.
- you can also assign new values by dereferencing (`*b = new_value`), in which case any variable pointing to or has the address of that memory address (e.g. `a`) also gets changed.
- if a pointer isn't initialised (`var foo *int`) it will be `nil` (pretty much always used when checking if a pointer is initialised) - if you tried to access it, it'd cause a runtime error
- when using a pointer to a `struct`, you could dereference the struct (e.g. `*doc1`), but to access a field of the struct you need to encapsulate the struct with parentheses `()` since `*` has a lower precedence than `.` and without `()` it'd try and access the field of a pointer etc, so `(*doc1).field` first dereferences, then accesses the field.
- there's also a shortcut to accessing fields - you can just do `doc1.field` and Go is smart enough to know you actually want the field of the underlying struct
- speaking of structs, you can initialise a pointer to a new empty struct with `new()`, e.g. `doc1 := new(Document)`
- if you want to create a pointer and initialise with values, you can just use the address of `&` syntax, e.g. `doc2 := &Document{name: "foo"}`