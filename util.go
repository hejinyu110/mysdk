package mysdk

import "encoding/json"

const (
	MessageTypeTxt = "RC:TxtMsg"
	MessageTypeVoice = "RC:VoiceMsg"
)

type RCMessage interface {
	JsonFormat() ([]byte, error)
	
}
type RCMessageUser struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Extra string `json:"extra"`
}
type RCTxtMesage struct {
	User RCMessageUser `json:"user"`
	Content string `json:"content"`
	Extra string `json:"extra"`
}


func (msg RCTxtMesage) JsonFormat() ([]byte, error)  {
	return json.Marshal(msg)
}

