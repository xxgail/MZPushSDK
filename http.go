package MZPushSDK

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func PostResult(requestPath string, fields map[string]string, retries int) (*Result, error) {
	result, err := postReq(requestPath, fields)
	if result.Code == Success && err == nil {
		return result, nil
	}
	// 重试
	for i := 0; i < retries; i++ {
		result, err = postReq(requestPath, fields)
		if result.Code == Success && err == nil {
			break
		}
	}
	return result, err
}

func postReq(requestPath string, fields map[string]string) (*Result, error) {
	form := url.Values{}
	for k, v := range fields {
		form.Add(k, v)
	}
	param := form.Encode()
	req, err := http.NewRequest("POST", fmt.Sprintf("%s", baseHost()+requestPath), strings.NewReader(param))
	fmt.Println("requestUrl:", baseHost()+requestPath)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("bbbb", string(body))
	res.Body.Close()

	if err != nil {
	}

	var result = &Result{}
	err = json.Unmarshal(body, result)

	return result, nil
}

func baseHost() string {
	return ProductionHost
}
