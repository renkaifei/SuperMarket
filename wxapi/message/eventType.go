package message

import (
	"encoding/xml"
)

type EventType struct {
	Event CDATA
}

func (a *EventType) Unmarshal(data []byte) error {
	err := xml.Unmarshal(data, a)
	return err
}
