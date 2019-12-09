package main

import (
	"bytes"
	"encoding/json"
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

func NewUserInfoRequestJson(pkey string) UserInfoRequestJson {
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

type UserNewInfoRequestJson struct {
	Body          UserNewInfoRequestJsonBody `json:"body"`
	Method        string                     `json:"method"`
	Schema        string                     `json:"schema"`
	ServiceAction string                     `json:"serviceAction"`
	ServiceId     string                     `json:"serviceId"`
}
type UserNewInfoRequestJsonBody struct {
	IdCard     string `json:"idCard"`
	PersonName string `json:"personName"`
	QueryBy    string `json:"queryBy"`
}

func NewUserNewInfoRequestJson(idCard string) UserNewInfoRequestJson {
	return UserNewInfoRequestJson{
		Body: UserNewInfoRequestJsonBody{
			IdCard:  idCard,
			QueryBy: "idCard",
		},
		Method:        "execute",
		Schema:        "chis.application.mpi.schemas.MPI_DemographicInfo",
		ServiceAction: "advancedSearch",
		ServiceId:     "chis.empiService",
	}
}

type UserNewInfoResp struct {
	Body     []*UserNewInfoBody `json:"body"`
	PageNo   int                `json:"pageNo"`
	PageSize int                `json:"pageSize"`
	Code     int                `json:"code"`
}

type UserNewInfoBody struct {
	CreateUserText          string      `json:"createUser_text"`
	InsuranceCodeText       string      `json:"insuranceCode_text"`
	Score                   string      `json:"score"`
	LastModifyUser          string      `json:"lastModifyUser"`
	BloodTypeCode           string      `json:"bloodTypeCode"`
	PersonName              string      `json:"personName"`
	NationCode              string      `json:"nationCode"`
	WorkPlace               string      `json:"workPlace"`
	BloodTypeCodeText       string      `json:"bloodTypeCode_text"`
	HomePlace               interface{} `json:"homePlace"`
	EducationCodeText       string      `json:"educationCode_text"`
	RegisteredPermanentText string      `json:"registeredPermanent_text"`
	PhoneNumber             interface{} `json:"phoneNumber"`
	ZipCode                 interface{} `json:"zipCode"`
	Cards                   []Card      `json:"cards"`
	EmpiID                  string      `json:"empiId"`
	Status                  string      `json:"status"`
	Zlls                    interface{} `json:"zlls"`
	WorkCode                string      `json:"workCode"`
	EducationCode           string      `json:"educationCode"`
	CreateUser              string      `json:"createUser"`
	Email                   interface{} `json:"email"`
	CreateUnit              string      `json:"createUnit"`
	InsuranceType           interface{} `json:"insuranceType"`
	MaritalStatusCode       string      `json:"maritalStatusCode"`
	VersionNumber           string      `json:"versionNumber"`
	InsuranceCode           string      `json:"insuranceCode"`
	Birthday                string      `json:"birthday"`
	CreateTime              string      `json:"createTime"`
	SexCodeText             string      `json:"sexCode_text"`
	LastModifyUnit          string      `json:"lastModifyUnit"`
	Contact                 string      `json:"contact"`
	CreateUnitText          string      `json:"createUnit_text"`
	ZllsText                interface{} `json:"zlls_text"`
	DefinePhrid             string      `json:"definePhrid"`
	RhBloodCode             string      `json:"rhBloodCode"`
	RegisteredPermanent     string      `json:"registeredPermanent"`
	NationalityCode         string      `json:"nationalityCode"`
	MaritalStatusCodeText   string      `json:"maritalStatusCode_text"`
	NationalityCodeText     string      `json:"nationalityCode_text"`
	WorkCodeText            string      `json:"workCode_text"`
	InsuranceText           interface{} `json:"insuranceText"`
	SexCode                 string      `json:"sexCode"`
	IDCard                  string      `json:"idCard"`
	ContactPhone            string      `json:"contactPhone"`
	LastModifyUnitText      string      `json:"lastModifyUnit_text"`
	NationCodeText          string      `json:"nationCode_text"`
	Photo                   string      `json:"photo"`
	StartWorkDate           interface{} `json:"startWorkDate"`
	Address                 string      `json:"address"`
	LastModifyUserText      string      `json:"lastModifyUser_text"`
	RhBloodCodeText         string      `json:"rhBloodCode_text"`
	MobileNumber            string      `json:"mobileNumber"`
	LastModifyTime          string      `json:"lastModifyTime"`
}

func (s *Server) RequestUserNewInfo(idcard string) (UserNewInfoResp, error) {
	url := "http://32.33.1.123:8082/pkehr/*.jsonRequest"
	cli := http.Client{}
	reqData := NewUserNewInfoRequestJson(idcard)
	bit, _ := json.Marshal(&reqData)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bit))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Cookie", s.Cookie)
	resp, err := cli.Do(request)
	if err != nil {
		return UserNewInfoResp{}, err
	}
	defer resp.Body.Close()
	bits, _ := ioutil.ReadAll(resp.Body)
	var result UserNewInfoResp
	err = json.Unmarshal(bits, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
