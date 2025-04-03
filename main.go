package main
import "fmt"
 
 func main() {
	// Slice 
	var score = []int{100, 200, 300, 400, 500}
	fmt.Println(score)
	score = append(score, 600, 700, 800)
	fmt.Sprintln(score)
	fmt.Println(score, len(score), cap(score))
 }
    

 	
 	