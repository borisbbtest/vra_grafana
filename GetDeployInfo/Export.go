package vra

import (
	"encoding/json"
	"fmt"
	"time"
)

func (d *ClinetObject) Export(p ...int) (result []Content, err error) {
	Auth := map[string]interface{}{
		"username": d.Login,
		"password": d.Password,
	}
	page := 0
	if len(p) > 0 {
		page = p[0]
	}
	top := 100
	if len(p) >= 1 {
		top = p[1]
	}
	nowTimeStpam := time.Now().UnixNano() / int64(time.Millisecond)
	var urlvro string = ""
	if urlvro = string(fmt.Sprintf("http://%s", d.FQDN)); d.SSL == true {
		urlvro = string(fmt.Sprintf("https://%s", d.FQDN))
	}
	data, _ := json.Marshal(Auth)
	var tokenst Token
	tokenGetLast, e := d.OpenTokenFile()
	if e != nil {
		_, _, tokenst = d.Gettoken(string(data), urlvro)
		fmt.Printf("Not found file  %v get service %v last get service  %v", nowTimeStpam, tokenst.Expires_in)
	} else {
		tokenst = tokenGetLast
		fmt.Printf("Open file")
	}

	var datares Content
	tr := true
	skip := 0
	for tr {
		if page > 0 {
			datares, err = d.GetDepoly(tokenst, urlvro, string(data), top, (top * page))
			result = append(result, datares)
			break
		}
		datares, err = d.GetDepoly(tokenst, urlvro, string(data), top, (top * skip))
		if err != nil {
			fmt.Printf("Erroe %v", err)
			return
		}
		for i := 0; i <= len(datares.Data)-1; i++ {
			// fmt.Printf("\n %v) %v", i, datares.Data[i])
		}
		///fmt.Printf("\n%v %v %v %v ", datares.TotalElements, datares.TotalPages, tr, (top * skip))
		skip++
		result = append(result, datares)
		tr = !datares.Last

	}
	d.Result = result
	return
}
