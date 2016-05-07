package rename

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMoveBinaryFile(t *testing.T) {
	reCreateSourceFile()
	src := "./fixtures/file.txt"
	dest := "./fixtures/dest.txt"

	defer reCreateSourceFile()
	MoveFile(src, dest)

	destContents, err := ioutil.ReadFile(dest)
	if err != nil {
		t.Errorf("Count not copy file. Error: %v", err)
	}

	if string(destContents) != "Hello World\n" {
		t.Errorf("The destination file does not contain the same contents as the source file")
	}

	if _, err := ioutil.ReadFile(src); err == nil {
		t.Errorf("The source file still exists. It should have been deleted")
	}
}

func reCreateSourceFile() {
	contents := "Hello World\n"

	fileInfos, err := ioutil.ReadDir("./fixtures")
	if err != nil {
		os.Exit(1)
	}
	srcExists := false
	for _, fileInfo := range fileInfos {
		if fileInfo.Name() == "file.txt" {
			srcExists = true
		}
	}
	if srcExists == false {
		err := ioutil.WriteFile("./fixtures/file.txt", []byte(contents), 0644)

		if err != nil {
			os.Exit(1)
		}
	}
}
