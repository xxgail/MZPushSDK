package MZPushSDK

import (
	"strings"
)

func MessageSend(message *Message, appId string, pushIds []string, appSecret string) (*Result, error) {
	fields := message.GetFields()
	fields["appId"] = appId
	fields["pushIds"] = strings.Join(pushIds, ",")
	fields["sign"] = generateSign(fields, appSecret)

	return PostResult(MessageURL, fields, 1)
}
