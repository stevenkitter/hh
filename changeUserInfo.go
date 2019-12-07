package main

import (
	"encoding/json"
	"reflect"
)

type ChangeUserInfoRequest struct {
	Body ChangeUserInfoRequestBody `json:"body"`
	Op   string                    `json:"op"`
	RequestJson
}

type ChangeUserInfoRequestBody struct {
	EmpiID                         string `json:"empiId"`
	PhrID                          string `json:"phrId"`
	MiddleID                       string `json:"middleId"`
	CardNo                         string `json:"cardNo"`
	DefinePhrid                    string `json:"definePhrid"`
	PersonName                     string `json:"personName"`
	SexCode                        string `json:"sexCode"`
	Birthday                       string `json:"birthday"`
	IDCard                         string `json:"idCard"`
	Address                        string `json:"address"`
	WorkPlace                      string `json:"workPlace"`
	MobileNumber                   string `json:"mobileNumber"`
	Contact                        string `json:"contact"`
	ContactPhone                   string `json:"contactPhone"`
	RegisteredPermanent            string `json:"registeredPermanent"`
	IsAgrRegister                  string `json:"isAgrRegister"`
	SignFlag                       string `json:"signFlag"`
	KnowFlag                       string `json:"knowFlag"`
	IsPovertyAlleviation           string `json:"isPovertyAlleviation"`
	IncomeSource                   string `json:"incomeSource"`
	NationCode                     string `json:"nationCode"`
	BloodTypeCode                  string `json:"bloodTypeCode"`
	RhBloodCode                    string `json:"rhBloodCode"`
	EducationCode                  string `json:"educationCode"`
	WorkCode                       string `json:"workCode"`
	MaritalStatusCode              string `json:"maritalStatusCode"`
	InsuranceCode                  string `json:"insuranceCode"`
	InsuranceCode1                 string `json:"insuranceCode1"`
	MasterFlag                     string `json:"masterFlag"`
	FamilyID                       string `json:"familyId"`
	DiseasetextCheckGm             string `json:"diseasetext_check_gm"`
	AQt1                           string `json:"a_qt1"`
	DiseasetextCheckBl             string `json:"diseasetext_check_bl"`
	DiseasetextRadioJb             string `json:"diseasetext_radio_jb"`
	DiseasetextCheckJb             string `json:"diseasetext_check_jb"`
	ConfirmdateGxy                 string `json:"confirmdate_gxy"`
	ConfirmdateGxb                 string `json:"confirmdate_gxb"`
	ConfirmdateExzl                string `json:"confirmdate_exzl"`
	ConfirmdateZxjsjb              string `json:"confirmdate_zxjsjb"`
	ConfirmdateGzjb                string `json:"confirmdate_gzjb"`
	ConfirmdateZyb                 string `json:"confirmdate_zyb"`
	ConfirmdatePx                  string `json:"confirmdate_px"`
	ConfirmdateQt                  string `json:"confirmdate_qt"`
	ConfirmdateTnb                 string `json:"confirmdate_tnb"`
	ConfirmdateMxzsxfjb            string `json:"confirmdate_mxzsxfjb"`
	ConfirmdateNzz                 string `json:"confirmdate_nzz"`
	ConfirmdateJhb                 string `json:"confirmdate_jhb"`
	ConfirmdateXtjx                string `json:"confirmdate_xtjx"`
	ConfirmdateSzjb                string `json:"confirmdate_szjb"`
	ConfirmdateQtfdcrb             string `json:"confirmdate_qtfdcrb"`
	DiseasetextZyb                 string `json:"diseasetext_zyb"`
	DiseasetextQtfdcrb             string `json:"diseasetext_qtfdcrb"`
	DiseasetextQt                  string `json:"diseasetext_qt"`
	DiseasetextSs                  string `json:"diseasetext_ss"`
	DiseasetextSs0                 string `json:"diseasetext_ss0"`
	StartdateSs0                   string `json:"startdate_ss0"`
	DiseasetextSs1                 string `json:"diseasetext_ss1"`
	StartdateSs1                   string `json:"startdate_ss1"`
	DiseasetextWs                  string `json:"diseasetext_ws"`
	DiseasetextWs0                 string `json:"diseasetext_ws0"`
	StartdateWs0                   string `json:"startdate_ws0"`
	DiseasetextWs1                 string `json:"diseasetext_ws1"`
	StartdateWs1                   string `json:"startdate_ws1"`
	DiseasetextSx                  string `json:"diseasetext_sx"`
	DiseasetextSx0                 string `json:"diseasetext_sx0"`
	StartdateSx0                   string `json:"startdate_sx0"`
	DiseasetextSx1                 string `json:"diseasetext_sx1"`
	StartdateSx1                   string `json:"startdate_sx1"`
	DiseasetextCheckFq             string `json:"diseasetext_check_fq"`
	QtFq1                          string `json:"qt_fq1"`
	DiseasetextCheckMQ             string `json:"diseasetextCheckMQ"`
	QtMq1                          string `json:"qt_mq1"`
	DiseasetextCheckXDJM           string `json:"diseasetextCheckXDJM"`
	QtXdjm1                        string `json:"qt_xdjm1"`
	DiseasetextCheckZN             string `json:"diseasetextCheckZN"`
	QtZn1                          string `json:"qt_zn1"`
	DiseasetextRedioYCBS           string `json:"diseasetextRedioYCBS"`
	DiseasetextYCBS                string `json:"diseasetextYCBS"`
	DiseasetextCheckCJ             string `json:"diseasetextCheckCJ"`
	CjqkQtcj1                      string `json:"cjqk_qtcj1"`
	ShhjCheckCFPFSS                string `json:"shhjCheckCFPFSS"`
	ShhjCheckRLLX                  string `json:"shhjCheckRLLX"`
	ShhjCheckYS                    string `json:"shhjCheckYS"`
	ShhjCheckCS                    string `json:"shhjCheckCS"`
	ShhjCheckQCL                   string `json:"shhjCheckQCL"`
	DeadFlag                       string `json:"deadFlag"`
	DeadDate                       string `json:"deadDate"`
	DeadReason                     string `json:"deadReason"`
	PersonalizedFamilyDoctorSigned string `json:"personalizedFamilyDoctorSigned"`
	FamilyDoctorSigned             string `json:"familyDoctorSigned"`
	RegionCode                     string `json:"regionCode"`
	ManaDoctorID                   string `json:"manaDoctorId"`
	ManaUnitID                     string `json:"manaUnitId"`
}

