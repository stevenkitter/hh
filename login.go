package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LoginRequest struct {
	HttpMethod string `json:"httpMethod"`
	Url        string `json:"url"`
}

func NewLoginRequest(urt int, uid, pwd string) LoginRequest {
	return LoginRequest{
		HttpMethod: "POST",
		Url:        fmt.Sprintf("logon/myApps?urt=%d&uid=%s&pwd=%s&deep=3", urt, uid, pwd),
	}
}

type LoginResp struct {
	Body LoginRespBody `json:"body"`
	Code int           `json:"code"`
}

type LoginApp struct {
	ID         string           `json:"id"`
	Name       string           `json:"name"`
	Type       string           `json:"type"`
	Properties LoginAppProperty `json:"properties"`
	PageCount  int              `json:"pageCount"`
	LastModify int64            `json:"lastModify"`
	FullID     string           `json:"fullId"`
}

type LoginAppProperty struct {
	EntryName string `json:"entryName"`
}
type LoginRespBody struct {
	MyDesktop                    string      `json:"myDesktop"`
	HypertensionRiskStartMonth   string      `json:"hypertensionRiskStartMonth"`
	UserDomain                   string      `json:"userDomain"`
	DebilityShowType             string      `json:"debilityShowType"`
	SysMessage                   bool        `json:"sysMessage"`
	UserPhoto                    string      `json:"userPhoto"`
	ServerDateTime               string      `json:"serverDateTime"`
	ChildrenDieAge               int         `json:"childrenDieAge"`
	OldPeopleAge                 int         `json:"oldPeopleAge"`
	AreaGridType                 string      `json:"areaGridType"`
	Fds                          string      `json:"fds"`
	ChildrenHealthCheckShowType  string      `json:"childrenHealthCheckShowType"`
	DiabetesRiskStartMonth       string      `json:"diabetesRiskStartMonth"`
	HypertensionMode             int         `json:"hypertensionMode"`
	HypertensionStartMonth       string      `json:"hypertensionStartMonth"`
	DiabetesPrecedeDays          int         `json:"diabetesPrecedeDays"`
	OldPeopleStartMonth          string      `json:"oldPeopleStartMonth"`
	DiabetesType                 string      `json:"diabetesType"`
	TabRemove                    bool        `json:"tabRemove"`
	AreaGridShowType             string      `json:"areaGridShowType"`
	Domain                       string      `json:"domain"`
	RoleType                     string      `json:"roleType"`
	RvcMode                      int         `json:"rvcMode"`
	DiabetesMode                 int         `json:"diabetesMode"`
	PhisActiveYW                 string      `json:"phisActiveYW"`
	DiabetesStartMonth           string      `json:"diabetesStartMonth"`
	OldPeopleMode                int         `json:"oldPeopleMode"`
	CenterUnit                   string      `json:"centerUnit"`
	PageCount                    int         `json:"pageCount"`
	PregnantMode                 int         `json:"pregnantMode"`
	Appwelcome                   bool        `json:"appwelcome"`
	CenterUnitName               string      `json:"centerUnitName"`
	HealthCheckType              string      `json:"healthCheckType"`
	PostnatalVisitType           string      `json:"postnatalVisitType"`
	DiabetesDelayDays            int         `json:"diabetesDelayDays"`
	TumourHighRiskStartMonth     string      `json:"tumourHighRiskStartMonth"`
	TumourPatientVisitStartMonth string      `json:"tumourPatientVisitStartMonth"`
	PsychosisType                string      `json:"psychosisType"`
	HypertensionType             string      `json:"hypertensionType"`
	PsychosisStartMonth          string      `json:"psychosisStartMonth"`
	CheckFollowUpShowType        string      `json:"checkFollowUpShowType"`
	Apps                         []LoginApp  `json:"apps"`
	TabNumber                    interface{} `json:"tabNumber"`
	ServerDate                   string      `json:"serverDate"`
	PhisActive                   bool        `json:"phisActive"`
	Postnatal42DayType           string      `json:"postnatal42dayType"`
	ChildrenRegisterAge          int         `json:"childrenRegisterAge"`
}

func (s *Server) Login(urt int, uid, pwd string) (LoginResp, error) {
	url := "http://32.33.1.123:8082/pkehr/logon/myApps"
	cli := http.Client{}
	url += fmt.Sprintf("?urt=%d&uid=%s&pwd=%s&deep=3", urt, uid, pwd)
	reqData := NewLoginRequest(urt, uid, pwd)
	bit, _ := json.Marshal(&reqData)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bit))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Cookie", s.Cookie)
	resp, err := cli.Do(request)
	setCookie := resp.Header.Get("Set-Cookie")
	if setCookie != "" {
		s.Cookie = setCookie
	}

	if err != nil {
		return LoginResp{}, err
	}
	defer resp.Body.Close()
	bits, _ := ioutil.ReadAll(resp.Body)
	var result LoginResp
	err = json.Unmarshal(bits, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
