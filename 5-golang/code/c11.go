/**
* ASN .1
 */
package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)

func main() {
	mdata, err := asn1.Marshal(true)
	checkError(err)

	var n bool
	_, err = asn1.Unmarshal(mdata, &n)
	checkError(err)

	fmt.Println("Afer marshal/unmarshal: ", n)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
