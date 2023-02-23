package dingding

import jsoniter "github.com/json-iterator/go"

// TextMessage text message struct
type TextMessage struct {
	MsgType Type `json:"msgtype"`
	At      struct {
		AtUserIDs []string `json:"atUserIds"`
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewTextMessage(text string) *TextMessage {
	return &TextMessage{
		Text: struct {
			Content string `json:"content"`
		}{
			Content: text,
		},
	}
}

func (m *TextMessage) SetAtUsers(atAll bool, userIDs []string, userMobiles []string) {
	m.At.IsAtAll = atAll
	m.At.AtUserIDs = userIDs
	m.At.AtMobiles = userMobiles
}

func (m *TextMessage) ToBytes() ([]byte, error) {
	m.MsgType = text
	bytes, err := jsoniter.Marshal(m)
	return bytes, err
}
