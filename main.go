package main
import "fmt"
 
 func main() {
	// Slice 
	
	nomes := [4]string{"livia", "laura", "gabrieli", "kelly"}
	
	fmt.Println(nomes)
	rangeOne := nomes[:2]
	fmt.Println(rangeOne)
	rangeTwo := nomes[2:]
	fmt.Println(rangeTwo)
	rangeThree := nomes[1:3]
	fmt.Println(rangeThree)
	
 }
    

 	
 	