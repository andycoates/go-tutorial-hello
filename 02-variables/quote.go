package main

import (
    "fmt"
    "rsc.io/quote"
)

// note: lowercase variables are local to the package and not are not exported
var (
    packageScoped int = 99
)

func main() {
    // can declare a variable fully..
    var singleVariable float32 = 1

    // or in a block
    var (
        packageScoped int = 88    // "shadowing" (overriding a higher scope)
    )

    // or let Go work it out with a short form declaration..
    myVariable := float32(packageScoped)  // cast the int to a float32 as well
    aString := "this is a string"

    fmt.Println(quote.Go())
    fmt.Printf("(%T) %v\n", singleVariable, singleVariable)
    fmt.Printf("(%T) %v\n", packageScoped, packageScoped)
    fmt.Printf("(%T) %f\n", myVariable, myVariable)  // being specific with the format
    fmt.Printf("(%T) %v\n", aString, aString)
}