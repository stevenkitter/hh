package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func (s *Server) RequestJson(b []byte) ([]byte, error) {
	url := "http://32.33.1.123:8082/pkehr/*.jsonRequest"
	cli := http.Client{}
	request, _ := http.NewRequest("POST", url, bytes.NewReader(b))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Cookie", s.Cookie)
	resp, err := cli.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
