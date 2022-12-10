package services

//Message holds a message from a service.
// A text message should only hold content.
// Adding fields, a title, an image, or a color creates a rich message.
// If ServerID is not specified, presume ChannelID to be a DM channel with a user and use Msg* methods.
type Message struct {
	UserID  string            `json:"userID,omitempty"`
	MessageID string          `json:"messageID,omitempty"`
	ChannelID string          `json:"channelID,omitempty"`
	ServerID  string          `json:"serverID,omitempty"`
	Title     string          `json:"title,omitempty"`
	Content   string          `json:"content,omitempty"`
	Image     string          `json:"image,omitempty"`
	Color     int            `json:"color,omitempty"`
	Fields    []*MessageField `json:"fields,omitempty"`
	Context   interface{}     `json:"context,omitempty"`
}

func NewMessage() *Message {
	return &Message{}
}
func (msg *Message) SetTitle(title string) *Message {
	msg.Title = title
	return msg
}
func (msg *Message) SetContent(content string) *Message {
	msg.Content = content
	return msg
}
func (msg *Message) SetColor(clr int) *Message {
	msg.Color = clr
	return msg
}
func (msg *Message) SetImage(img string) *Message {
	msg.Image = img
	return msg
}
func (msg *Message) AddField(name, value string) *Message {
	if msg.Fields == nil {
		msg.Fields = make([]*MessageField, 0)
	}
	msg.Fields = append(msg.Fields, &MessageField{Name: name, Value: value})
	return msg
}

type MessageField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}