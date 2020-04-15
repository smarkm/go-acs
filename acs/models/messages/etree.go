package messages

import (
	"encoding/json"
	"fmt"

	"github.com/beevik/etree"
)

//FindValue RT
func FindValue(e *etree.Element, tag string) string {
	rs := e.FindElement(tag)
	if rs != nil {
		return rs.Text()
	}
	return ""
}

func Json(obj interface{}) {
	data, err := json.Marshal(obj)
	fmt.Println(string(data), err)
}
