package structs

import (
	"encoding/json"
	"errors"

	"main/functions/sessions"
)
type MessageType struct {
	Server string
	Basic  string
}
type MessageUser struct {
	Name string `json:"name"`
	ID string `json:"id"`
}
type Message struct {
	User *MessageUser `json:"user"`
	Text string `json:"text"`
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

func (msg *Message) MapToUser(umap map[string]any) {
	by, err := json.Marshal(umap)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(by, &msg); err != nil {
		panic(err)
	}
}
func (msg *Message) ToMap() map[string]any {
	by, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	UMap := map[string]any{}
	if err := json.Unmarshal(by, &UMap); err != nil {
		panic(err)
	}
	return UMap
}

func NewMessage(text string, session sessions.Session) *Message {
	user := session.Get("user").(map[string]any)


	return &Message{
		Text: text,
		User: &MessageUser{
			Name: user["name"].(string),
			ID: user["id"].(string),

		},

	}
}