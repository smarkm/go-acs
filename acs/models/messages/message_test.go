package messages_test

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"go-acs/acs/models/messages"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestInform(t *testing.T) {

	//informResponse := messages.InformResponse{}
	//fmt.Println(string(informResponse.CreateXML()))
	xmls, _ := xml.Marshal(&messages.Envelope{XmlnsEnv: "http://schemas.xmlsoap.org/soap/envelope/", Body: "sdf"})
	fmt.Println(string(xmls))
	e := messages.EnvelopeU{}
	fmt.Println(xml.Unmarshal(xmls, &e))
	fmt.Println(e)
}

func TestACS(t *testing.T) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		line, _ := bufio.NewReader(r.Body).ReadBytes(io.SeekEnd)
		envelope := messages.EnvelopeU{}
		err := xml.Unmarshal(line, &envelope)
		//fmt.Println(envelope.Body)
		if strings.Contains(envelope.Body.Data, "cwmp:Inform") {
			inform := messages.Inform{}
			inform.Parse(envelope.Body.Data)
		}
		fmt.Println(err)
	})

	log.Fatal(http.ListenAndServe(":7547", nil))
}
