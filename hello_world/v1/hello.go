package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishEndSuffix = "!"

func Hello(name string) string {
	return englishHelloPrefix + name + englishEndSuffix
}

func main() {
	fmt.Println(Hello("Adi"))
}
