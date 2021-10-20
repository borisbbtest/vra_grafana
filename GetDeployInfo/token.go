package vra

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func (d *ClinetObject) GetRTWeb(Auth string, urlvro string) (reftok RefreshToken, err error) {
	client := &http.Client{}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// fmt.Printf("%v", fmt.Sprintf("%s/csp/gateway/am/api/login?access_token", urlvro))

	// Временный токен
	req, err := http.NewRequest(
		"POST", fmt.Sprintf("%s/csp/gateway/am/api/login?access_token", urlvro),
		bytes.NewBuffer([]byte(Auth)),
	)
	// добавляем заголовки
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}
	//log.Println("Response Body:", string(body))

	json.Unmarshal(body, &reftok)
	WriteTokenFile(string(body), "/tmp/ref_token_last.json")

	//rft := fmt.Sprintf("%s", resj["refresh_token"])
	return
}
func (d *ClinetObject) GeBaseToken(rft RefreshToken, urlvro string) (body []byte, httpstatus int, err error) {
	client := &http.Client{}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	refresh_token := map[string]interface{}{
		"refreshToken": rft.RefToken,
	}
	data_rft, _ := json.Marshal(refresh_token)
	fmt.Printf("%v", string(data_rft))
	req, err := http.NewRequest(
		"POST", fmt.Sprintf("%s/iaas/api/login", urlvro),
		bytes.NewBuffer([]byte(string(data_rft))),
	)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	httpstatus = resp.StatusCode
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}
	return
}

func (d *ClinetObject) Gettoken(Auth string, urlvro string) (token string, timestamp int, tokenst Token, err error) {

	// fmt.Printf("%v", fmt.Sprintf("%s/csp/gateway/am/api/login?access_token", urlvro))
	var rft RefreshToken

	tokenGetLast, err := d.OpenRefTokenFile("/tmp/ref_token_last.json")
	if err != nil {
		rft, err = d.GetRTWeb(Auth, urlvro)
		//fmt.Printf("Not found file  %v get service %v last get service  %v", nowTimeStpam, tokenst.Expires_in)
	} else {
		rft = tokenGetLast
		fmt.Printf("Open file")
	}
	// Получаем токен
	body, httpstatus, err := d.GeBaseToken(rft, urlvro)
	if httpstatus >= 300 {
		fmt.Printf("Gen new ref token")
		rft, err = d.GetRTWeb(Auth, urlvro)
		// fmt.Printf("%s  ", rft)
		body, httpstatus, err = d.GeBaseToken(rft, urlvro)
	}
	err = json.Unmarshal(body, &tokenst)
	err = WriteTokenFile(string(body), "/tmp/token_last.json")
	//fmt.Printf("Response Body: %s \n", string(body))
	//fmt.Sprintf("%v", reflect.TypeOf(resj["validity"]))
	//ts = fmt.Sprintf("%v", tp)
	if httpstatus >= 300 && err == nil {
		err = errors.New(fmt.Sprintf("HTTP CODE %s ", httpstatus))
	}
	return
}

func (d *ClinetObject) OpenTokenFile(path string) (tokenfile Token, e error) {
	jsonFile, err := os.Open(path)
	e = err
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//var resj map[string]interface{}

	json.Unmarshal(byteValue, &tokenfile)
	//fmt.Printf(" token %v ", tokenfile)
	return
}

func (d *ClinetObject) OpenRefTokenFile(path string) (tokenfile RefreshToken, e error) {
	jsonFile, err := os.Open(path)
	e = err
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//var resj map[string]interface{}

	json.Unmarshal(byteValue, &tokenfile)
	//fmt.Printf(" token %v ", tokenfile)
	return
}

func WriteTokenFile(text string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}
	return file.Sync()
}
