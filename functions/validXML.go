package functions

import "encoding/xml"

func IsXMLValid(data []byte) bool {
	return xml.Unmarshal(data, new(interface{})) == nil
}