package MZPushSDK

import "encoding/json"

type Message struct {
	Fields map[string]string
}

type MessageJson struct {
	NoticeBarInfo    NoticeBarInfo    `json:"noticeBarInfo"`
	NoticeExpandInfo NoticeExpandInfo `json:"noticeExpandInfo"`
	ClickTypeInfo    ClickTypeInfo    `json:"clickTypeInfo"`
	PushTimeInfo     PushTimeInfo     `json:"pushTimeInfo"`
}

type NoticeBarInfo struct {
	NoticeBarType int    `json:"noticeBarType"`
	Title         string `json:"title"`
	Content       string `json:"content"`
}

type NoticeExpandInfo struct {
	NoticeExpandType    int    `json:"noticeExpandType"` // 0-标准、1-文本
	NoticeExpandContent string `json:"noticeExpandContent"`
}

type ClickTypeInfo struct {
	ClickType       int               `json:"clickType"` // 点击动作 0-打开应用（默认）、1-打开应用页面、2-打开URI页面、3-应用客户端自定义
	Url             string            `json:"url"`
	Parameters      map[string]string `json:"parameters"`
	Activity        string            `json:"activity"`
	CustomAttribute string            `json:"customAttribute"`
}

type PushTimeInfo struct {
	OffLine   int `json:"offLine"`
	ValidTime int `json:"validTime"`
}

func InitMessage(title string, desc string) *Message {
	var messageJson MessageJson
	messageJson = MessageJson{
		NoticeBarInfo: NoticeBarInfo{
			NoticeBarType: 0,
			Title:         title,
			Content:       desc,
		},
		NoticeExpandInfo: NoticeExpandInfo{
			NoticeExpandType:    0,
			NoticeExpandContent: "",
		},
		ClickTypeInfo: ClickTypeInfo{
			ClickType:  0,
			Parameters: map[string]string{},
			Activity:   "",
		},
		PushTimeInfo: PushTimeInfo{
			OffLine:   0,
			ValidTime: 0,
		},
	}
	messageJsonStr, _ := json.Marshal(messageJson)
	var fields map[string]string
	fields = map[string]string{
		"messageJson": string(messageJsonStr),
	}
	return &Message{
		Fields: fields,
	}
}

func (m *Message) GetFields() map[string]string {
	return m.Fields
}
