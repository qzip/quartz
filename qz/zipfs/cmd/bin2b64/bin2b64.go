/*
Copyright (c) 2018, QzIP Blockchain Technology LLP

http://www.apache.org/licenses/LICENSE-2.0

Author: Ashish Banerjee, <ashish@qzip.in>


*/
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	undef := "undefined"
	fmt.Printf("Bin file to/From Base64 encode/decode util\n")
	var binFile = flag.String("bin", undef, "Output file, Base64 string")
	var b64File = flag.String("b64", undef, "Input bin file")
	var outType = flag.String("out", "b64", "output type bin or b64(default) ")
	flag.Parse()

	if strings.EqualFold(undef, *binFile) || strings.EqualFold(undef, *b64File) {
		fmt.Printf("Please specify input and output file. Example \nbin2b64 -out b64/bin -bin my.zip -b64 b64.txt\n")
		return
	}
	err := genB64(*binFile, *b64File, strings.EqualFold("b64", *outType))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Success generated\n")
}
func genB64(binFile string, b64File string, b64Out bool) (err error) {
	inFile := binFile
	outFile := b64File
	if !b64Out {
		inFile = b64File
		outFile = binFile
	}
	fileIn, err := os.Open(inFile)
	if err != nil {
		return err
	}
	defer fileIn.Close()

	b, err := ioutil.ReadAll(fileIn)
	if err != nil {
		return
	}
	if b64Out {
		b64 := base64.StdEncoding.EncodeToString(b)
		err = ioutil.WriteFile(outFile, []byte(b64), 0644) // fileMode = "-rw-r--r--"
	} else {
		data, err := base64.StdEncoding.DecodeString(string(b))
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(outFile, data, 0644)
	}
	return
}
