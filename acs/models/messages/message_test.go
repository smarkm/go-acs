package messages_test

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"go-acs/acs/models/messages"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"text/scanner"
	"time"
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

func TestReboot(t *testing.T) {

}

func TestACS(t *testing.T) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		line, _ := bufio.NewReader(r.Body).ReadBytes(io.SeekEnd)
		fmt.Printf("line: %s\n", line)
		if r.ContentLength == 0 {
			//reboot := messages.Reboot{}
			//w.Write(reboot.CreateXML())
			return
		}

		envelope := messages.EnvelopeU{}
		err := xml.Unmarshal(line, &envelope)
		if err != nil {
			t.Error(err)
		}
		if strings.Contains(envelope.Body.Data, "cwmp:Inform") {
			inform := messages.NewInform()
			inform.Parse(envelope.Body.Data)
			informResponse := messages.InformResponse{MaxEnvelopes: 1, ID: "FFFFFF123456"}
			w.Write(informResponse.CreateXML())
			fmt.Printf("%s\n", informResponse.CreateXML())
		}
	})

	go http.ListenAndServe(":7547", nil)
	fmt.Println("================")
	var s scanner.Scanner
	s.Init(os.Stdin)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}
	time.Sleep(10000 * time.Second)

}
