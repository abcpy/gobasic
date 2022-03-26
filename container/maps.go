package main

import "fmt"

func main() {
	fmt.Println("Creating map")
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "glang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) //map[] m2 == empty

	var m3 map[string]int //map[] m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")

	courseName, ok := m["course"]
	fmt.Println(courseName, ok) //glang

	if courseName2, ok := m["cou"]; ok {
		fmt.Println(courseName2, ok)
	} else {
		fmt.Println("Key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok) // false

}
