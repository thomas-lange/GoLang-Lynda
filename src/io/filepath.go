package main

import (
    "fmt"
    "os"
    "path/filepath"
    "io/ioutil"
    )
    
func main() {
    var cursor = "."
    var files []os.FileInfo
    
    for true {
        
        cursor, _ = filepath.Abs(cursor)
        files, _ = ioutil.ReadDir(cursor)
        var directories = make(map[int]os.FileInfo)
        fmt.Println("Currently in Path:", cursor)
        
        numberOfDirectories := 0
        for _, item := range files {
            if (item.IsDir()) {
                numberOfDirectories++
                directories[numberOfDirectories] = item
                fmt.Println(numberOfDirectories, "--", item.Name())
            }
            
        }
        if numberOfDirectories > 0 {
            fmt.Println("---------------------------------------------------")
            fmt.Println("Press 0 to move down or choose one of the directories above!")
        } else {
            fmt.Println("Press 0 to move down in the direcory tree!")
        }
        
selection, _ := GetIntegerInput()
        if selection == 0 {
            cursor = cursor + "/.."
        } else {
            cursor = cursor + "/" + directories[selection].Name()
        }

    }
    
}


func GetIntegerInput() (int, error) {
    var i int
    _, err := fmt.Scanf("%d", &i)
    return i, err
}
