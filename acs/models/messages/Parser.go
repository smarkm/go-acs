package messages

//ParseXML parse xml msg
func ParseXML(data []byte) (msg Message, err error) {
	// doc := xmlx.New()
	// doc.LoadBytes(data, nil)
	// bodyNode := doc.SelectNode("*", "Body")
	// if bodyNode != nil {
	// 	var name string
	// 	if len(bodyNode.Children) > 1 {
	// 		name = bodyNode.Children[1].Name.Local
	// 	} else {
	// 		name = bodyNode.Children[0].Name.Local
	// 	}
	// 	switch name {
	// 	case "Inform":
	// 		msg = NewInform()
	// 		msg.Parse(doc)
	// 	case "GetParameterValuesResponse":
	// 		msg = &GetParameterValuesResponse{}
	// 		msg.Parse(doc)
	// 	case "SetParameterValuesResponse":
	// 		msg = &SetParameterValuesResponse{}
	// 		msg.Parse(doc)
	// 	case "DownloadResponse":
	// 		msg = &DownloadResponse{}
	// 	case "TransferComplete":
	// 		msg = &TransferComplete{}
	// 	case "GetRPCMethodsResponse":
	// 		msg = &GetRPCMethodsResponse{}
	// 	case "RebootResponse":
	// 		msg = &RebootResponse{}
	// 	}
	// 	if msg != nil {
	// 		msg.Parse(doc)
	// 	}
	// }
	return msg, err
}
