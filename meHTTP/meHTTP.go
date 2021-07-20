package meHTTP

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

type HttpObj struct {
	HostName string
	TimeOut time.Duration
	Header map[string]string
	Body  io.Reader
	Method string
}
func httpclient() *http.Client{
	var client = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	return client
}

func (h *HttpObj )Run()(*http.Response,error) {
	var client = httpclient()
	client.Timeout = h.TimeOut
	request, err := http.NewRequest(h.Method,h.HostName,h.Body)
	if err != nil {
		return &http.Response{},err
	}
	for key,value := range h.Header{
		request.Header.Add(key,value)
	}
	resp, err := client.Do(request)
	if err != nil {
		return &http.Response{},err
	}
	return resp,nil
}
