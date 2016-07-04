package testlib


func SumAll(elements ...int) (int, int) {
    var sum int = 0;
    for _, element := range elements {
        sum += element
    }
    
    return sum, len(elements)
}