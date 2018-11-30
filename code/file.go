package main

import "fmt"
import "os"

func main() {
	file := filecreate("D:/Go/workspace/filewrite.txt")
	defer closefile(file)
	filewrite(file)

}

func filecreate(filename string) *os.File {
	fmt.Println("Creating file...... ")
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return file
}
func filewrite(filename *os.File) {
	fmt.Println("Writing now")
	fmt.Fprintln(filename, "adminadminadmin")
}
func closefile(filename *os.File) {
	fmt.Println("Close the file")
	filename.Close()
}
