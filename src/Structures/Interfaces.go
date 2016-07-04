package main

import "fmt"


type Animal interface {
    Speak() string
    // Walk()
}

type Dog struct {
}



func (d Dog) Speak() string {
    return "Woof"
}

type Cat struct {
}



func (c Cat) Speak() string {
    return "Meow"
}




func main() {
    var foo Animal = Dog{}
    fmt.Println(foo.Speak())
    
    var bar Animal = Cat{}
    fmt.Println(bar.Speak())
    
    
    
}