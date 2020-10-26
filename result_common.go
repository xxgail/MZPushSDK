package MZPushSDK

type Result struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	Value    string `json:"value"`
	Redirect string `json:"redirect"`
	MsgId    string `json:"msgId"`
}
