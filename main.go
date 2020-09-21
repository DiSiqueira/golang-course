package main

import "fmt"

type (
	Printer interface {
		Print(string)
	}

	CLIPrinter struct {}
)

func (CLIPrinter) Print(text string) {
	fmt.Println(text)
}

func main() {
	p := &CLIPrinter{}
	PrintWithSuffix(p, "hello!")
}

func PrintWithSuffix(p Printer, text string) {
	p.Print("suffix:"+text)
}


