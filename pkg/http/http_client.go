package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	RequestMethodGet  = "GET"
	RequestMethodPost = "POST"
)

type ResponseErr struct {
	Err  error
	Code int
}

type Response struct {
	ResponseErr
	Data interface{}
}

func Get(url string) (buf []byte, err error) {
	var (
		client http.Client
		resp   *http.Response
	)
	client = http.Client{Timeout: 5 * time.Second}
	resp, err = client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	buf, err = ioutil.ReadAll(resp.Body)
	return
}

//application/json; charset=utf-8
func Post(url string, data interface{}, timeOutSecond int) (respBuf []byte, err error) {
	var (
		jsonBuf []byte
		req     *http.Request
		client  *http.Client
		resp    *http.Response
	)
	jsonBuf, err = json.Marshal(data)
	if err != nil {
		return
	}
	req, err = http.NewRequest(RequestMethodPost, url, bytes.NewBuffer(jsonBuf))
	if err != nil {
		return
	}
	req.Close = true
	req.Header.Add("content-type", "application/json; charset=utf-8")

	client = &http.Client{Timeout: time.Duration(timeOutSecond) * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respBuf, err = ioutil.ReadAll(resp.Body)
	return
}

func PostReturn(url string, input, output interface{}, timeOut int) (err error) {
	var (
		respBuf []byte
	)
	respBuf, err = Post(url, input, timeOut)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(respBuf, output); err != nil {
		return
	}
	return
}
