package messages

import (
	"encoding/xml"
)

const (
	//XsdString string type
	XsdString string = "xsd:string"
	//XsdUnsignedint uint type
	XsdUnsignedint string = "xsd:unsignedInt"
)

const (
	//SoapArray array type
	SoapArray string = "SOAP-ENC:Array"
)

const (
	//EventBootStrap first connection
	EventBootStrap string = "0 BOOTSTRAP"
	//EventBoot reset or power on
	EventBoot string = "1 BOOT"
	//EventPeriodic periodic inform
	EventPeriodic string = "2 PERIODIC"
	//EventScheduled scheduled infrorm
	EventScheduled string = "3 SCHEDULED"
	//EventValueChange value change event
	EventValueChange string = "4 VALUE CHANGE"
	//EventKicked acs notify cpe
	EventKicked string = "5 KICKED"
	//EventConnectionRequest cpe request connection
	EventConnectionRequest string = "6 CONNECTION REQUEST"
	//EventTransferComplete download complete
	EventTransferComplete string = "7 TRANSFER COMPLETE"
	//EventClientChange custom event client online/offline
	EventClientChange string = "8 CLIENT CHANGE"
)

//Message tr069 msg interface
type Message interface {
	Parse(doc *string)
	CreateXML() []byte
	GetName() string
	GetID() string
}

//Envelope tr069 body
type Envelope struct {
	XMLName   xml.Name    `xml:"soap_env:Envelope"`
	XmlnsEnv  string      `xml:"xmlns:soap_env,attr"`
	XmlnsEnc  string      `xml:"xmlns:soap_enc,attr"`
	XmlnsXsd  string      `xml:"xmlns:xsd,attr"`
	XmlnsXsi  string      `xml:"xmlns:xsi,attr"`
	XmlnsCwmp string      `xml:"xmlns:cwmp,attr"`
	Header    interface{} `xml:"soap_env:Header"`
	Body      interface{} `xml:"soap_env:Body"`
}
type EnvelopeU struct {
	XMLName   xml.Name      `xml:"Envelope"`
	XmlnsEnv  string        `xml:"soap_env,attr"`
	XmlnsEnc  string        `xml:"soap_enc,attr"`
	XmlnsXsd  string        `xml:"xsd,attr"`
	XmlnsXsi  string        `xml:"xsi,attr"`
	XmlnsCwmp string        `xml:"cwmp,attr"`
	Header    HeaderStructU `xml:"Header"`
	Body      BodyU         `xml:"Body"`
}
type BodyU struct {
	Data string `xml:",innerxml"`
}

//HeaderStruct tr069 header
type HeaderStruct struct {
	ID     IDStruct    `xml:"cwmp:ID"`
	NoMore interface{} `xml:"cwmp:NoMoreRequests,ommitempty"`
}
type HeaderStructU struct {
	ID     IDStructU   `xml:"ID"`
	NoMore interface{} `xml:"NoMoreRequests,ommitempty"`
}

//IDStruct msg id
type IDStruct struct {
	Attr  string `xml:"SOAP-ENV:mustUnderstand,attr,ommitempty"`
	Value string `xml:",chardata"`
}
type IDStructU struct {
	Attr  string `xml:"mustUnderstand,attr,ommitempty"`
	Value string `xml:",chardata"`
}

//NodeStruct node
type NodeStruct struct {
	Type  interface{} `xml:"xsi:type,attr"`
	Value string      `xml:",chardata"`
}

//EventStruct event
type EventStruct struct {
	Type   string            `xml:"SOAP-ENC:arrayType,attr"`
	Events []EventNodeStruct `xml:"EventStruct"`
}

//EventNodeStruct event node
type EventNodeStruct struct {
	EventCode  NodeStruct `xml:"EventCode"`
	CommandKey string     `xml:"CommandKey"`
}

//ParameterListStruct param list
type ParameterListStruct struct {
	Type   string                 `xml:"SOAP-ENC:arrayType,attr"`
	Params []ParameterValueStruct `xml:"ParameterValueStruct"`
}

//ParameterValueStruct param value
type ParameterValueStruct struct {
	Name  NodeStruct `xml:"Name"`
	Value NodeStruct `xml:"Value"`
}

//FaultStruct error
type FaultStruct struct {
	FaultCode   int
	FaultString string
}

//ValueStruct value
type ValueStruct struct {
	Type  string
	Value string
}
