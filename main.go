package main

import "fmt"

func main() {
    age := 45
	fmt.Println(age <= 50 )
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("menor que 30 anos")
	} else if age < 40 {
		fmt.Println("menor que 40 anos")
	} else { 
		fmt.Println("não é mesnor que 40 anos")
	}
	names := []string{"beatriz", "livia", "joana", "mel"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continue apos posição", index, "e valor", value)
			continue
		}
		if index > 2 { 
			fmt.Println("sair apos" , index)
		}
		fmt.Println("valor: ", value)
	}
}
 