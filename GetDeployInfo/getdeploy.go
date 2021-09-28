package vra

import (
	"encoding/json"
	"fmt"
)

func (d *ClinetObject) GetDepoly(token Token, Url string, auth string, stings_page ...int) (value Content, err error) {
	top := 50
	skip := 0
	if len(stings_page) > 0 {
		top = stings_page[0]
	}
	if len(stings_page) >= 2 {
		skip = stings_page[1]
	}
	//https://712b-vra.mpk.lcl/deployment/api/swagger/swagger-ui.html?urls.primaryName=2020-08-25#/Deployments/getDeploymentsV3UsingGET
	path := string(fmt.Sprintf("deployment/api/deployments?$skip=%v&$top=%v&expand=project", skip, top))
	url := string(fmt.Sprintf("%s/%s", Url, path))
	res, _, code := d.GETReq(token, url)
	// Если не получилось подключится попробуем еще разок и сгенерируем новый ключ
	if code == 401 {
		fmt.Printf("Gen new")
		_, _, token, _ = d.Gettoken(auth, Url)
		res, err, code = d.GETReq(token, url)
	}

	if code >= 200 && code < 300 {
		var vl Content
		// var ty map[string]interface{}
		// ttp := json.Unmarshal([]byte(res), &ty)
		// fmt.Printf("%v", string([]byte(res)))
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
