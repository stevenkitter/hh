package main

import "testing"

func TestServer_Login(t *testing.T) {
	s := Server{}
	res, err := s.DoctorInfo("01144817", "1")
	if err != nil {
		t.Errorf("err %v", err)
	}
	t.Logf("ok %d", res.Code)
}
