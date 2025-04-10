package main

import "fmt"

func main() {
    var numeros [5]int
    var soma int

    fmt.Println("Digite 5 números:")

    for i := 0; i < 5; i++ {
        fmt.Scan(&numeros[i])
        soma += numeros[i]
    }

    fmt.Println("Soma:", soma)
}
 