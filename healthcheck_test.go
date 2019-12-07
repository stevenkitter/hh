package main

import (
	"strings"
	"testing"
)

func TestServer_RequestHealthCheckList(t *testing.T) {
	s := Server{Cookie: CookieStr}
	res, err := s.RequestHealthCheckList("320122194011290426")
	if err != nil {
		t.Errorf("request err %v", err)
	}
	for _, item := range res.Body {
		if strings.Contains(item.CheckDate, "2018") {
			resDetail, err := s.RequestHealthCheckDetail(item.HealthCheck)
			if err != nil {
				t.Errorf("request err %v", err)
			}
			changeRes, err := s.ChangeHealthCheckRequest(resDetail.Body, User{})
			if err != nil {
				t.Errorf("request err %v", err)
			}
			t.Logf("resDetail %d", changeRes.Code)
		}
	}
	t.Logf("request ok %v", res)
}
