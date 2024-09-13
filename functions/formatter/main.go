package main

import "fmt"

func reformat(message string, formatter func(string) string) string {
	once := formatter(message)
	twice := formatter(once)
	thrice := formatter(twice)
	prefix := "TEXTIO:"
	return prefix + thrice
}

func myFormatter(s string) string {
	return s + "!"
}

func main() {
	str := reformat("hello world", myFormatter)

	fmt.Println(str)
}
