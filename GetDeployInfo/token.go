package vra

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func (d *ClinetObject) Gettoken(Auth string, urlvro string) (token string, timestamp int, tokenst Token, err error) {
	client := &http.Client{}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// fmt.Printf("%v", fmt.Sprintf("%s/csp/gateway/am/api/login?access_token", urlvro))
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

	var resj map[string]interface{}
	json.Unmarshal([]byte(string(body)), &resj)
	json.Unmarshal([]byte(string(body)), &tokenst)
	token = fmt.Sprintf("%v", resj["access_token"])
	timestamp = int(resj["expires_in"].(float64))

	//fmt.Sprintf("%v", reflect.TypeOf(resj["validity"]))
	//ts = fmt.Sprintf("%v", tp)
	WriteTokenFile(string(body))
	return
}
func (d *ClinetObject) OpenTokenFile() (tokenfile Token, e error) {
	jsonFile, err := os.Open("/tmp/token_last.json")
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

func WriteTokenFile(text string) error {
	file, err := os.Create("/tmp/token_last.json")
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
