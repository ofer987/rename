package rename

import (
	"io/ioutil"
	"log"
	"os"
)

func MoveFile(src string, dest string) {
	srcFile := file(src)
	srcFileInfo := info(srcFile)
	srcMode := srcFileInfo.Mode()
	srcContents := readFile(src)

	ioutil.WriteFile(dest, srcContents, srcMode)
	os.Remove(src)
}

func file(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open %v. Error: %v", path, err)
	}

	return file
}

func info(file *os.File) os.FileInfo {
	fileInfo, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}

	return fileInfo
}

func readFile(path string) []byte {
	contents, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return contents
}
