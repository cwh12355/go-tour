package main

import (
	"fmt"
	"os"
)

func main() {
	f := creatFile("./temp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func creatFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "Caowenhao NB NB NB !!!")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
