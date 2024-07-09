/*
Copyright (c) 2021, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

/*
DB based dynamically configurable web service

HTTP server component, first checks the URL against publicNamespace, if found serves the file


*/

package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"qz/commands"
	"reflect"
	"strconv"

	"golang.org/x/crypto/acme/autocert"
)

/*
TODO:
	configLoader
	contentLoader
*/

/***** NOT USED ****
//HTTPHandlerFactory creates the HttpHandler
type HTTPHandlerFactory interface {
	commands.ComponentFactory
	Create(ctx context.Context, cfg map[string]interface{}, errChan chan error) (component http.Handler)
}
******/

// SrvHTTP interface that all commands must implement
type SrvHTTP struct {
}

// Name implements qz/commands/Commander interface
func (srv *SrvHTTP) Name() string {
	return reflect.TypeOf(*srv).String()
}

// Help implements qz/commands/Commander interface
func (srv *SrvHTTP) Help() string {
	return `
   {
	   "command": "cmd.SrvHTTP",
	   "debug": true,
	   "domain": "traceit.info",
	   "port": 443,
	   "ssl" : true,
	   "dbURL":               "postgresql://root@localhost:26257?sslcert=%2Fhome%2Fashish%2FAB%2FDev2019%2FLabs%2Fwlt2019-certs%2Fclient.root.crt&sslkey=%2Fhome%2Fashish%2FAB%2FDev2019%2FLabs%2Fwlt2019-certs%2Fclient.root.key&sslmode=verify-full&sslrootcert=%2Fhome%2Fashish%2FAB%2FDev2019%2FLabs%2Fwlt2019-certs%2Fca.crt",
	   "dbDriver":            "postgres",
	   "publicNamespace": "web",
	   "HTTPmuxHandler": "MuxHandler", 
	   "loadConfigHandler":  "dbConfigLoader",
		"helpers": [
			{
				"component": "dbConfigLoader"
				"param": {}
			}
			
		]
   }   
 
   `
}

const (
	ctxHandlerHelper = "HTTPmuxHandler"
	paramPort        = "port"
	paramDomain      = "domain"
)

// Exec Executes the configuration dynamically from db
func (srv *SrvHTTP) Exec(ctx context.Context, cfg map[string]interface{}, errCh chan error) {
	// load http handler factory
	key, err := srv.getHandler(cfg)
	if err != nil {
		errCh <- commands.NewFatalError(err.Error())
		return
	}
	handler := ctx.Value(key)
	if handler == nil {
		errCh <- commands.NewFatalError("cmd.SrvHTTP.Exec: ctx helper " + key + " not found")
		return
	}
	var mux http.Handler
	ok := true
	if mux, ok = handler.(http.Handler); !ok {
		err := commands.NewFatalError(fmt.Sprintf("cmd.SrvHTTP.Exec: expected helper of type http.Handler but got %v", reflect.TypeOf(handler).String()))
		errCh <- err
		return
	}
	port := 7070
	sport := os.Getenv("PORT") // for cloud run compatibility
	if sport != "" {
		if port, err = strconv.Atoi(sport); err != nil {
			errCh <- err
			return
		}
	}

	domain := ""

	iport, ok := cfg[paramPort]
	if ok {
		port, ok = iport.(int)
		if !ok {
			errCh <- commands.NewFatalError(fmt.Errorf("cmd.SrvHTTP.Exec: param port %v not int", iport).Error())
			return
		}
	}

	idomain, ok := cfg[paramDomain]
	if ok {
		domain, ok = idomain.(string)
		if !ok {
			errCh <- commands.NewFatalError(fmt.Errorf("cmd.SrvHTTP.Exec: param domain %v not a string", idomain).Error())
			return
		}
	}

	if domain != "" {
		errCh <- fmt.Errorf("cmd.SrvHTTP.Exec: INFO attempting to serve ssl on doamin %v", domain)
		err := http.Serve(autocert.NewListener(domain), mux)
		if err != nil {
			errCh <- commands.NewFatalError(err.Error())
		}
	} else {
		errCh <- fmt.Errorf("cmd.SrvHTTP.Exec: INFO attempting to serve localhost port: %v", port)
		uport := ":" + strconv.Itoa(port)
		err := http.ListenAndServe(uport, mux)
		if err != nil {
			errCh <- commands.NewFatalError(err.Error())
		}

	}
}

func (srv *SrvHTTP) getHandler(cfg map[string]interface{}) (string, error) {
	handlerFact, ok := cfg[ctxHandlerHelper]
	if !ok {
		return "", fmt.Errorf("cmd.SrvHTTP.createHandler: Http handler helper param %v not specified", ctxHandlerHelper)
	}
	helper, ok := handlerFact.(string)
	if !ok {
		return "", fmt.Errorf("cmd.SrvHTTP.createHandler:  param %v not a string", ctxHandlerHelper)
	}

	return helper, nil
}
