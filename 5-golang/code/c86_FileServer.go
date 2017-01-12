/**
* File Server
**/

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// deliver files from the directory /
	fileServer := http.FileServer(http.Dir("/"))

	// register the handler and deliver requests to it
	err := http.ListenAndServe(":8090", fileServer)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
