package main
import (
        "fmt"
    )

    func main() {
    z := [3][4]int{
        {0,1,2,3,},
        {4,5,6,7},
        {8,9,10,11},
    }
    val:=z[1][1]
    fmt.Println(val)
    val = z[2][3]
    fmt.Println(val)
    fmt.Println(z)

    }
