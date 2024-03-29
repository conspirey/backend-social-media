package structs

import (
	"encoding/json"
	"errors"
	"regexp"

	// "fmt"

	"main/functions/sessions"
)

type MessageType struct {
	Server      string
	Basic       string
	ServerTimer string
}
type MessageUser struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
type Message struct {
	User          *MessageUser   `json:"user"`
	Text          string         `json:"text"`
	Type          string         `json:"type"`
	ServerMessage *ServerMessage `json:"server_message"`
	TextStyle     string         `json:"text_style,omitempty"`
	NameStyle     string         `json:"name_style,omitempty"`
}
type ServerMessage struct {
	Timer   int64  `json:"timer"`
	Delay   int64  `json:"delay"`
	Message string `json:"message"`
}

var (
	ServerMessageTimerRegex = regexp.MustCompile(`({timer}|\[timer\])`)
)

func (msg *Message) SetUser(session sessions.Session) {
	user := session.Get("user").(map[string]any)
	// fmt.Println()
	msg.User.ID = user["id"].(string)
	msg.User.Name = user["name"].(string)

}
func (mt *MessageType) Apply(mtype string) (success error) {
	if mtype == "server" {
		mt.Server = mtype
		return nil
	} else if mtype == "basic" {
		mt.Basic = mtype
		return nil
	} else if mtype == "server_timer" {
		mt.ServerTimer = mtype
		return nil
	}
	return errors.New("invalid message type")
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

func NewMessage(text string, msgType string, session sessions.Session, ServerMessageV *ServerMessage) *Message {
	user := session.Get("user").(map[string]any)
	message := &Message{
		Text: text,
		Type: msgType, // basic|server
		User: &MessageUser{
			Name: user["name"].(string),
			ID:   user["id"].(string),
		},
	}
	if msgType == "server_timer" {
		message.ServerMessage = ServerMessageV
	}
	return message
}
func (msg *MessageUser) MapToUser(umap map[string]any) {
	by, err := json.Marshal(umap)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(by, &msg); err != nil {
		panic(err)
	}
}
func (msg *MessageUser) ToMap() map[string]any {
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
