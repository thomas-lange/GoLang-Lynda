package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strings"
    // "math/big"
    )
    
    
func main() {
    url := "http://services.explorecalifornia.org/json/tours.php"
    
    resp, err := http.Get(url)
    
    if err != nil {
        panic(err)
    }
    
    
    fmt.Printf("ResponseType: %T", resp)
    
    defer resp.Body.Close()
    
    bytes, err := ioutil.ReadAll(resp.Body)
    
    var content string = string(bytes)
    
    // fmt.Println(content)
    tours := SimpleParseJsonToursString(content)
    PrintTours(tours)
    
    tours = ParseJsonToursString(content)
    PrintTours(tours)
    
}


type Tour struct {
    Name, Price string
}



func ParseJsonToursString(jsonString string) []Tour {
    tours := make([]Tour, 0, 20)
    
    decoder := json.NewDecoder(strings.NewReader(jsonString))
    _, err := decoder.Token()
    
    checkError(err)
    
    var tour Tour
    for decoder.More() {
        err := decoder.Decode(&tour)
        checkError(err)
        tours = append(tours, tour)
    }
    return tours
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