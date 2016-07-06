package main

import (
    "fmt"
    "io"
    "bufio"
    "bytes"
    "io/ioutil"
    "os"
    "math/rand"
    "time"
)




func main() {

    
    content := generateRandomString(80)
    ln, _ := WriteStringToFile(content, "../../ressources/fromString.txt")
    
    for i := 0; i < 10; i++ {
        content := GetStringFromFile("../../ressources/repeatedInsert.txt") +
            "\n" +
            generateRandomString(80)
            WriteStringToFile(content, "../../ressources/repeatedInsert.txt")
    }
    
    fmt.Println(ReadFileWithIo("../../ressources/repeatedInsert.txt"))
    
    
    fmt.Printf("Written %v Characters to File!\n", ln)
    
    WriteBytesToFile([]byte(content), "../../ressources/fromBytes")
    
}

/**
    This does not work
**/
func ReadFileWithIo(fileName string) string {
    var s string
    var e error
    var buffer bytes.Buffer
    file, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }
    r := bufio.NewReader(file)
    s, e = Readln(r)
    for e == nil {
        s, e = Readln(r)
        buffer.WriteString(s + "\n")
    }
    return buffer.String()
}



func Readln(r *bufio.Reader) (string, error) {
    var buffer bytes.Buffer
    var (isPrefix bool = true
        err error = nil
        line []byte 
    )
    for isPrefix && err == nil {
        line, isPrefix, err = r.ReadLine()
        buffer.WriteString(string(line))
    }
    return buffer.String(), err
}


func WriteStringToFile(content, fileName string) (ln int, err error) {
    file, err := os.Create(fileName)
    checkError(err)
    defer file.Close()
    if err != nil {
        return 0, err
    }
    
    ln, err = io.WriteString(file, content)
    checkError(err)
    return ln, err
}


func WriteBytesToFile(content []byte, fileName string) {
    ioutil.WriteFile(fileName, content, 0777)
}



func generateRandomString(length int) string {
    rand.Seed(time.Now().UnixNano())
    
    var validCharacters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    var validCharSize = len(validCharacters)
    
    var buffer = make([]rune, length)
    
    for i := 0; i < length; i++ {
        buffer[i] = validCharacters[rand.Intn(validCharSize)]
    }
    return string(buffer)
    
}

func GetStringFromFile(file string) string {
    content, err := ioutil.ReadFile(file)
    if err != nil {
        return ""
    }
    return string(content)
}




func checkError(err error) {
    if err != nil {
        fmt.Println(error.Error)
        panic(err)
    }
}