package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
*Reading files requires checking most calls for errors.
This helper will streamline our error checks below.
*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	/** Perhaps the most basic file reading task is slurping
	a file’s entire contents into memory. */
	dat, err := os.ReadFile("./data/text.txt")

	check(err)
	fmt.Print(string("file content", dat))

	/**
	You’ll often want more control over how and what parts of a file are read.
	For these tasks, start by Opening a file to obtain an os.File value.
	*/
	f, err := os.Open("./data/text.txt")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
