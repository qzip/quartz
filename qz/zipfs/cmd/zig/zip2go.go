/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

http://www.apache.org/licenses/LICENSE-2.0

Author: Ashish Banerjee, <ashish@qzip.in>

*/
package main
import (
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/base64"
)

func main() {
	fmt.Printf("Bin file go\n")
	var binFile = flag.String("zip", "", "input zip or any file")
	var placeHolderName = flag.String("paceholder", "b64", "place holder that will be replaced by adding $[<placeholder>]\n see: zipfs/srvzip/zipcontent.go")
	var outFile = flag.String("gout", "", "golang outfile")
    var inFile = flag.String("gin", "", "golang outfile")
    flag.Parse()
	if  len(*binFile) == 0 || len(*outFile) == 0 || len(*inFile) == 0 {
		flag.PrintDefaults()
	} 
	zipBuf , err := ioutil.ReadFile(*binFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goBuf , err := ioutil.ReadFile(*inFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	b64 := base64.StdEncoding.EncodeToString(zipBuf)
	replaceFrom := "$[" + strings.TrimSpace(*placeHolderName)+"]"
	goOut := strings.Replace(string(goBuf), replaceFrom,b64,1 )
	err = ioutil.WriteFile(*outFile, []byte(goOut), 0644) // fileMode = "-rw-r--r--"
	if err != nil {
		fmt.Println(err.Error())
		
	}	
}