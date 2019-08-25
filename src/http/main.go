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
}
