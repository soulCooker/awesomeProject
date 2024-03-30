package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

type Singer interface {
	Sing()
}

type Mark struct {
}

func (m Mark) Sing() {
	fmt.Println("I'm singing")
}

func (m Mark) Name() string {
	return "Mark"
}

func (m Mark) String() string {
	return "My name is Mark"
}

type Weekday int

func main() {
	var singer Singer
	mark := &Mark{}
	singer = mark
	singer.Sing()
	fmt.Println(mark.Name())

	output := crc32.NewIEEE()
	mulWriter := io.MultiWriter(output, os.Stdout)
	_, err := fmt.Fprint(mulWriter, singer)
	if err != nil {
		return
	}
	fmt.Println(output.Sum32())

	bufferWriter := new(bytes.Buffer)
	dumper := hex.Dumper(bufferWriter)

	dumper.Write([]byte(mark.String()))

	dumper.Close()

	str, err := bufferWriter.ReadString('\n')

	fmt.Println(str)
}
