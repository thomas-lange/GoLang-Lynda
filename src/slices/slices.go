
/**
*   Slices and arrays
*
**/

package main;


import (
    "fmt"
    "os"
    "bufio"
    "sort"
    "strconv"
    "strings"

)



func iterationWithArray() {
        var array [10]int;
    
    for i := 0; i < 10; i++ {
        array[i] = i * i
    }
    
    
    for i := 0; i < 10; i++ {
        fmt.Println(array[i])
    }
}



func iterationWithSlice() {
    var slice = []int{};
    
    for i := 0; i < 10; i++ {
        slice = append(slice, i)
    }
    for i := 0; i < 10; i++ {
        fmt.Println(slice[i])
    }
}

func efficientSlices() {
    // Slices can be made more efficient, with giving them a initial size
    // with the MAKE function
    
    numbers := make([]int, 1, 5)
    for i := 0; i < len(numbers); i++ {
        fmt.Println(numbers[i])
    }
    
}


func getUserInput() int {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a number: ")
    str, _ := reader.ReadString('\n')
    
    str = strings.TrimSpace(str)
    val, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println(err)
        panic("Error occured!")
    }
    return val
}







func main(){
    // iterationWithSlice()
    // fmt.Println("=============================")
    // efficientSlices()
    slice := make([]int, 0, 20)
    var val int = 0
    for val >= 0 {
        val = getUserInput()
        if val > 0 {
            slice = append(slice, val)
        }
    }
    fmt.Println(slice)
    
    sort.Ints(slice)

    fmt.Println(slice)
    
}