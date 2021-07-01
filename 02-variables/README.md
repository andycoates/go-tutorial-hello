# 02-variables

Based on https://golang.org/doc/tutorial/getting-started#call and following on from previous tutorial..

## What did we learn?

- `import` can use a block syntax `( .... )` to import multiple modules instead of typing `import` on each line.
- ... this also applies to `var` definitions! (and probably other definitions I'm guessing)
- Go supports package level variable scopes, and local block scopes - and defining the same variable in a local scope is called `shadowing` (it "shadows" the original)
- lowerCase variables defined at the package level are scoped to the package, whereas UpperCase variables are exported/globally visable outside the package
- Casting a variable to another type is pretty much same as most languages, i.e. `a = type(b)` e.g. `a := int(b)`
- `fmt.Printf()` is pretty funky, it has `%v` (value) and `%T` (type) if you wanted to debug easily
- Go is pretty strict on using variables!  If you don't use a variable it will complain at compile time.
- Uninitialised variables have a "zero" value, i.e. `0` for numeric types, `""` (empty string) for strings, and `false` for booleans
- string creation with double-quotes `""` are internally stored as bytes (uint8) and encoded as UTF-8, so 1 character can be 1-3 bytes
- single-quotes `''` mean something different - you are indicating its a `rune`, a single `int32` value, e.g. `'明'`
- ... so `"明"` would have a `len()` of 3 as its actually stored as 3x bytes (`uint8`), whereas `'明'` would have a `len()` of 1 as its stored as 1x `int32`

Output from this code:
```
Don't communicate by sharing memory, share memory by communicating.
(float32) 1.1
(int) 88
(float32) 88.000000
(string) this is a string
defaults: bool=false, int=0, string=
 string size: hello 明 (string of length 9)
as byte size: [104 101 108 108 111 32 230 152 142] ([]uint8 of length 9)
as rune size: [104 101 108 108 111 32 26126] ([]int32 of length 7)
```