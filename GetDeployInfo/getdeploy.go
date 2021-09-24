package vra

import (
	"encoding/json"
	"fmt"
)

func GetDepoly(token Token, Url string, auth string) (value Content, err error) {
	path := "deployment/api/deployments?$skip=0&$top=200&expand=blueprint"
	url := string(fmt.Sprintf("%s/%s", Url, path))
	res, _, code := GETReq(token, url)
	// Если не получилось подключится попробуем еще разок и сгенерируем новый ключ
	if code == 401 {
		fmt.Printf("Gen new")
		_, _, token = Gettoken(auth, Url)
		res, err, code = GETReq(token, url)
	}
	if code >= 200 && code < 300 {
		var vl Content
		err := json.Unmarshal([]byte(res), &vl)
		if err != nil {
			fmt.Printf("err %v", err)
			return vl, err
		}
		value = vl
	}
	// fmt.Println("Metric", res, code)
	return
}
