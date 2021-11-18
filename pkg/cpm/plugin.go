package cpm

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	"github.com/infamousjoeg/go-cpm-plugin/pkg/types"
)

// ParseAccountXML parses account data from a CPM-provided XML file.
func ParseAccountXML(xmlPath string) (types.AccountObject, error) {
	// Read the XML file
	xmlFile, err := os.Open(xmlPath)
	// If there's an error, return it
	if err != nil {
		return types.AccountObject{}, err
	}

	// Close the file at the end
	defer xmlFile.Close()

	// Read the file into a byte array
	byteValue, _ := ioutil.ReadAll(xmlFile)

	var AccountObject types.AccountObject

	// Unmarshal the byte array into the AccountObject struct
	err = xml.Unmarshal(byteValue, &AccountObject)
	// If there's an error, return it
	if err != nil {
		return types.AccountObject{}, err
	}

	return AccountObject, nil
}
