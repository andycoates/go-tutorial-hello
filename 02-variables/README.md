# 02-variables

Based on https://golang.org/doc/tutorial/getting-started#call and following on from previous tutorial..

## What did we learn?

- `import` can use a block syntax `( .... )` to import multiple modules instead of typing `import` on each line.
- this also applies to `var` definitions! (and probably other definitions I'm guessing)
- Go supports package level variable scopes, and local block scopes - and defining the same variable in a local scope is called `shadowing` (it "shadows" the original)
- lowerCase variables defined at the package level are scoped to the package, whereas UpperCase variables are exported/globally visable outside the package
- casting a variable to another type is pretty much same as most languages, i.e. `a = type(b)` e.g. `a := int(b)`
- `fmt.Printf()` is pretty funky, it has `%v` (value) and `%T` (type) if you wanted to debug easily
- Go is pretty strict on using variables!  If you don't use a variable it will complain at compile time.

Output from this code:
```
Don't communicate by sharing memory, share memory by communicating.
(float32) 1
(int) 88
(float32) 88.000000
(string) this is a string
```