package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserInfoResp struct {
	Body UserHealthInfo `json:"body"`
	Code int            `json:"code"`
}

type UserHealthInfo struct {
	InsuranceCode       Insurance `json:"insuranceCode"`
	Birthday            string    `json:"birthday"`
	CreateTime          string    `json:"createTime"`
	LastModifyUser      Insurance `json:"lastModifyUser"`
	Contact             string    `json:"contact"`
	BloodTypeCode       Insurance `json:"bloodTypeCode"`
	PersonName          string    `json:"personName"`
	NationCode          Insurance `json:"nationCode"`
	WorkPlace           string    `json:"workPlace"`
	HomePlace           string    `json:"homePlace"`
	RhBloodCode         Insurance `json:"rhBloodCode"`
	DefinePhrid         string    `json:"definePhrid"`
	PhoneNumber         string    `json:"phoneNumber"`
	RegisteredPermanent Insurance `json:"registeredPermanent"`
	NationalityCode     Insurance `json:"nationalityCode"`
	ZipCode             string    `json:"zipCode"`
	Cards               []Card    `json:"cards"`
	InsuranceText       string    `json:"insuranceText"`
	SexCode             Insurance `json:"sexCode"`
	EmpiId              string    `json:"empiId"`
	IdCard              string    `json:"idCard"`
	ContactPhone        string    `json:"contactPhone"`
	Status              string    `json:"status"`
	Zlls                Insurance `json:"zlls"`
	Photo               string    `json:"photo"`
	StartWorkDate       string    `json:"startWorkDate"`
	WorkCode            Insurance `json:"workCode"`
	EducationCode       Insurance `json:"educationCode"`
	CreateUser          Insurance `json:"createUser"`
	Address             string    `json:"address"`
	Email               string    `json:"email"`
	CreateUnit          Insurance `json:"createUnit"`
	MobileNumber        string    `json:"mobileNumber"`
	InsuranceType       string    `json:"insuranceType"`
	MaritalStatusCode   Insurance `json:"maritalStatusCode"`
	VersionNumber       string    `json:"versionNumber"`
	LastModifyTime      string    `json:"lastModifyTime"`
}

type Card struct {
	CardID           string      `json:"cardId"`
	CreateTime       interface{} `json:"createTime"`
	EmpiID           string      `json:"empiId"`
	Status           string      `json:"status"`
	LastModifyUnit   interface{} `json:"lastModifyUnit"`
	LastModifyUser   interface{} `json:"lastModifyUser"`
	CardTypeCodeText string      `json:"cardTypeCode_text"`
	CardTypeCode     string      `json:"cardTypeCode"`
	StatusText       string      `json:"status_text"`
	CardNo           string      `json:"cardNo"`
	CreateUser       interface{} `json:"createUser"`
	CreateUnit       interface{} `json:"createUnit"`
	ValidTime        interface{} `json:"validTime"`
	LastModifyTime   interface{} `json:"lastModifyTime"`
}

type Insurance struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type UserInfoRequestJson struct {
	Action        string   `json:"action"`
	Body          BodyJson `json:"body"`
	Method        string   `json:"method"`
	Pkey          string   `json:"pkey"`
	Schema        string   `json:"schema"`
	ServiceAction string   `json:"serviceAction"`
	ServiceId     string   `json:"serviceId"`
}

type BodyJson struct {
	PKey string `json:"pkey"`
}

func NewUserInfoRequestJson(idCard string) UserInfoRequestJson {
	pkey := fmt.Sprintf("%s00000000000000", idCard)
	return UserInfoRequestJson{
		Action: "create",
		Body: BodyJson{
			PKey: pkey,
		},
		Method:        "execute",
		Pkey:          pkey,
		Schema:        "chis.application.mpi.schemas.MPI_DemographicInfo",
		ServiceAction: "getDemographicInfo",
		ServiceId:     "chis.empiService",
	}
}

func (s *Server) RequestUserInfo(idcard string) (UserInfoResp, error) {
	url := "http://32.33.1.123:8082/pkehr/*.jsonRequest"
	cli := http.Client{}
	reqData := NewUserInfoRequestJson(idcard)
	bit, _ := json.Marshal(&reqData)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bit))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Cookie", s.Cookie)
	resp, err := cli.Do(request)
	if err != nil {
		return UserInfoResp{}, err
	}
	defer resp.Body.Close()
	bits, _ := ioutil.ReadAll(resp.Body)
	var result UserInfoResp
	err = json.Unmarshal(bits, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
