package main

import "fmt"

func main() {
	m := map[string]int{"golang": 1, "python": 2, "java": 3}
	delete(m, "python")
	fmt.Printf("%d,%d,%d", m["golang"], m["python"], m["java"], m["ruby"])
}
