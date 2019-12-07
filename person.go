package main

import "encoding/json"

type PersonInfoRequest struct {
	RequestJson
	Body struct {
		IdCard string `json:"idCard"`
	} `json:"body"`
	Op string `json:"op"`
}

func NewPersonInfoRequest(idCard string) PersonInfoRequest {
	return PersonInfoRequest{
		RequestJson: RequestJson{
			Method:        "execute",
			Schema:        "",
			ServiceAction: "LoadBasicPersonalInformation",
			ServiceId:     "chis.basicPersonalInformationService",
		},
		Body: struct {
			IdCard string `json:"idCard"`
		}{IdCard: idCard},
		Op: "create",
	}
}

type PersonalInfoResp struct {
	Body PersonalInfoRespBody `json:"body"`
	Code int                  `json:"code"`
}

type PersonalInfoRespBody struct {
	MiddleID                       string    `json:"middleId"`
	DiseasetextRedioCJ             string    `json:"diseasetextRedioCJ"`
	DiseasetextCheckZN             Insurance `json:"diseasetextCheckZN"`
	LastModifyUser                 string    `json:"lastModifyUser"`
	BloodTypeCode                  Insurance `json:"bloodTypeCode"`
	PersonName                     string    `json:"personName"`
	NationCode                     Insurance `json:"nationCode"`
	WorkPlace                      string    `json:"workPlace"`
	HomePlace                      Insurance `json:"homePlace"`
	ShhjCheckCS                    Insurance `json:"shhjCheckCS"`
	PhoneNumber                    string    `json:"phoneNumber"`
	ZipCode                        string    `json:"zipCode"`
	DiseasetextRadioGm             string    `json:"diseasetext_radio_gm"`
	ManaDoctorID                   Insurance `json:"manaDoctorId"`
	IsAgrRegister                  Insurance `json:"isAgrRegister"`
	ShhjCheckRLLX                  Insurance `json:"shhjCheckRLLX"`
	DiseasetextRedioXDJM           string    `json:"diseasetextRedioXDJM"`
	DiseasetextCheckCJ             Insurance `json:"diseasetextCheckCJ"`
	DeadFlag                       Insurance `json:"deadFlag"`
	EmpiID                         string    `json:"empiId"`
	DiseasetextSx0                 string    `json:"diseasetext_sx0"`
	MasterFlag                     Insurance `json:"masterFlag"`
	RegionCodeText                 string    `json:"regionCode_text"`
	DiseasetextCheckSs             string    `json:"diseasetext_check_ss"`
	ShhjCheckYS                    Insurance `json:"shhjCheckYS"`
	Status                         string    `json:"status"`
	Zlls                           string    `json:"zlls"`
	DiseasetextCheckGm             Insurance `json:"diseasetext_check_gm"`
	ShhjCheckCFPFSS                Insurance `json:"shhjCheckCFPFSS"`
	FamilyDoctorSigned             Insurance `json:"familyDoctorSigned"`
	WorkCode                       Insurance `json:"workCode"`
	CreateUser                     string    `json:"createUser"`
	EducationCode                  Insurance `json:"educationCode"`
	Email                          Insurance `json:"email"`
	CreateUnit                     string    `json:"createUnit"`
	PhrID                          string    `json:"phrId"`
	DiseasetextRedioMQ             string    `json:"diseasetextRedioMQ"`
	InsuranceType                  string    `json:"insuranceType"`
	MaritalStatusCode              Insurance `json:"maritalStatusCode"`
	DiseasetextCheckSx             string    `json:"diseasetext_check_sx"`
	VersionNumber                  string    `json:"versionNumber"`
	DiseasetextYCBS                string    `json:"diseasetextYCBS"`
	IncomeSource                   Insurance `json:"incomeSource"`
	InsuranceCode                  Insurance `json:"insuranceCode"`
	DiseasetextCheckWs             string    `json:"diseasetext_check_ws"`
	StartdateWs1                   Insurance `json:"startdate_ws1"`
	DiseasetextRedioFq             string    `json:"diseasetext_redio_fq"`
	CreateTime                     string    `json:"createTime"`
	Birthday                       string    `json:"birthday"`
	DiseasetextSs0                 string    `json:"diseasetext_ss0"`
	ManaUnitID                     Insurance `json:"manaUnitId"`
	LastModifyUnit                 string    `json:"lastModifyUnit"`
	PersonalizedFamilyDoctorSigned Insurance `json:"personalizedFamilyDoctorSigned"`
	DiseasetextCheckFq             Insurance `json:"diseasetext_check_fq"`
	StartdateSs0                   Insurance `json:"startdate_ss0"`
	Contact                        string    `json:"contact"`
	DiseasetextRadioBl             string    `json:"diseasetext_radio_bl"`
	DiseasetextRadioJb             Insurance `json:"diseasetext_radio_jb"`
	SignFlag                       Insurance `json:"signFlag"`
	DefinePhrid                    string    `json:"definePhrid"`
	RhBloodCode                    Insurance `json:"rhBloodCode"`
	NationalityCode                string    `json:"nationalityCode"`
	RegisteredPermanent            Insurance `json:"registeredPermanent"`
	DiseasetextSs                  Insurance `json:"diseasetext_ss"`
	DiseasetextSx                  Insurance `json:"diseasetext_sx"`
	KnowFlag                       Insurance `json:"knowFlag"`
	InsuranceText                  Insurance `json:"insuranceText"`
	StartdateSx0                   Insurance `json:"startdate_sx0"`
	DiseasetextCheckJb             Insurance `json:"diseasetext_check_jb"`
	InsuranceCode1                 Insurance `json:"insuranceCode1"`
	SexCode                        Insurance `json:"sexCode"`
	RegionCode                     Insurance `json:"regionCode"`
	ContactPhone                   string    `json:"contactPhone"`
	IDCard                         string    `json:"idCard"`
	DiseasetextCheckXDJM           Insurance `json:"diseasetextCheckXDJM"`
	DiseasetextRedioYCBS           Insurance `json:"diseasetextRedioYCBS"`
	DeadDate                       Insurance `json:"deadDate"`
	DiseasetextRedioZN             string    `json:"diseasetextRedioZN"`
	StartWorkDate                  Insurance `json:"startWorkDate"`
	Photo                          string    `json:"photo"`
	DiseasetextWs                  Insurance `json:"diseasetext_ws"`
	DeadReason                     Insurance `json:"deadReason"`
	DiseasetextCheckMQ             Insurance `json:"diseasetextCheckMQ"`
	Address                        string    `json:"address"`
	DiseasetextWs1                 string    `json:"diseasetext_ws1"`
	DiseasetextCheckBl             Insurance `json:"diseasetext_check_bl"`
	MobileNumber                   string    `json:"mobileNumber"`
	ShhjCheckQCL                   Insurance `json:"shhjCheckQCL"`
	LastModifyTime                 string    `json:"lastModifyTime"`
	IsPovertyAlleviation           Insurance `json:"isPovertyAlleviation"`
}

/**
 *  查询用户是否有个人信息
 */
func (s *Server) PersonalInfo(idCard string) (PersonalInfoResp, error) {
	r := NewPersonInfoRequest(idCard)
	b, err := json.Marshal(&r)
	if err != nil {
		return PersonalInfoResp{}, err
	}
	bs, err := s.RequestJson(b)
	if err != nil {
		return PersonalInfoResp{}, err
	}
	var result PersonalInfoResp
	err = json.Unmarshal(bs, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
