package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type ChangeHealthCheckRequest struct {
	RequestJson
	Op   string                       `json:"op"`
	Body ChangeHealthCheckRequestBody `json:"body"`
}

type ChangeHealthCheckRequestBody struct {
	IdCard             string
	HcData             HcRequestData        `json:"hcData"`
	LsData             LsRequestData        `json:"lsData"`
	ExaData            ExaRequestData       `json:"exaData"`
	AeData             AeRequestData        `json:"aeData"`
	HaData             HaRequestData        `json:"haData"`
	InhospitalListData []InhospitalListData `json:"inhospitalListData"`
	MedicineListData   []MedicineListData   `json:"medicineListData"`
	NiListData         []NiListData         `json:"niListData"`
}

type InhospitalListData struct {
	Type                     int    `json:"type"`
	SituationId              string `json:"situationId"`
	InhospitalDate           string `json:"inhospitalDate"`
	OuthospitalDate          string `json:"outhospitalDate"`
	InhospitalReason         string `json:"inhospitalReason"`
	MedicalestablishmentName string `json:"medicalestablishmentName"`
	MedicalrecordNumber      string `json:"medicalrecordNumber"`
}
type MedicineListData struct {
	SituationId   string `json:"situationId"`
	Medicine      string `json:"medicine"`
	Use           string `json:"use"`
	EachDose      string `json:"eachDose"`
	UseDate       string `json:"useDate"`
	MedicineYield string `json:"medicineYield"`
}
type NiListData struct {
	RecordId        string `json:"recordId"`
	Name            string `json:"name"`
	InoculationDate string `json:"inoculationDate"`
	InoculationUnit string `json:"inoculationUnit"`
}

func NewChangeHealthCheckRequest(body HealthCheckDetailBody, user ChangeHealthCheckRequestBody) ChangeHealthCheckRequest {
	op := "update"
	if body.HcData.HealthCheck == "" {
		op = "create"
	}
	return ChangeHealthCheckRequest{
		RequestJson: RequestJson{
			Method:        "execute",
			Schema:        "chis.application.hc.schemas.HC_HealthCheck",
			ServiceAction: "saveHealthCheckHtml",
			ServiceId:     "chis.healthCheckService",
		},
		Op: op,
		Body: ChangeHealthCheckRequestBody{
			HcData:  NewHcRequestData(body.HcData, user),
			LsData:  NewLsRequestData(body.LsData, user),
			ExaData: NewExaRequestData(body.ExaData),
			AeData:  NewAeRequestData(body.AeData, user),
			HaData:  NewHaRequestData(body.HaData),
			InhospitalListData: []InhospitalListData{NewInhospitalListRequestData(1),
				NewInhospitalListRequestData(1),
				NewInhospitalListRequestData(2),
				NewInhospitalListRequestData(2)},
			MedicineListData: []MedicineListData{
				NewMedicineListRequestData(),
				NewMedicineListRequestData(),
				NewMedicineListRequestData(),
				NewMedicineListRequestData(),
				NewMedicineListRequestData(),
				NewMedicineListRequestData(),
			},
			NiListData: []NiListData{
				NewNiListRequestData(),
				NewNiListRequestData(),
				NewNiListRequestData(),
			},
		},
	}
}

/**
 *  FUCK
 */
