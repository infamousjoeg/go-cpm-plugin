package main

import (
	"fmt"
	"log"

	"github.com/infamousjoeg/go-cpm-plugin/pkg/cpm"
)

var (
	// XMLPath is the path to the XML file to parse
	XMLPath string
)

func main() {
	// Set the XMLPath variable to the path to the XML file
	XMLPath = "tmp/WinDomain-D-Win-DomainAdmins-Root-Operating System-WinDomain-joegarcia.dev-adm_stan-11-17-21 13-40-45.xml"

	// Parse the XML file
	AccountObject, err := cpm.ParseAccountXML(XMLPath)
	// If there's an error, handle it
	if err != nil {
		log.Fatalln(err)
	}

	// Print the account object
	fmt.Println(AccountObject)
}
