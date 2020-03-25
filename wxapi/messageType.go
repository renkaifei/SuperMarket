package wxapi

import (
	"encoding/xml"
)

type MessageType struct {
	XMLName xml.Name `xml:"xml"`
	MsgType CDATA
}

func (a *MessageType) UnMarshal(arrbyte []byte) error {
	err := xml.Unmarshal(arrbyte, a)
	return err
}
