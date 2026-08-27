package main

import "fmt"

const greet = "Hello "

func Hello(name string) string {
	return greet + name + "!"
}

func main() {
	fmt.Println(Hello(""))
}
