package main


import "fmt"



type Dog struct {
    Breed string
    Weight int
    Sound string
}


// ->(d Dog) = Binding to Type
func (d Dog) Speak() {
    fmt.Println(d.Sound)
}

// Use Reference instead of Clone
func (d *Dog) SpeakWPointer() {
    fmt.Println(d.Sound)
    d.Sound = "Whatever!"
}


func main() {
    testDog := Dog{"Foo", 34, "Woof"}
    testDog.Speak()
    
    testDog.Sound = "Woooooof"
    testDog.SpeakWPointer()
    testDog.SpeakWPointer()
}

