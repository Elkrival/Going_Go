package main

import (
	"bufio"
	"fmt"
)

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Println(len(p))
	return len(p), nil
}

func main() {
	fmt.Println("Unbuffered I/O")
	w := new(Writer)
	w.Write([]byte{'a'})
	w.Write([]byte{'b'})
	w.Write([]byte{'c'})
	w.Write([]byte{'d'})
	fmt.Println("Buffered I/O")
	bw := bufio.NewWriterSize(w, 3)
	bw.Write([]byte{'a'})
	fmt.Println(bw.Available(), "available")
	bw.Write([]byte{'b'})
	bw.Write([]byte{'c'})
	fmt.Println(bw.Available(), "available")
	bw.Write([]byte{'d'})
	fmt.Println(bw.Available(), "available")
	bw.Write([]byte{'e'})
	fmt.Println(bw.Available(), "available")
	err := bw.Flush()
	if err != nil {
		panic(err)
	}
	fmt.Println("Buffered I/O 2")
	bcw := bufio.NewWriterSize(w, 2)
	bcw.Write([]byte{'a'})
	bcw.Write([]byte{'b'})
	bcw.Write([]byte{'c'})
	er := bcw.Flush()
	if er != nil {
		panic(er)
	}
	fmt.Println("Large Writes")
	lw := new(Writer)
	blw := bufio.NewWriterSize(lw, 3)
	blw.Write([]byte("try it"))
	w1 := new(Writer1);
	bw1 := bufio.NewWriterSize(w1, 2)
	bw1.Write([]byte("abcd"))
	fmt.Println(bw1.Available(), "available")
	bw1.Write([]byte("efgh"))
	fmt.Println(bw1.Available(), "available")
	w2 := new(Writer2)
	bw1.Reset(w2)
	bw1.Write([]byte("ef"))
	fmt.Println(bw1.Available(), "available")
	bw1.Flush()
	wbrs := new(Writer)
	bwbrs := bufio.NewWriterSize(wbrs, 10)
	fmt.Println(bwbrs.Buffered())
	bwbrs.WriteRune("¢")
	fmt.Println(bwbrs.Buffered())
	bwbrs.WriteRune("ł")
	fmt.Println(bwbrs.Buffered())
	bwbrs.WriteString("bbbbb")
	fmt.Println(bwbrs.Buffered())
}
func (*Writer) LWrites(p []byte) (n int, err error) {
	fmt.Printf("%q\n", p)
	return len(p), nil
}

/*
======== Reset ==========
To allocate memory we need to reset the buffer for garbage collection and
for resources to run smoothly
*/

type Writer1 int
type Writer2 int
func (*Writer1) Write(p []byte) (n int, err error) {
 fmt.Printf("writer#1: %q\n", p)
 // We are returning the lenght of the byte or null the null is for error
 return len(p), nil
}
func (*Writer2) Write(p[]byte)(n int, err error) {
	fmt.Printf("writer#2 : %q\n", p)
	return len(p), nil
}
