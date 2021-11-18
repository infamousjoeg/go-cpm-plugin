package types

import "encoding/xml"

// Location metadata for account object
type Location struct {
	XMLName    xml.Name `xml:"location"`
	SafeName   string   `xml:"safename"`
	FolderName string   `xml:"foldername"`
	ObjectName string   `xml:"objectname"`
}

// Policy metadata for account object
type Policy struct {
	XMLName           xml.Name `xml:"policy"`
	ChangePassTimeout int      `xml:"changepasstimeout"`
}

// Property metadata for account object
type Properties struct {
	XMLName                 xml.Name `xml:"properties"`
	ManagementType          string   `xml:"ManagementType"`
	LogonDomain             string   `xml:"LogonDomain"`
	DeviceType              string   `xml:"DeviceType"`
	PolicyID                string   `xml:"PolicyID"`
	Address                 string   `xml:"Address"`
	UserName                string   `xml:"UserName"`
	CreationMethod          string   `xml:"CreationMethod"`
	LastSuccessVerification int      `xml:"LastSuccessVerification"`
	SequenceID              int      `xml:"SequenceID"`
	LastSuccessChange       int      `xml:"LastSuccessChange"`
	RetriesCount            int      `xml:"RetriesCount"`
	ResetImmediately        string   `xml:"ResetImmediately"`
}

// ExtraPass1Location metadata for account object's logon account
type ExtraPass1Location struct {
	XMLName    xml.Name `xml:"location"`
	SafeName   string   `xml:"safename"`
	FolderName string   `xml:"foldername"`
	ObjectName string   `xml:"objectname"`
}

// ExtraPass1Properties metadata for account object's logon account
type ExtraPass1Properties struct {
	XMLName                 xml.Name `xml:"properties"`
	ManagementType          string   `xml:"ManagementType"`
	LogonDomain             string   `xml:"LogonDomain"`
	DeviceType              string   `xml:"DeviceType"`
	PolicyID                string   `xml:"PolicyID"`
	Address                 string   `xml:"Address"`
	UserName                string   `xml:"UserName"`
	CreationMethod          string   `xml:"CreationMethod"`
	LastSuccessVerification int      `xml:"LastSuccessVerification"`
	SequenceID              int      `xml:"SequenceID"`
	LastSuccessChange       int      `xml:"LastSuccessChange"`
	RetriesCount            int      `xml:"RetriesCount"`
	CPMDisabled             string   `xml:"CPMDisabled"`
	CPMStatus               string   `xml:"CPMStatus"`
}

// ExtraPass1 metadata for account object's logon account
type ExtraPass1 struct {
	XMLName    xml.Name             `xml:"extrapass1"`
	Location   ExtraPass1Location   `xml:"location"`
	Properties ExtraPass1Properties `xml:"properties"`
}

// ExtraPass2Location metadata for account object's reconcile account
type ExtraPass2Location struct {
	XMLName    xml.Name `xml:"location"`
	SafeName   string   `xml:"safename"`
	FolderName string   `xml:"foldername"`
	ObjectName string   `xml:"objectname"`
}

// ExtraPass2Properties metadata for account object's reconcile account
type ExtraPass2Properties struct {
	XMLName                 xml.Name `xml:"properties"`
	ManagementType          string   `xml:"ManagementType"`
	LogonDomain             string   `xml:"LogonDomain"`
	DeviceType              string   `xml:"DeviceType"`
	PolicyID                string   `xml:"PolicyID"`
	Address                 string   `xml:"Address"`
	UserName                string   `xml:"UserName"`
	CreationMethod          string   `xml:"CreationMethod"`
	LastSuccessVerification int      `xml:"LastSuccessVerification"`
	SequenceID              int      `xml:"SequenceID"`
	LastSuccessChange       int      `xml:"LastSuccessChange"`
	RetriesCount            int      `xml:"RetriesCount"`
	CPMDisabled             string   `xml:"CPMDisabled"`
	CPMStatus               string   `xml:"CPMStatus"`
}

// ExtraPass2 metadata for account object's reconcile account
type ExtraPass2 struct {
	XMLName    xml.Name             `xml:"extrapass2"`
	Location   ExtraPass2Location   `xml:"location"`
	Properties ExtraPass2Properties `xml:"properties"`
}

// ExtraPass3Location metadata for account object's enable account
type ExtraPass3Location struct {
	XMLName    xml.Name `xml:"location"`
	SafeName   string   `xml:"safename"`
	FolderName string   `xml:"foldername"`
	ObjectName string   `xml:"objectname"`
}

// ExtraPass3Properties metadata for account object's enable account
type ExtraPass3Properties struct {
	XMLName                 xml.Name `xml:"properties"`
	ManagementType          string   `xml:"ManagementType"`
	LogonDomain             string   `xml:"LogonDomain"`
	DeviceType              string   `xml:"DeviceType"`
	PolicyID                string   `xml:"PolicyID"`
	Address                 string   `xml:"Address"`
	UserName                string   `xml:"UserName"`
	CreationMethod          string   `xml:"CreationMethod"`
	LastSuccessVerification int      `xml:"LastSuccessVerification"`
	SequenceID              int      `xml:"SequenceID"`
	LastSuccessChange       int      `xml:"LastSuccessChange"`
	RetriesCount            int      `xml:"RetriesCount"`
	CPMDisabled             string   `xml:"CPMDisabled"`
	CPMStatus               string   `xml:"CPMStatus"`
}

// ExtraPass3 metadata for account object's enable account
type ExtraPass3 struct {
	XMLName    xml.Name             `xml:"extrapass3"`
	Location   ExtraPass3Location   `xml:"location"`
	Properties ExtraPass3Properties `xml:"properties"`
}

// ExtraPass metadata for account object's logon, reconcile, and enable accounts
type ExtraPass struct {
	XMLName    xml.Name   `xml:"extrapass,omitempty"`
	ExtraPass1 ExtraPass1 `xml:"extrapass1,omitempty"`
	ExtraPass2 ExtraPass2 `xml:"extrapass2,omitempty"`
	ExtraPass3 ExtraPass3 `xml:"extrapass3,omitempty"`
}

// AccountObject metadata for account object
type AccountObject struct {
	XMLName    xml.Name   `xml:"settings"`
	Location   Location   `xml:"location"`
	Policy     Policy     `xml:"policy"`
	Properties Properties `xml:"properties"`
	ExtraPass  ExtraPass  `xml:"extrapass,omitempty"`
}
