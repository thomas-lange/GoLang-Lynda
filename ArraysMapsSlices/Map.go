package main;



import (
    "fmt"
    )
    
    
    
func main() {
    item := make(map[string]string)
    
    item["foo"] = "bar"
    item["bar"] = "foo"
    item["123"] = "123"
    item["123"] = "323"
    
    
    fmt.Println(item)
    
}