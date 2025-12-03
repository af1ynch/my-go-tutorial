package main

import "fmt"

type Reader interface {
	Read() (string, error)
}

type Writer interface {
	Write(s string) (string, error)
}

type File struct {
	FilePath string
}

func NewFile(fp string) *File {
	return &File{
		FilePath: fp,
	}
}

func (f *File) Read() (string, error) {
	return fmt.Sprintf("文件%s读取成功", f.FilePath), nil
}

func (f *File) Write(s string) (string, error) {
	return fmt.Sprintf("%s已被成功写入文件\n", s), nil
}

func main() {

	var f1 Reader = new(File)
	f2 := NewFile("test.txt")
	if f, ok := f1.(Reader); ok {
		fmt.Printf("f1 is %v, ok is %v\n", f, ok)
	} else {
		fmt.Printf("f1 is %v, ok is %v\n", f, ok)
	}
	fmt.Printf("f2 is %v\n", f2)

}
