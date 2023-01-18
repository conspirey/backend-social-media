package structs

import (
	"errors"
)
type MessageType struct {
	Server string
	Basic  string
}

func (mt *MessageType) Apply(mtype string) (success error) {
	if mtype == "server" {
		mt.Server = mtype
		return nil
	} else if mtype == "basic" {
		mt.Basic = mtype
		return nil
	}
	return errors.New("Invalid message type")
}

func (mt *MessageType) CreateServer() {

}