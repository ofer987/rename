package fileutils

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMoveFile(t *testing.T) {
	src := "file.txt"
	dest := "dest.txt"
	createFile(src)

	defer createFile(src)
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

func TestMoveFileToExistingFile(t *testing.T) {
	reCreateSourceFile()
	src := "file.txt"
	dest := "dest.txt"

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

func createFile(path) {
	contents := "Hello World\n"

	fileInfos, err := ioutil.ReadDir("./")
	if err != nil {
		os.Exit(1)
	}
	srcExists := false
	for _, fileInfo := range fileInfos {
		if fileInfo.Name() == path {
			srcExists = true
		}
	}
	if srcExists == false {
		err := ioutil.WriteFile(path, []byte(contents), 0644)

		if err != nil {
			os.Exit(1)
		}
	}
}
