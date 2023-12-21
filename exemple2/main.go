
package main

import (
	"fmt"
	"Desktop\M1 WMD\Go Language\exemple\dictionary"
)


func main() {

    myDictionary := make(dictionary.Dictionary)

    
    myDictionary.Add("golang", "A programming language developed by Google.")
    myDictionary.Add("python", "A high-level programming language.")


    fmt.Println("Definition of golang:", myDictionary.Get("golang"))

 
    myDictionary.Remove("python")

    fmt.Println("Dictionary:", myDictionary.List())


    myDictionary.ExamplePrint()
}