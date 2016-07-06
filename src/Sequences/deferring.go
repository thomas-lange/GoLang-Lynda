package main

import "fmt"


func main() {
    defer fmt.Println("Some text1")
    fmt.Println("Some text2")

    defer fmt.Println("Some text3")
    defer fmt.Println("Some text4")
    fmt.Println("Some text5")
    
    
    
}