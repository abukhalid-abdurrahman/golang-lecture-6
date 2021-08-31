package main

import (
	"fmt"
)

func main() {
	human := struct {
		name string
		age int
		id int
	}{
		name: "Faridun",
		age: 19,
		id: 1893
	}
	const PI = 3.14
	const radius float64 = 55.5
	var mainNumber int = int(radius * PI)
	humanId := human.id
	fmt.Printf("%+v", human)
}