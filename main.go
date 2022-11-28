package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"weavelab.xyz/puffin"
)

func main() {
	service := "data-seeding-api"

	servicePath, err := GetPathToDirectory("./integrationtests", service)
	if err != nil {
		log.Fatalln(err.Error())
	}

	exec := puffin.NewOsExec()
	testCmd := exec.Command("/bin/sh", "-c", "mkdir -p /tmp/reports && go test ./... -json -count 1 -tags integration 2>&1")
	testCmd.SetDir(servicePath)

	out, _ := testCmd.Output()

	fmt.Printf("%v\n", string(out))
}

func GetPathToDirectory(root, targetDir string) (string, error) {
	var pathToDir string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && info.Name() == targetDir {
			pathToDir = path
		}
		return nil
	})
	return pathToDir, err
}

func GetPathToDirectory_WalkDir(root, targetDir string) (string, error) {
	var pathToDir string
	err := filepath.WalkDir(root, func(path string, dir os.DirEntry, err error) error {
		if dir.IsDir() && dir.Name() == targetDir {
			pathToDir = path
		}
		return nil
	})
	return pathToDir, err
}
