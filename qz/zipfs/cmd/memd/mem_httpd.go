/*
Copyright (c) 2018-2019, QzIP Blockchain Technology LLP

http://www.apache.org/licenses/LICENSE-2.0

Author: Ashish Banerjee, <ashish@qzip.in>

  A Micro Lite Http File Server
  http://golang.org/pkg/net/http/#example_FileServer

*/
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"qz/zipfs/srvzip"
	"time"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	inPort := flag.String("port", "7070", "Input Port Number, default 7070")
	dbg := flag.Bool("debug", false, "set true for enabling Http Debug")
	domain := flag.String("domain", "", "(optional) domain for Letsencrypt autocert, like:\n -domain=\"wzip.in\"")
	flag.Parse()
	srvzip.Debug = *dbg
	httpPort := ":" + *inPort
	fmt.Printf("Port %s\n", httpPort)

	mzs, err := srvzip.NewMemZipServer(srvzip.MemZippedB64)
	if err != nil {
		log.Fatal(err)
	}

	if len(*domain) == 0 {
		go func() {
			time.Sleep(200 * time.Millisecond)
			srvzip.OpenBrowser("http://localhost:" + *inPort)

		}()
		log.Fatal(http.ListenAndServe(httpPort, mzs.MemZipServer()))
	} else {
		log.Fatal(http.Serve(autocert.NewListener(*domain), mzs.MemZipServer()))
	}
	/*

		mzip, err := srvzip.NewMemZip(srvzip.MemZippedB64)
		if err != nil {
			log.Fatal(err)
		}
		// Simple static webserver:
		log.Fatal(http.ListenAndServe(httpPort, mzip.FileServer()))
	*/

}
