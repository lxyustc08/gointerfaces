// Package bytecounter is the code of interface examples in
// go programming language
package bytecounter

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Bytecounter is the test type of interface
type Bytecounter int

func (c *Bytecounter) Write(p []byte) (int, error) {
	*c += Bytecounter(len(p))
	return len(p), nil
}

// Test is the test function for interface
func Test() {
	var c Bytecounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

// Wordcounter is the counter for words
type Wordcounter int

// Linecounter is the counter for lines
type Linecounter int

func (c *Wordcounter) Write(p []byte) (int, error) {
	count, err := retCount(p, bufio.ScanWords)
	*c += Wordcounter(count)
	return len(p), err
}

func (c *Linecounter) Write(p []byte) (int, error) {
	count, err := retCount(p, bufio.ScanLines)
	*c += Linecounter(count)
	return len(p), err
}

func retCount(p []byte, f bufio.SplitFunc) (count int, err error) {
	s := string(p)
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(f)
	count = 0
	for scanner.Scan() {
		count++
	}
	if err = scanner.Err(); err != nil {
		err = fmt.Errorf("scanning err %v", err)
	}
	return count, err
}

// TestF is the code of exercise 7.1 in go programming language
func TestF() {
	var x Wordcounter
	var y Linecounter
	x.Write([]byte("hello world ustc"))
	fmt.Println(x)
	y.Write([]byte("hello world ustc,\n am i test\n, line count test\n"))
	fmt.Println(y)
}

type bytecounter struct {
	io.Writer
	countting int64
}

func (bc *bytecounter) Write(p []byte) (n int, err error) {
	n, err = bc.Writer.Write(p)
	bc.countting += int64(n)
	return n, err
}

// CountingWriter accepts an io.Writer, then returns a new Writer that wraps the original,
// and a pointer to an int64 variable that any monent contains the number of bytes written
// to the new writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWriter := &bytecounter{w, 0}
	return newWriter, &newWriter.countting
}