type HcRequestData struct {
	HealthCheck                  string `json:"healthCheck"`
	EmpiID                       string `json:"empiId"`
	PhrID                        string `json:"phrId"`
	CheckDate                    string `json:"checkDate"`
	CheckWay                     string `json:"checkWay"`
	Sf                           string `json:"sf"`
	Symptom                      string `json:"symptom"`
	SymptomOt                    string `json:"symptomOt"`
	Temperature                  string `json:"temperature"`
	Breathe                      string `json:"breathe"`
	Pulse                        string `json:"pulse"`
	ConstrictionL                string `json:"constriction_L"`
	DiastolicL                   string `json:"diastolic_L"`
	Constriction                 string `json:"constriction"`
	Diastolic                    string `json:"diastolic"`
	Fbs                          string `json:"fbs"`
	Alt                          string `json:"alt"`
	Ast                          string `json:"ast"`
	Alb                          string `json:"alb"`
	Tbil                         string `json:"tbil"`
	Dbil                         string `json:"dbil"`
	Height                       string `json:"height"`
	Weight                       string `json:"weight"`
	Waistline                    string `json:"waistline"`
	Bmi                          string `json:"bmi"`
	HealthStatus                 string `json:"healthStatus"`
	SelfCare                     string `json:"selfCare"`
	Cognitive                    string `json:"cognitive"`
	CognitiveZf                  string `json:"cognitiveZf"`
	Emotion                      string `json:"emotion"`
	EmotionZf                    string `json:"emotionZf"`
	CerebrovascularDiseases      string `json:"cerebrovascularDiseases"`
	OthercerebrovascularDiseases string `json:"othercerebrovascularDiseases"`
	HeartDisease                 string `json:"heartDisease"`
	OtherheartDisease            string `json:"otherheartDisease"`
	KidneyDiseases               string `json:"kidneyDiseases"`
	OtherkidneyDiseases          string `json:"otherkidneyDiseases"`
	VascularDisease              string `json:"VascularDisease"`
	OtherVascularDisease         string `json:"otherVascularDisease"`
	EyeDiseases                  string `json:"eyeDiseases"`
	OthereyeDiseases             string `json:"othereyeDiseases"`
	NeurologicalDiseases         string `json:"neurologicalDiseases"`
	NeurologicalDiseasesDesc     string `json:"neurologicalDiseasesDesc"`
	OtherDiseasesone             string `json:"otherDiseasesone"`
	OtherDiseasesoneDesc         string `json:"otherDiseasesoneDesc"`
	ManaDoctorID                 string `json:"manaDoctorId"`
	InhospitalFlag               string `json:"inhospitalFlag"`
	InfamilybedFlag              string `json:"infamilybedFlag"`
	MedicineFlag                 string `json:"medicineFlag"`
	NonimmuneFlag                string `json:"nonimmuneFlag"`
	ManaUnitID                   string `json:"manaUnitId"`
	CreateUser                   string `json:"createUser"`
	CreateUnit                   string `json:"createUnit"`
	CreateDate                   string `json:"createDate"`
	LastModifyUnit               string `json:"lastModifyUnit"`
	LastModifyUser               string `json:"lastModifyUser"`
	LastModifyDate               string `json:"lastModifyDate"`
}
type LsRequestData struct {
	HealthCheck               string `json:"healthCheck"`
	PhysicalExerciseFrequency string `json:"physicalExerciseFrequency"`
	EveryPhysicalExerciseTime string `json:"everyPhysicalExerciseTime"`
	Insistexercisetime        string `json:"insistexercisetime"`
	ExerciseStyle             string `json:"exerciseStyle"`
	DietaryHabit              string `json:"dietaryHabit"`
	WehtherSmoke              string `json:"wehtherSmoke"`
	BeginSmokeTime            string `json:"beginSmokeTime"`
	StopSmokeTime             string `json:"stopSmokeTime"`
	Smokes                    string `json:"smokes"`
	DrinkingFrequency         string `json:"drinkingFrequency"`
	AlcoholConsumption        string `json:"alcoholConsumption"`
	WhetherDrink              string `json:"whetherDrink"`
	StopDrinkingTime          string `json:"stopDrinkingTime"`
	GeginToDrinkTime          string `json:"geginToDrinkTime"`
	IsDrink                   string `json:"isDrink"`
	MainDrinkingVvarieties    string `json:"mainDrinkingVvarieties"`
	DrinkOther                string `json:"drinkOther"`
	Occupational              string `json:"occupational"`
	Jobs                      string `json:"jobs"`
	WorkTime                  string `json:"workTime"`
	Dust                      string `json:"dust"`
	DustPro                   string `json:"dustPro"`
	DustProDesc               string `json:"dustProDesc"`
	Ray                       string `json:"ray"`
	RayPro                    string `json:"rayPro"`
	RayProDesc                string `json:"rayProDesc"`
	PhysicalFactor            string `json:"physicalFactor"`
	PhysicalFactorPro         string `json:"physicalFactorPro"`
	PhysicalFactorProDesc     string `json:"physicalFactorProDesc"`
	Chemicals                 string `json:"chemicals"`
	ChemicalsPro              string `json:"chemicalsPro"`
	ChemicalsProDesc          string `json:"chemicalsProDesc"`
	Other                     string `json:"other"`
	OtherPro                  string `json:"otherPro"`
	OtherProDesc              string `json:"otherProDesc"`
	CreateUser                string `json:"createUser"`
	CreateUnit                string `json:"createUnit"`
	CreateDate                string `json:"createDate"`
	LastModifyUser            string `json:"lastModifyUser"`
	LastModifyUnit            string `json:"lastModifyUnit"`
	LastModifyDate            string `json:"lastModifyDate"`
}
type ExaRequestData struct {
	HealthCheck       string `json:"healthCheck"`
	Fundus            string `json:"fundus"`
	FundusDesc        string `json:"fundusDesc"`
	Skin              string `json:"skin"`
	SkinDesc          string `json:"skinDesc"`
	Sclera            string `json:"sclera"`
	ScleraDesc        string `json:"scleraDesc"`
	Lymphnodes        string `json:"lymphnodes"`
	LymphnodesDesc    string `json:"lymphnodesDesc"`
	BarrelChest       string `json:"barrelChest"`
	BreathSound       string `json:"breathSound"`
	BreathSoundDesc   string `json:"breathSoundDesc"`
	Rales             string `json:"rales"`
	RalesDesc         string `json:"ralesDesc"`
	HeartRate         string `json:"heartRate"`
	Rhythm            string `json:"rhythm"`
	HeartMurmur       string `json:"heartMurmur"`
	HeartMurmurDesc   string `json:"heartMurmurDesc"`
	AbdominAltend     string `json:"abdominAltend"`
	AbdominAltendDesc string `json:"abdominAltendDesc"`
	AdbominAlmass     string `json:"adbominAlmass"`
	AdbominAlmassDesc string `json:"adbominAlmassDesc"`
	LiverBig          string `json:"liverBig"`
	LiverBigDesc      string `json:"liverBigDesc"`
	Splenomegaly      string `json:"splenomegaly"`
	SplenomegalyDesc  string `json:"splenomegalyDesc"`
	Dullness          string `json:"dullness"`
	DullnessDesc      string `json:"dullnessDesc"`
	Edema             string `json:"edema"`
	FootPulse         string `json:"footPulse"`
	Dre               string `json:"dre"`
	DreDesc           string `json:"dreDesc"`
	Breast            string `json:"breast"`
	BreastDesc        string `json:"breastDesc"`
	Vulva             string `json:"vulva"`
	VulvaDesc         string `json:"vulvaDesc"`
	Vaginal           string `json:"vaginal"`
	VaginalDesc       string `json:"vaginalDesc"`
	Cervix            string `json:"cervix"`
	CervixDesc        string `json:"cervixDesc"`
	Palace            string `json:"palace"`
	PalaceDesc        string `json:"palaceDesc"`
	Attachment        string `json:"attachment"`
	AttachmentDesc    string `json:"attachmentDesc"`
	Tjother           string `json:"tjother"`
	CreateUser        string `json:"createUser"`
	CreateUnit        string `json:"createUnit"`
	CreateDate        string `json:"createDate"`
	LastModifyUser    string `json:"lastModifyUser"`
	LastModifyUnit    string `json:"lastModifyUnit"`
	LastModifyDate    string `json:"lastModifyDate"`
}
type AeRequestData struct {
	HealthCheck    string `json:"healthCheck"`
	Lip            string `json:"lip"`
	Denture        string `json:"denture"`
	LeftUp         string `json:"leftUp"`
	LeftDown       string `json:"leftDown"`
	RightUp        string `json:"rightUp"`
	RightDown      string `json:"rightDown"`
	Pharyngeal     string `json:"pharyngeal"`
	LeftEye        string `json:"leftEye"`
	RightEye       string `json:"rightEye"`
	RecLeftEye     string `json:"recLeftEye"`
	RecRightEye    string `json:"recRightEye"`
	Hearing        string `json:"hearing"`
	Motion         string `json:"motion"`
	Hgb            string `json:"hgb"`
	Wbc            string `json:"wbc"`
	Platelet       string `json:"platelet"`
	BloodOther     string `json:"bloodOther"`
	Proteinuria    string `json:"proteinuria"`
	Glu            string `json:"glu"`
	Dka            string `json:"dka"`
	Oc             string `json:"oc"`
	UrineOther     string `json:"urineOther"`
	Fbs            string `json:"fbs"`
	Fbs2           string `json:"fbs2"`
	Ecg            string `json:"ecg"`
	EcgText        string `json:"ecgText"`
	Malb           string `json:"malb"`
	Fob            string `json:"fob"`
	Hba1C          string `json:"hba1c"`
	Hbsag          string `json:"hbsag"`
	Alt            string `json:"alt"`
	Ast            string `json:"ast"`
	Alb            string `json:"alb"`
	Tbil           string `json:"tbil"`
	Dbil           string `json:"dbil"`
	Cr             string `json:"cr"`
	Bun            string `json:"bun"`
	Kalemia        string `json:"kalemia"`
	Natremia       string `json:"natremia"`
	Tc             string `json:"tc"`
	Tg             string `json:"tg"`
	Ldl            string `json:"ldl"`
	Hdl            string `json:"hdl"`
	X              string `json:"x"`
	XText          string `json:"xText"`
	B              string `json:"b"`
	BText          string `json:"bText"`
	Ps             string `json:"ps"`
	PsText         string `json:"psText"`
	FuOther        string `json:"fuOther"`
	CreateUser     string `json:"createUser"`
	CreateUnit     string `json:"createUnit"`
	CreateDate     string `json:"createDate"`
	LastModifyUser string `json:"lastModifyUser"`
	LastModifyUnit string `json:"lastModifyUnit"`
	LastModifyDate string `json:"lastModifyDate"`
}
type HaRequestData struct {
	HealthCheck        string `json:"healthCheck"`
	Recognize          string `json:"recognize"`
	Abnormality        string `json:"abnormality"`
	Abnormality1       string `json:"abnormality1"`
	Abnormality2       string `json:"abnormality2"`
	Abnormality3       string `json:"abnormality3"`
	Abnormality4       string `json:"abnormality4"`
	Mana               string `json:"mana"`
	RiskfactorsControl string `json:"riskfactorsControl"`
	TargetWeight       string `json:"targetWeight"`
	Vaccine            string `json:"vaccine"`
	PjOther            string `json:"pjOther"`
	CreateUser         string `json:"createUser"`
	CreateUnit         string `json:"createUnit"`
	CreateDate         string `json:"createDate"`
	LastModifyUser     string `json:"lastModifyUser"`
	LastModifyUnit     string `json:"lastModifyUnit"`
	LastModifyDate     string `json:"lastModifyDate"`
}

