package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

// Custom writer that adhere to the Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just write this number of bytes:", len(bs))

	return len(bs), nil
}

// ==========================================================
//                Assignment 1 Interfaces
// ==========================================================

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println("The area of this shape is:", s.getArea())
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))
	// io.Copy(os.Stdout, res.Body)

	lw := logWriter{}
	io.Copy(lw, res.Body)

	t := triangle{base: 12, height: 7}
	s := square{sideLength: 23}

	printArea(t)
	printArea(s)
}
