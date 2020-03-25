package wxapi

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"sort"
	"strings"
)

type CDATA struct {
	Text string `xml:",cdata"`
}

type TextMessage struct {
	XMLName      xml.Name `xml:"xml"`
	URL          CDATA
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int
	MsgType      CDATA
	Content      CDATA
	MsgId        int64
}

func (a *TextMessage) Marshal() (result string, err error) {
	arrbyte, err := xml.Marshal(a)
	result = string(arrbyte)
	return result, err
}

func ValidateSignature(signature string, timestamp string, nonce string) bool {
	tempArr := []string{token, timestamp, nonce}
	sort.Strings(tempArr)
	temp := strings.Join(tempArr, "")
	result := fmt.Sprintf("%x", sha1.Sum([]byte(temp)))
	if result == signature {
		return true
	} else {
		return false
	}
}

func ListenMessage(msg []byte) (txtMsg *TextMessage, err error) {
	v := &TextMessage{}
	err = xml.Unmarshal(msg, &v)
	return v, err
}