func Bmi(height, weight string) string {
	w, _ := strconv.ParseFloat(weight, 10)
	h, _ := strconv.ParseFloat(height, 10)
	r := w * 100 * 100 / (h * h)
	return fmt.Sprintf("%0.2f", r)
}
func NewHcRequestData(data HealthCheckDetailHcData, user ChangeHealthCheckRequestBody) HcRequestData {
	hc := HcRequestData{}
	t := reflect.TypeOf(hc)
	p := reflect.ValueOf(&hc).Elem()
	for k := 0; k < t.NumField(); k++ {
		fName := t.Field(k).Name
		dataV := GetStringValueFromStructByName(data, fName)
		p.FieldByName(fName).Set(reflect.ValueOf(dataV))
	}

	for k := 0; k < t.NumField(); k++ {
		fName := t.Field(k).Name
		dataV := GetStringValueFromStructByName(user, fName)
		p.FieldByName(fName).Set(reflect.ValueOf(dataV))
	}

	hc.ManaDoctorID = "10094418"
	hc.ManaUnitID = "320111001005"
	return hc
}
func NewLsRequestData(data HealthCheckDetailIsData, user ChangeHealthCheckRequestBody) LsRequestData {
	hc := LsRequestData{}
	t := reflect.TypeOf(hc)
	p := reflect.ValueOf(&hc).Elem()
	for k := 0; k < t.NumField(); k++ {
		fName := t.Field(k).Name
		dataV := GetStringValueFromStructByName(data, fName)
		p.FieldByName(fName).Set(reflect.ValueOf(dataV))
	}

	for k := 0; k < t.NumField(); k++ {
		fName := t.Field(k).Name
		dataV := GetStringValueFromStructByName(user, fName)
		p.FieldByName(fName).Set(reflect.ValueOf(dataV))
	}

	return hc
}
func NewExaRequestData(data HealthCheckDetailExaData) ExaRequestData {
	hc := ExaRequestData{}
	t := reflect.TypeOf(hc)
	p := reflect.ValueOf(&hc).Elem()
	for k := 0; k < t.NumField(); k++ {
		fName := t.Field(k).Name
		dataV := GetStringValueFromStructByName(data, fName)
		p.FieldByName(fName).Set(reflect.ValueOf(dataV))
	}
	d := fmt.Sprintf("%d-%02d-%02d 00:00:00", time.Now().Year(), time.Now().Month(), time.Now().Day())
	hc.CreateDate = d
	hc.LastModifyDate = d
	hc.HeartRate = ""
	return hc
}
func NewAeRequestData(data HealthCheckDetailAeData, user ChangeHealthCheckRequestBody) AeRequestData {
	hc := AeRequestData{}
	t := reflect.TypeOf(hc)
	p := reflect.ValueOf(&hc).Elem()
	for k := 0; k < t.NumField(); k++ {
		fName := t.Field(k).Name
		dataV := GetStringValueFromStructByName(data, fName)
		p.FieldByName(fName).Set(reflect.ValueOf(dataV))
	}

	for k := 0; k < t.NumField(); k++ {
		fName := t.Field(k).Name
		dataV := GetStringValueFromStructByName(user, fName)
		p.FieldByName(fName).Set(reflect.ValueOf(dataV))
	}
	return hc
}
func NewHaRequestData(data HealthCheckDetailHaData) HaRequestData {
	hc := HaRequestData{}
	t := reflect.TypeOf(hc)
	p := reflect.ValueOf(&hc).Elem()
	for k := 0; k < t.NumField(); k++ {
		fName := t.Field(k).Name
		dataV := GetStringValueFromStructByName(data, fName)
		p.FieldByName(fName).Set(reflect.ValueOf(dataV))
	}
	hc.Abnormality = "1"
	d := fmt.Sprintf("%d-%02d-%02d 00:00:00", time.Now().Year(), time.Now().Month(), time.Now().Day())
	hc.CreateDate = d
	hc.LastModifyDate = d
	return hc
}

