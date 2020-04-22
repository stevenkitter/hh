package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	UserInfoURL = "http://32.33.1.123:8082/pkehr/*.jsonRequest"
)

type UserInfoReq struct {
	Cnd           []interface{} `json:"cnd"`
	Method        string        `json:"method"`
	PageNo        int           `json:"pageNo"`
	PageSize      int           `json:"pageSize"`
	Schema        string        `json:"schema"`
	ServiceAction string        `json:"serviceAction"`
	ServiceId     string        `json:"serviceId"`
}

type UserInfoListResp struct {
	TotalCount int            `json:"totalCount"`
	PageNo     int            `json:"pageNo"`
	PageSize   int            `json:"pageSize"`
	Code       int            `json:"code"`
	Msg        string         `json:"msg"`
	Body       []UserInfoItem `json:"body"`
}

type UserInfoDetail struct {
	Body struct {
		MiddleID           string `json:"middleId"`
		DiseasetextRedioCJ string `json:"diseasetextRedioCJ"`
		DiseasetextCheckZN struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"diseasetextCheckZN"`
		LastModifyUser string `json:"lastModifyUser"`
		BloodTypeCode  struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"bloodTypeCode"`
		PersonName string `json:"personName"`
		NationCode struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"nationCode"`
		WorkPlace   string      `json:"workPlace"`
		HomePlace   interface{} `json:"homePlace"`
		ShhjCheckCS struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"shhjCheckCS"`
		PhoneNumber        string `json:"phoneNumber"`
		ZipCode            string `json:"zipCode"`
		DiseasetextRadioGm string `json:"diseasetext_radio_gm"`
		ManaDoctorID       struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"manaDoctorId"`
		IsAgrRegister struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"isAgrRegister"`
		ShhjCheckRLLX struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"shhjCheckRLLX"`
		DiseasetextRedioXDJM string `json:"diseasetextRedioXDJM"`
		DiseasetextCheckCJ   struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"diseasetextCheckCJ"`
		DeadFlag struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"deadFlag"`
		EmpiID         string `json:"empiId"`
		DiseasetextSx0 string `json:"diseasetext_sx0"`
		MasterFlag     struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"masterFlag"`
		RegionCodeText     string `json:"regionCode_text"`
		DiseasetextCheckSs string `json:"diseasetext_check_ss"`
		ShhjCheckYS        struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"shhjCheckYS"`
		Status             string `json:"status"`
		Zlls               string `json:"zlls"`
		DiseasetextCheckGm struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"diseasetext_check_gm"`
		ShhjCheckCFPFSS struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"shhjCheckCFPFSS"`
		FamilyDoctorSigned struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"familyDoctorSigned"`
		WorkCode struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"workCode"`
		CreateUser    string `json:"createUser"`
		EducationCode struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"educationCode"`
		Email              interface{} `json:"email"`
		CreateUnit         string      `json:"createUnit"`
		PhrID              string      `json:"phrId"`
		DiseasetextRedioMQ string      `json:"diseasetextRedioMQ"`
		InsuranceType      interface{} `json:"insuranceType"`
		MaritalStatusCode  struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"maritalStatusCode"`
		DiseasetextCheckSx string `json:"diseasetext_check_sx"`
		VersionNumber      string `json:"versionNumber"`
		DiseasetextYCBS    string `json:"diseasetextYCBS"`
		IncomeSource       struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"incomeSource"`
		InsuranceCode struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"insuranceCode"`
		DiseasetextCheckWs string      `json:"diseasetext_check_ws"`
		StartdateWs1       interface{} `json:"startdate_ws1"`
		DiseasetextRedioFq string      `json:"diseasetext_redio_fq"`
		CreateTime         string      `json:"createTime"`
		Birthday           string      `json:"birthday"`
		DiseasetextSs0     string      `json:"diseasetext_ss0"`
		ManaUnitID         struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"manaUnitId"`
		LastModifyUnit                 string `json:"lastModifyUnit"`
		PersonalizedFamilyDoctorSigned struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"personalizedFamilyDoctorSigned"`
		DiseasetextCheckFq struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"diseasetext_check_fq"`
		StartdateSs0       interface{} `json:"startdate_ss0"`
		Contact            string      `json:"contact"`
		DiseasetextRadioBl string      `json:"diseasetext_radio_bl"`
		DiseasetextRadioJb struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"diseasetext_radio_jb"`
		SignFlag struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"signFlag"`
		DefinePhrid string `json:"definePhrid"`
		RhBloodCode struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"rhBloodCode"`
		NationalityCode     string `json:"nationalityCode"`
		RegisteredPermanent struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"registeredPermanent"`
		DiseasetextSs struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"diseasetext_ss"`
		DiseasetextSx struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"diseasetext_sx"`
		KnowFlag struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"knowFlag"`
		InsuranceText      interface{} `json:"insuranceText"`
		StartdateSx0       interface{} `json:"startdate_sx0"`
		DiseasetextCheckJb struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"diseasetext_check_jb"`
		InsuranceCode1 interface{} `json:"insuranceCode1"`
		SexCode        struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"sexCode"`
		RegionCode struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"regionCode"`
		ContactPhone         string `json:"contactPhone"`
		IDCard               string `json:"idCard"`
		DiseasetextCheckXDJM struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"diseasetextCheckXDJM"`
		DiseasetextRedioYCBS struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"diseasetextRedioYCBS"`
		DeadDate           interface{} `json:"deadDate"`
		DiseasetextRedioZN string      `json:"diseasetextRedioZN"`
		StartWorkDate      interface{} `json:"startWorkDate"`
		Photo              interface{} `json:"photo"`
		DiseasetextWs      struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"diseasetext_ws"`
		DeadReason         interface{} `json:"deadReason"`
		DiseasetextCheckMQ struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"diseasetextCheckMQ"`
		Address            string `json:"address"`
		DiseasetextWs1     string `json:"diseasetext_ws1"`
		DiseasetextCheckBl struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"diseasetext_check_bl"`
		MobileNumber string `json:"mobileNumber"`
		ShhjCheckQCL struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"shhjCheckQCL"`
		LastModifyTime       interface{} `json:"lastModifyTime"`
		IsPovertyAlleviation struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"isPovertyAlleviation"`
	} `json:"body"`
	Code int `json:"code"`
}

type UserInfoItem struct {
	Address                  string      `json:"address"`
	Birthday                 string      `json:"birthday"`
	BloodTypeCode            string      `json:"bloodTypeCode"`
	BloodTypeCodeText        string      `json:"bloodTypeCode_text"`
	CancellationDate         interface{} `json:"cancellationDate"`
	CancellationReason       interface{} `json:"cancellationReason"`
	CancellationReasonText   interface{} `json:"cancellationReason_text"`
	CancellationUnit         interface{} `json:"cancellationUnit"`
	CancellationUnitText     interface{} `json:"cancellationUnit_text"`
	CancellationUser         string      `json:"cancellationUser"`
	CancellationUserText     string      `json:"cancellationUser_text"`
	ContactPhone             string      `json:"contactPhone"`
	CreateDate               string      `json:"createDate"`
	CreateUnit               string      `json:"createUnit"`
	CreateUnitText           string      `json:"createUnit_text"`
	CreateUser               string      `json:"createUser"`
	CreateUserText           string      `json:"createUser_text"`
	DataSource               interface{} `json:"dataSource"`
	DataSourceText           interface{} `json:"dataSource_text"`
	DeadDate                 interface{} `json:"deadDate"`
	DeadFlag                 string      `json:"deadFlag"`
	DeadFlagText             string      `json:"deadFlag_text"`
	DeadReason               interface{} `json:"deadReason"`
	DefinePhrid              string      `json:"definePhrid"`
	EmpiID                   string      `json:"empiId"`
	ExistHealthCheck         interface{} `json:"existHealthCheck"`
	ExistHealthCheckText     interface{} `json:"existHealthCheck_text"`
	FamilyDoctorID           interface{} `json:"familyDoctorId"`
	FamilyDoctorIDText       interface{} `json:"familyDoctorId_text"`
	FamilyDoctorSigned       string      `json:"familyDoctorSigned"`
	FamilyDoctorSignedText   string      `json:"familyDoctorSigned_text"`
	FamilyID                 string      `json:"familyId"`
	FatherID                 interface{} `json:"fatherId"`
	FatherName               interface{} `json:"fatherName"`
	IDCard                   string      `json:"idCard"`
	IncomeSource             string      `json:"incomeSource"`
	IncomeSourceText         string      `json:"incomeSource_text"`
	IsAgrRegister            string      `json:"isAgrRegister"`
	IsAgrRegisterText        string      `json:"isAgrRegister_text"`
	IsDiabetes               string      `json:"isDiabetes"`
	IsDiabetesText           string      `json:"isDiabetes_text"`
	IsHypertension           string      `json:"isHypertension"`
	IsHypertensionText       string      `json:"isHypertension_text"`
	IsPovertyAlleviation     string      `json:"isPovertyAlleviation"`
	IsPovertyAlleviationText string      `json:"isPovertyAlleviation_text"`
	KnowFlag                 string      `json:"knowFlag"`
	KnowFlagText             string      `json:"knowFlag_text"`
	LastModifyDate           string      `json:"lastModifyDate"`
	LastModifyUnit           string      `json:"lastModifyUnit"`
	LastModifyUnitText       string      `json:"lastModifyUnit_text"`
	LastModifyUser           string      `json:"lastModifyUser"`
	LastModifyUserText       string      `json:"lastModifyUser_text"`
	ManaDoctorID             string      `json:"manaDoctorId"`
	ManaDoctorIDText         string      `json:"manaDoctorId_text"`
	ManaUnitID               string      `json:"manaUnitId"`
	ManaUnitIDText           string      `json:"manaUnitId_text"`
	MasterFlag               string      `json:"masterFlag"`
	MasterFlagText           string      `json:"masterFlag_text"`
	MobileNumber             string      `json:"mobileNumber"`
	MotherID                 interface{} `json:"motherId"`
	MotherName               interface{} `json:"motherName"`
	OldlastVisitDate         interface{} `json:"oldlastVisitDate"`
	PartnerID                interface{} `json:"partnerId"`
	PartnerName              interface{} `json:"partnerName"`
	PersonName               string      `json:"personName"`
	PhrID                    string      `json:"phrId"`
	RegionCode               string      `json:"regionCode"`
	RegionCodeText           string      `json:"regionCode_text"`
	RegisteredPermanent      string      `json:"registeredPermanent"`
	RegisteredPermanentText  string      `json:"registeredPermanent_text"`
	RelaCode                 string      `json:"relaCode"`
	RelaCodeText             string      `json:"relaCode_text"`
	SexCode                  string      `json:"sexCode"`
	SexCodeText              string      `json:"sexCode_text"`
	SignFlag                 string      `json:"signFlag"`
	SignFlagText             string      `json:"signFlag_text"`
	Status                   string      `json:"status"`
	StatusText               string      `json:"status_text"`
	Sxsj                     interface{} `json:"sxsj"`
	WorkCode                 string      `json:"workCode"`
	WorkCodeText             string      `json:"workCode_text"`
	Zlls                     string      `json:"zlls"`
	ZllsText                 string      `json:"zlls_text"`
	Zzsj                     interface{} `json:"zzsj"`
}

// GongweiUserInfo g
// {"cnd":["and",["eq",["$","a.status"],["s","0"]],["like",["$","regionCode"],["s","320111001010%"]]],"method":"execute","pageNo":1,"pageSize":25,"schema":"chis.application.hr.schemas.EHR_HealthRecord","serviceAction":"","serviceId":"chis.simpleQuery"}
func (s *Server) GongweiUserInfo(regionCode string, start int) (dest UserInfoListResp, err error) {
	//var start = 297
	req := UserInfoReq{
		Method:        "execute",
		PageNo:        start + 1,
		PageSize:      25,
		Schema:        "chis.application.hr.schemas.EHR_HealthRecord",
		ServiceAction: "",
		ServiceId:     "chis.simpleQuery",
		Cnd: []interface{}{"and",
			[]interface{}{"eq", []string{"$", "a.status"}, []string{"s", "0"}},
			[]interface{}{"like", []string{"$", "regionCode"}, []string{"s", regionCode}},
		},
	}
	b, _ := json.Marshal(&req)
	d, err := s.RequestUserInfoJson(b, start)
	if err != nil {
		return
	}
	dest = UserInfoListResp{}
	err = json.Unmarshal(d, &dest)
	if err != nil {
		return
	}
	//log.Printf("dest %v", dest)
	return
}

func (s *Server) RequestUserInfoJson(b []byte, start int) ([]byte, error) {
	//url := "http://32.33.1.123:8082/pkehr/*.jsonRequest"
	url := fmt.Sprintf(UserInfoURL+"?start=%d&limit=%d", start, 25)
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

type EMRegion struct {
	EmpiId     string `json:"empiId"`
	RegionCode string `json:"regionCode"`
}

type UserInfoDetailReq struct {
	Body          EMRegion `json:"body"`
	Method        string   `json:"method"`
	Op            string   `json:"op"`
	ServiceAction string   `json:"serviceAction"`
	ServiceId     string   `json:"serviceId"`
}

// empiId: 32011120090918601800000000000000
func (s *Server) UserInfoDetail(empiId, regionCode string) (u UserInfoDetail, err error) {
	req := UserInfoDetailReq{
		Body: EMRegion{
			EmpiId:     empiId,
			RegionCode: regionCode,
		},
		Method:        "execute",
		Op:            "update",
		ServiceAction: "LoadBasicPersonalInformation",
		ServiceId:     "chis.basicPersonalInformationService",
	}
	d, _ := json.Marshal(&req)
	data, err := s.RequestJson(d)
	if err != nil {
		return
	}
	u = UserInfoDetail{}
	err = json.Unmarshal(data, &u)
	if err != nil {
		return
	}
	return
}
