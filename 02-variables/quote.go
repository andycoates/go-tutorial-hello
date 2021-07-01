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
    var singleVariable float32 = 1.1

    // or in a block
    var (
        packageScoped int = 88    // "shadowing" (overriding a higher scope)
    )

    // or let Go work it out with a short form declaration..
    myVariable := float32(packageScoped)  // cast the int to a float32 as well
    aString := "this is a string"

    // uninitialised variables have a 0/false/empty value by default
    var myBool bool  // false
    var myInt int  // 0
    var myString string // ""

    helloString := "hello 明"  // utf-8: 6 bytes for "hello " and 3 bytes for "明" = 9 bytes (uint8)
    helloBytes := []byte(helloString) // slice of string now represented as 9x uint8
    helloRune := []rune(helloString) // slice of string now represented as 7x int32

    fmt.Println(quote.Go())
    fmt.Printf("(%T) %v\n", singleVariable, singleVariable)
    fmt.Printf("(%T) %v\n", packageScoped, packageScoped)
    fmt.Printf("(%T) %f\n", myVariable, myVariable)  // being specific with the format
    fmt.Printf("(%T) %v\n", aString, aString)
    fmt.Printf("defaults: %T=%v, %T=%v, %T=%v\n", myBool, myBool, myInt, myInt, myString, myString)
    fmt.Printf(" string size: %v (%T of length %d)\n", helloString, helloString,  len(helloString))
    fmt.Printf("as byte size: %v (%T of length %d)\n", helloBytes, helloBytes,  len(helloBytes))
    fmt.Printf("as rune size: %v (%T of length %d)\n", helloRune, helloRune, len(helloRune))
}