/**
 *  根据获取的详情，表里的信息修改请求
 */
func NewChangeUserInfoRequest(user ChangeUserInfoRequestBody, detail PersonalInfoRespBody) ChangeUserInfoRequest {
	b := ChangeUserInfoRequestBody{}
	t := reflect.TypeOf(b)
	p := reflect.ValueOf(&b).Elem()

	for k := 0; k < t.NumField(); k++ {
		fName := t.Field(k).Name
		dataV := GetStringValueFromStructByName(user, fName)
		p.FieldByName(fName).Set(reflect.ValueOf(dataV))
	}
	b.EmpiID = detail.EmpiID
	b.PhrID = detail.PhrID
	b.MiddleID = detail.MiddleID

	op := "update"
	if detail.RegionCode.Key != "" {
		b.RegionCode = detail.RegionCode.Key
	}
	if detail.ManaDoctorID.Key != "" {
		b.ManaDoctorID = detail.ManaDoctorID.Key
	}
	if detail.ManaUnitID.Key != "" {
		b.ManaUnitID = detail.ManaUnitID.Key
	}
	if detail.EmpiID == "" {
		op = "create"
	}

	return ChangeUserInfoRequest{
		Body: b,
		Op:   op,
		RequestJson: RequestJson{
			Method:        "execute",
			ServiceAction: "saveBasicPersonalInformation",
			ServiceId:     "chis.basicPersonalInformationService",
		},
	}
}

type ChangeUserInfoResp struct {
	Body ChangeUserInfoRespBody `json:"body"`
	Code int                    `json:"code"`
}
type ChangeUserInfoRespBody struct {
	MiddleID   string    `json:"middleId"`
	EmpiID     string    `json:"empiId"`
	RegionCode Insurance `json:"regionCode"`
	ManaUnitID Insurance `json:"manaUnitId"`
	PhrID      string    `json:"phrId"`
}

func (s *Server) ChangeUserInfoRequest(user ChangeUserInfoRequestBody, detail PersonalInfoRespBody) (ChangeUserInfoResp, error) {
	r := NewChangeUserInfoRequest(user, detail)
	b, err := json.Marshal(&r)
	if err != nil {
		return ChangeUserInfoResp{}, err
	}
	bs, err := s.RequestJson(b)
	if err != nil {
		return ChangeUserInfoResp{}, err
	}
	var result ChangeUserInfoResp
	err = json.Unmarshal(bs, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
