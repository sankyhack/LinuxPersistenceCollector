package copy

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Copies files given on "destingation path" from "src"
func CopyFile(src, dst string) error {
	source, err := os.Open(src)

	if err != nil {
		return err
	}

	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		fmt.Println("unable to create file", dst)
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)

	return err
}

//this functions creates directory with same name as that of our source path
// also iterates through directory for files and send it to CopyFile function

func CopyDir(src, dst string) error {
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		destPath := filepath.Join(dst, path)

		if info.IsDir() {
			//fmt.Println("Isndie ISDir", info.Name())
			if err := os.MkdirAll(destPath, info.Mode().Perm()); err != nil {
				fmt.Printf("error creating directory %q: %v\n", destPath, err)
				return err
			}
			//	return nil
		} else {

			/* if !info.Mode().IsRegular() {
				fmt.Printf("skipping non-regular file %q\n", path)
				return nil
			} */

			if err := CopyFile(path, destPath); err != nil {
				fmt.Printf("error copying file %q to %q: %v\n", path, destPath, err)
				return err
			}
		}
		return nil
	})
	return err
}
