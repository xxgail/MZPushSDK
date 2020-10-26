package MZPushSDK

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
)

func generateSign(params map[string]string, appSecret string) string {
	keys := make([]string, 0, len(params))
	for k, _ := range params {
		keys = append(keys, k)
	}

	str := ""
	sort.Strings(keys)
	for _, v := range keys {
		str += fmt.Sprintf("%v=%v", v, params[v])
	}
	str += appSecret
	fmt.Println("ssss", str)

	u := md5.New()
	u.Write([]byte(str))
	return hex.EncodeToString(u.Sum(nil))
}
