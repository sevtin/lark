package xhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	HTTP_REQUEST_TIME_OUT_SECOND = time.Second * 10
)

type HeaderOption struct {
	Key   string
	Value string
}

func getQueryUrl(params map[string]interface{}) (query string) {
	if params == nil || len(params) == 0 {
		return
	}
	var (
		buffer bytes.Buffer
		key    string
		val    interface{}
	)
	buffer.WriteString("?")
	for key, val = range params {
		if val == nil {
			continue
		}
		buffer.WriteString(fmt.Sprintf("%s=%v&", key, val))
	}
	buffer.Truncate(buffer.Len() - 1)
	return buffer.String()
}

func Get(url string, params map[string]interface{}, headerOptions ...*HeaderOption) (buf []byte, err error) {
	var (
		req    *http.Request
		option *HeaderOption
	)
	url += getQueryUrl(params)
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return
	}
	for _, option = range headerOptions {
		req.Header.Set(option.Key, option.Value)
	}

	var (
		client = http.Client{Timeout: HTTP_REQUEST_TIME_OUT_SECOND}
		resp   *http.Response
	)
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	buf, err = ioutil.ReadAll(resp.Body)
	return
}

func Post(url string, params map[string]interface{}, headerOptions ...*HeaderOption) (buf []byte, err error) {
	var (
		jsonBuf []byte
		req     *http.Request
		option  *HeaderOption
	)
	if len(params) > 0 {
		jsonBuf, err = json.Marshal(params)
		if err != nil {
			return
		}
	}
	req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBuf))
	if err != nil {
		return
	}
	for _, option = range headerOptions {
		req.Header.Set(option.Key, option.Value)
	}
	var (
		client = &http.Client{Timeout: HTTP_REQUEST_TIME_OUT_SECOND}
		resp   *http.Response
	)
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	buf, err = ioutil.ReadAll(resp.Body)
	return
}
