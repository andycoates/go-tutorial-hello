# 03-structures

Touching on various structures and concepts

## What did we learn?

- https://gobyexample.com/constants
- https://github.com/golang/go/wiki/Iota
- `iota` is a special identifier used with constant block definitions that will start at 0 and increment for every entry, i.e. a = iota (a would equal 0), b = iota (b would equal 1)
- if you omit `iota` and the definition of further entries, it will magically repeat the same action from the last use of it
- underscore `_` (called the "Blank Identifier") can be used as a variable to essentially tell Go you don't care about what is being assigned (a function may return something, but you just don't care about it...)
- https://gobyexample.com/arrays / https://gobyexample.com/slices
- Go has both fixed length arrays (`[3]int`), or any length "slices" (`[]int`), with fixed being more efficient as the memory allocation is fixed ahead of time etc
- the range slicing syntax is the same as python, i.e. `[x:]`, `[:y]`, and `[x:y]` (from index `x` up to, but not including, index `y`)
- you can initialise values of arrays/slices with the syntax `[4]int{1, 2, 3, 4}` etc
- ... but Go is smart enough to know that initialising an `array` with 4 values will make its size 4, so you can just do `[...]int{1, 2, 3, 4}` instead
- Speaking of `...`, when used like `[]int{1, 2, 3, 4}...` Go will expand/stretch out the slice like they were individual parameters, e.g. `append(foo, 1, 2, 3, 4)` is the same as `append(foo, []int{1, 2, 3, 4}...)`
- `maps` are very similar to other languages (e.g. Groovy), and basically a flexible `[k]=v` structure (like `dicts` in python) - define them as `map[<key type>]<value type>`, e.g. `map[string]int` (key is a `string`, value is an `int`) and initialise them just like other variables with `{  }`
- https://blog.golang.org/maps / https://gobyexample.com/maps
```
var m map[string]int

Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn't point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don't do that. To initialize a map, use the built in make function:

m = make(map[string]int)
```
- use `delete(m, key)` to delete a key
- both `slices` and `maps` are accessed by reference (a pointer) - if you point a variable at one and change it, the original changes too.
- https://golang.org/ref/spec#Struct_types / https://gobyexample.com/structs
- `structs` are similar to structs in other languages like C - you can define a complex data structure of different field types, and you access the fields with the `.` dot notation, e.g `myStruct.fieldName`
- When you initialise a `struct`, you don't actually have to specify the key when assigning the value - if you omit the key it will treat the assignment per the struct definition ordering, e.g.
```
type Document struct {
    name string
    size int
}
```
then..
```
myDocument := Document{
    name: "filename.ext",
    size: 123456,
}
```
is the same as
```
myDocument := Document{
    "filename.ext",
    123456,
}
```
... but be very wary of using that syntax - if the definition order changed you could get caught out
- you can also define anonymous `structs` with direct initialisation, e.g.
```
myDocument := struct{ name string; size int }{ name: "filename.ext", size: 123456}
```
or define and assign later:
```
var myDocument struct{name string; size int}
myDocument.name = "filename.ext"
```
- `structs` aren't like a `map` or a `slice` where assignment is by reference (a pointer), they are by value (a copy)