package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DoctorInfoRequest struct {
	Pwd string `json:"pwd"`
	Uid string `json:"uid"`
	Url string `json:"url"`
}

func NewDoctorInfoRequest(uid, pwd string) DoctorInfoRequest {
	return DoctorInfoRequest{
		Pwd: pwd,
		Uid: uid,
		Url: "logon/myRoles",
	}
}

type DoctorInfoResp struct {
	Body DoctorInfoRespBody `json:"body"`
	Code int                `json:"code"`
}

type DoctorInfoProperty struct {
	UserID     string `json:"userId"`
	CenterUnit string `json:"centerUnit"`
	RefRoleID  string `json:"refRoleId"`
	Fds        string `json:"fds"`
	RefUserID  string `json:"refUserId"`
}

type DoctorInfoRoleProperty struct {
	Version string `json:"version"`
}

type DoctorInfoRole struct {
	ID         string                 `json:"id"`
	Properties DoctorInfoRoleProperty `json:"properties"`
	Name       string                 `json:"name"`
	PageCount  int                    `json:"pageCount"`
	Type       string                 `json:"type"`
	LastModify int64                  `json:"lastModify"`
}
type DoctorInfoManageProperty struct {
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	PrintPort string `json:"printPort"`
}
type DoctorInfoManageUnit struct {
	ID         string                   `json:"id"`
	Properties DoctorInfoManageProperty `json:"properties"`
	Name       string                   `json:"name"`
	Type       string                   `json:"type"`
	Ref        string                   `json:"ref"`
	PyCode     string                   `json:"pyCode"`
}
type DoctorInfoToken struct {
	Properties     DoctorInfoProperty   `json:"properties"`
	ID             int                  `json:"id"`
	UserID         string               `json:"userId"`
	RoleID         string               `json:"roleId"`
	ManageUnitID   string               `json:"manageUnitId"`
	OrganID        string               `json:"organId"`
	RegionCode     string               `json:"regionCode"`
	LastLoginTime  string               `json:"lastLoginTime"`
	Domain         string               `json:"domain"`
	Logoff         string               `json:"logoff"`
	DisplayName    string               `json:"displayName"`
	Role           DoctorInfoRole       `json:"role"`
	UserName       string               `json:"userName"`
	ManageUnitName string               `json:"manageUnitName"`
	ManageUnit     DoctorInfoManageUnit `json:"manageUnit"`
	OrganName      string               `json:"organName"`
	RoleName       string               `json:"roleName"`
}

type DoctorInfoRespBody struct {
	UserPhoto string            `json:"userPhoto"`
	Tokens    []DoctorInfoToken `json:"tokens"`
}

/**
 *  登陆功能
 */
func (s *Server) DoctorInfo(uid, pwd string) (DoctorInfoResp, error) {
	url := "http://32.33.1.123:8082/pkehr/logon/myRoles"
	cli := http.Client{}
	reqData := NewDoctorInfoRequest(uid, pwd)
	bit, _ := json.Marshal(&reqData)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bit))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Cookie", s.Cookie)
	resp, err := cli.Do(request)
	setCookie := resp.Header.Get("Set-Cookie")
	s.Cookie = setCookie
	if err != nil {
		return DoctorInfoResp{}, err
	}
	defer resp.Body.Close()
	bits, _ := ioutil.ReadAll(resp.Body)
	var result DoctorInfoResp
	err = json.Unmarshal(bits, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
