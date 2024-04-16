package main

import "fmt"

func main() {
	fmt.Println("Maps")
	user := map[string]string{
		"nome" : "Enzzo",
		"sobrenome": "Jaco",
	}
	fmt.Println(user)
	delete(user, "sobrenome")
	user["idade"] = "21"
	fmt.Println(user)
}