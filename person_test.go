package main

import "testing"

func TestServer_PersonalInfo(t *testing.T) {
	s := Server{Cookie: "JSESSIONID=359E2121A53735BAB90146981485FB14; 01144817=%u738B%u5E73@photo/01144817.jpg"}
	res, err := s.PersonalInfo("341424200903276111")
	if err != nil {
		t.Errorf("err %v", err)
	}
	t.Logf("ok %d", res.Code)
}
