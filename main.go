package main

import "log"

func main() {

	name, name2, age := a("Lee", 20)
	log.Println(name + name2 + string(age))
}

func a(name string, age int) (string, string, int) {

	return name, name, 30

}
