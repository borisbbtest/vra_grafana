package vra

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (d *ClinetObject) PostReq(token Token, uri string, bodys string) (result string, err error) {

	//fmt.Printf("\n%v\n", uri)
	client := &http.Client{}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest(
		"POST", uri, bytes.NewBuffer([]byte(bodys)),
	)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", token.Token_Type, token.Token_Access))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
		return
	}
	result = string(body)
	// fmt.Println("Response Body:", string(body))
	return
}

func (d *ClinetObject) GETReq(token Token, uri string) (result string, err error, httpstatus int) {

	//fmt.Printf("\n%v\n", uri)
	client := &http.Client{}
	httpstatus = 0
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest(
		"GET", uri, nil,
	)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("%s %s", token.Token_Type, token.Token_Access))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpstatus = resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Cod: %s - Couldn't parse response body. %+v", err, httpstatus)
		return
	}
	result = string(body)
	// fmt.Println("Response Body:", string(body))
	return
}