func GetStringValueFromStructByName(data interface{}, name string) string {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for k := 0; k < t.NumField(); k++ {
		fName := t.Field(k).Name
		a := strings.ToLower(fName)
		b := strings.ToLower(name)
		if a == b {
			kind := v.Field(k).Kind()
			if kind == reflect.Int {
				intV := v.Field(k).Int()
				return fmt.Sprintf("%d", intV)
			}
			if kind == reflect.Float32 {
				floatV := v.Field(k).Float()
				return fmt.Sprintf("%.2f", floatV)
			}
			if kind == reflect.String {
				return v.Field(k).String()
			}
			if kind == reflect.Struct {
				keyValue := v.Field(k).FieldByName("Key")
				s := keyValue.String()
				return s
			}
		}
	}
	return ""
}

func NewInhospitalListRequestData(typeInt int) InhospitalListData {
	return InhospitalListData{Type: typeInt}
}
func NewMedicineListRequestData() MedicineListData {
	return MedicineListData{
		SituationId:   "",
		Medicine:      "",
		Use:           "",
		EachDose:      "",
		UseDate:       "",
		MedicineYield: "",
	}
}
func NewNiListRequestData() NiListData {
	return NiListData{
		RecordId:        "",
		Name:            "",
		InoculationDate: "",
		InoculationUnit: "",
	}
}

type ChangeHealthCheckResp struct {
	Body ChangeHealthCheckRespBody `json:"body"`
	Code int                       `json:"code"`
}
type ChangeHealthCheckRespBody struct {
	LastModifyUnit Insurance `json:"lastModifyUnit"`
	LastModifyUser Insurance `json:"lastModifyUser"`
	LastModifyDate string    `json:"lastModifyDate"`
}

/**
 *  修改健康检查表
 *  完成
 */
func (s *Server) ChangeHealthCheckRequest(body HealthCheckDetailBody, user ChangeHealthCheckRequestBody) (ChangeHealthCheckResp, error) {
	url := "http://32.33.1.123:8082/pkehr/*.jsonRequest"
	cli := http.Client{}
	reqData := NewChangeHealthCheckRequest(body, user)
	bit, _ := json.Marshal(&reqData)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bit))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Cookie", s.Cookie)
	resp, err := cli.Do(request)
	if err != nil {
		return ChangeHealthCheckResp{}, err
	}
	defer resp.Body.Close()
	bits, _ := ioutil.ReadAll(resp.Body)
	var result ChangeHealthCheckResp
	err = json.Unmarshal(bits, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
