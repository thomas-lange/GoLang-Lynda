package main


import (
    "fmt"
    "testlib"
    )
    
    
/**

Here a customly created library is used
1. Set GoPath
    export GOPATH=/home/ubuntu/workspace
2. go build && go install
3. import
4. use
**/
func usingASelfCreatedLibrary() {
    fmt.Println("=========================================")
    sum, val := testlib.SumAll(12, 123, 234 ,2 ,12)
    
    fmt.Printf("Sum: %v from %v values!", sum, val)
}




func main() {
    // operateOnAllIntegers("testFunc", 1, 2, 33, 4)
    

    
}