/*
Copyright (c) 2018, QzIP Blockchain Technology LLP

http://www.apache.org/licenses/LICENSE-2.0

Author: Ashish Banerjee, <ashish@qzip.in>


*/
package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"os"
	"qz/zipfs/srvzip"
)

func main() {
	mzip, err := srvzip.NewMemZip(srvzip.MemZippedB64)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	files, cmnt, err := mzip.GetFiles()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(cmnt)
	fmt.Println("----- DIRS ----")
	prnDirs(files)
	fmt.Println("----- JSON START------")
	prnFiles(files)
	fmt.Println("----- JSON END------")

}
func prnDirs(files []*zip.File) {
	for _, f := range files {
		if f.FileInfo().IsDir() {
			fmt.Println(f.Name)
		}
	}
}
func prnFiles(files []*zip.File) {
	b, err := json.Marshal(files)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
