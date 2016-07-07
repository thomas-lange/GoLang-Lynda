package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    )
    
    
    
type Tour struct {
    Name, Price string
}

type Hello struct {}


func (h Hello) ServeHTTP (writer http.ResponseWriter, reader *http.Request) {
    tours := GetJson()
    PrintToursToWriter(tours, writer)
}
    
func main() {
    var h Hello
    err := http.ListenAndServe("localhost:4004", h)
    checkError(err)

    
}




func PrintToursToWriter(tours []Tour, writer http.ResponseWriter) {
    for index, item := range tours {
        fmt.Fprintf( writer, "Tour %d >> Price: %v, Name: %v\n", index, item.Price, item.Name )
    }
}



func GetJson() []Tour {
    url := "http://services.explorecalifornia.org/json/tours.php"
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    bytes, err := ioutil.ReadAll(resp.Body)
    var content string = string(bytes)
    
    return SimpleParseJsonToursString(content)
    
}



func SimpleParseJsonToursString(jsonString string) []Tour {
    var tours []Tour
    json.Unmarshal([]byte(jsonString), &tours)
    

    return tours
    
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}


func PrintTours(tours []Tour) {
    fmt.Println();
    for index, item := range tours {
        fmt.Printf("Tour %d >> Price: %v, Name: %v\n", index, item.Price, item.Name )
    }
    fmt.Println();
}