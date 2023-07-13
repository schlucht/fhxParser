package main

import "fmt"

type Desktop struct {
	Title string
}

func main() {
	var desktop Desktop
	desktop.Title = "Meine App"
	fmt.Println(desktop.Title)
}
