package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RequestJson struct {
	Method        string `json:"method"`
	Schema        string `json:"schema"`
	ServiceAction string `json:"serviceAction"`
	ServiceId     string `json:"serviceId"`
}

type HealthCheckDetailRequest struct {
	Body HealthCheckDetailRequestBody `json:"body"`
	RequestJson
}
type HealthCheckDetailRequestBody struct {
	HealthCheck string `json:"healthCheck"`
}

func NewHealthCheckDetailRequest(healthCheck string) HealthCheckDetailRequest {
	return HealthCheckDetailRequest{
		Body: HealthCheckDetailRequestBody{
			HealthCheck: healthCheck,
		},
		RequestJson: RequestJson{
			Method:        "execute",
			Schema:        "chis.application.hc.schemas.HC_HealthCheck",
			ServiceAction: "getHMNIListOfHTML",
			ServiceId:     "chis.healthCheckService",
		},
	}
}

type HealthCheckDetailResp struct {
	Body HealthCheckDetailBody `json:"body"`
	Code int                   `json:"code"`
}

type HealthCheckDetailBody struct {
	IhList  []interface{}            `json:"ihList"`
	NiList  []interface{}            `json:"niList"`
	HhList  []MsListItem             `json:"hhList"`
	HaData  HealthCheckDetailHaData  `json:"haData"`
	MsList  []interface{}            `json:"msList"`
	HcData  HealthCheckDetailHcData  `json:"hcData"`
	ExaData HealthCheckDetailExaData `json:"exaData"`
	AeData  HealthCheckDetailAeData  `json:"aeData"`
	LsData  HealthCheckDetailIsData  `json:"lsData"`
}

type MsListItem struct {
	CreateUser_text     string `json:"createUser_text"`
	LastModifyUnit_text string `json:"lastModifyUnit_text"`
	LastModifyUnit      string `json:"lastModifyUnit"`
	MedicineYield       string `json:"medicineYield"`
	LastModifyUser      string `json:"lastModifyUser"`
	Use                 string `json:"use"`
	HealthCheck         string `json:"healthCheck"`
	UseDate             string `json:"useDate"`
	MedicineYield_text  string `json:"medicineYield_text"`
	CreateUnit_text     string `json:"createUnit_text"`
	SituationId         string `json:"situationId"`
	EachDose            string `json:"eachDose"`
	CreateUser          string `json:"createUser"`
	CreateUnit          string `json:"createUnit"`
	LastModifyUser_text string `json:"lastModifyUser_text"`
	LastModifyDate      string `json:"lastModifyDate"`
	CreateDate          string `json:"createDate"`
	Medicine            string `json:"medicine"`
	Descr               string `json:"descr"`
}

type HealthCheckDetailHaData struct {
	Abnormality        Insurance `json:"abnormality"`
	PjOther            string    `json:"pjOther"`
	LastModifyUnit     Insurance `json:"lastModifyUnit"`
	LastModifyUser     Insurance `json:"lastModifyUser"`
	Mana               Insurance `json:"mana"`
	TargetWeight       float32   `json:"targetWeight"`
	HealthCheck        string    `json:"healthCheck"`
	Vaccine            string    `json:"vaccine"`
	Abnormality1       string    `json:"abnormality1"`
	CreateUser         Insurance `json:"createUser"`
	Abnormality4       string    `json:"abnormality4"`
	Recognize          string    `json:"recognize"`
	Abnormality3       string    `json:"abnormality3"`
	CreateUnit         Insurance `json:"createUnit"`
	AssessmentId       string    `json:"assessmentId"`
	Abnormality2       string    `json:"abnormality2"`
	RiskfactorsControl Insurance `json:"riskfactorsControl"`
	LastModifyDate     string    `json:"lastModifyDate"`
	CreateDate         string    `json:"createDate"`
}
type HealthCheckDetailHcData struct {
	SelfCare                 Insurance `json:"selfCare"`
	CheckWay                 Insurance `json:"checkWay"`
	LastModifyUser           Insurance `json:"lastModifyUser"`
	CerebrovascularDiseases  Insurance `json:"cerebrovascularDiseases"`
	OtherheartDisease        string    `json:"otherheartDisease"`
	Diastolic                int       `json:"diastolic"`
	HeartDisease             Insurance `json:"heartDisease"`
	Temperature              float32
	Height                   float32
	CheckDate                string
	ManaDoctorId             Insurance `json:"manaDoctorId"`
	KidneyDiseases           Insurance `json:"kidneyDiseases"`
	NonimmuneFlag            Insurance `json:"nonimmuneFlag"`
	EmpiId                   string    `json:"empiId"`
	HealthStatus             Insurance `json:"healthStatus"`
	Breathe                  int       `json:"breathe"`
	VascularDisease          Insurance `json:"VascularDisease"`
	Sf                       string    `json:"sf"`
	OtherVascularDisease     string    `json:"otherVascularDisease"`
	Emotion                  Insurance `json:"emotion"`
	PersonalCode             string    `json:"personalCode"`
	CreateUser               Insurance `json:"createUser"`
	Cognitive                Insurance `json:"cognitive"`
	OtherDiseasesone         Insurance `json:"otherDiseasesone"`
	CreateUnit               Insurance `json:"createUnit"`
	PhrId                    string    `json:"phrId"`
	LastModifyDate           string    `json:"lastModifyDate"`
	OtherDiseasesoneDesc     string    `json:"otherDiseasesoneDesc"`
	NeurologicalDiseasesDesc string    `json:"neurologicalDiseasesDesc"`

	Symptom                      Insurance `json:"symptom"`
	Weight                       float32   `json:"weight"`
	EmotionZf                    string    `json:"emotionZf"`
	CognitiveZf                  string    `json:"cognitiveZf"`
	Bmi                          float32   `json:"bmi"`
	ManaUnitId                   Insurance `json:"manaUnitId"`
	LastModifyUnit               Insurance `json:"lastModifyUnit"`
	NeurologicalDiseases         Insurance `json:"neurologicalDiseases"`
	EyeDiseases                  Insurance `json:"eyeDiseases"`
	HealthCheck                  string    `json:"healthCheck"`
	Pulse                        int       `json:"pulse"`
	Diastolic_L                  float32   `json:"diastolic_L"`
	Waistline                    float32   `json:"waistline"`
	OthercerebrovascularDiseases string    `json:"othercerebrovascularDiseases"`
	SymptomOt                    string    `json:"symptomOt"`
	CreateDate                   string    `json:"createDate"`
	Constriction                 int       `json:"constriction"`
	InhospitalFlag               Insurance `json:"inhospitalFlag"`
	OthereyeDiseases             string    `json:"othereyeDiseases"`
	MedicineFlag                 Insurance `json:"medicineFlag"`
	OtherkidneyDiseases          string    `json:"otherkidneyDiseases"`
	InfamilybedFlag              Insurance `json:"infamilybedFlag"`
	Constriction_L               float32   `json:"constriction_L"`
}
type HealthCheckDetailExaData struct {
	LymphnodesDesc    string    `json:"lymphnodesDesc"`
	SkinDesc          string    `json:"skinDesc"`
	LiverBig          Insurance `json:"liverBig"`
	HeartRate         int       `json:"heartRate"`
	LastModifyUser    Insurance `json:"lastModifyUser"`
	Skin              Insurance `json:"skin"`
	ScleraDesc        string    `json:"scleraDesc"`
	PalaceDesc        string    `json:"palaceDesc"`
	HeartMurmurDesc   string    `json:"heartMurmurDesc"`
	AbdominAltendDesc string    `json:"abdominAltendDesc"`
	FundusDesc        string    `json:"fundusDesc"`
	VaginalDesc       string    `json:"vaginalDesc"`
	AdbominAlmassDesc string    `json:"adbominAlmassDesc"`
	Fundus            Insurance `json:"fundus"`
	Breast            string    `json:"breast"`
	CreateUser        Insurance `json:"createUser"`
	Edema             Insurance `json:"edema"`
	CreateUnit        Insurance `json:"createUnit"`
	LiverBigDesc      string    `json:"liverBigDesc"`
	LastModifyDate    string    `json:"lastModifyDate"`
	Cervix            string    `json:"cervix"`
	Palace            string    `json:"palace"`
	FootPulse         Insurance `json:"footPulse"`
	Vaginal           string    `json:"vaginal"`
	DullnessDesc      string    `json:"dullnessDesc"`
	Examination       string    `json:"examination"`
	Splenomegaly      Insurance `json:"splenomegaly"`
	LastModifyUnit    Insurance `json:"lastModifyUnit"`
	HealthCheck       string    `json:"health_check"`
	AbdominAltend     Insurance `json:"abdominAltend"`
	BarrelChest       Insurance `json:"barrelChest"`
	RalesDesc         string    `json:"ralesDesc"`
	Rhythm            Insurance `json:"rhythm"`
	AttachmentDesc    string    `json:"attachmentDesc"`
	Dre               string    `json:"dre"`
	CreateDate        string    `json:"createDate"`
	VulvaDesc         string    `json:"vulvaDesc"`
	HeartMurmur       Insurance `json:"heartMurmur"`
	DreDesc           string    `json:"dreDesc"`
	Dullness          Insurance `json:"dullness"`
	Rales             Insurance `json:"rales"`
	Attachment        string    `json:"attachment"`
	Sclera            Insurance `json:"sclera"`
	CervixDesc        string    `json:"cervix_desc"`
	BreathSoundDesc   string    `json:"breathSoundDesc"`
	Tjother           string    `json:"tjother"`
	Vulva             string    `json:"vulva"`
	Lymphnodes        Insurance `json:"lymphnodes"`
	BreastDesc        string    `json:"breastDesc"`
	BreathSound       Insurance `json:"breathSound"`
	SplenomegalyDesc  string    `json:"splenomegalyDesc"`
	AdbominAlmass     Insurance `json:"adbominAlmass"`
}
type HealthCheckDetailAeData struct {
	Kalemia        string    `json:"kalemia"`
	RecordId       string    `json:"recordId"`
	Hbsag          string    `json:"hbsag"`
	Fbs2           string    `json:"fbs2"`
	Tg             float32   `json:"tg"`
	LastModifyUser Insurance `json:"lastModifyUser"`
	RecRightEye    string    `json:"recRightEye"`
	PsText         string    `json:"psText"`
	Fbs            float32   `json:"fbs"`
	Glu            string    `json:"glu"`
	Oc             string    `json:"oc"`
	Tc             float32   `json:"tc"`
	RightEye       float32   `json:"rightEye"`
	Motion         Insurance `json:"motion"`
	BloodOther     string    `json:"bloodOther"`
	Dbil           string    `json:"dbil"`
	LeftEye        float32   `json:"leftEye"`
	Proteinuria    string    `json:"proteinuria"`
	Tbil           string    `json:"tbil"`
	Hearing        Insurance `json:"hearing"`
	Cr             float32   `json:"cr"`
	CreateUser     Insurance `json:"createUser"`
	Ldl            float32   `json:"ldl"`
	CreateUnit     Insurance `json:"createUnit"`
	XText          string    `json:"xText"`
	LastModifyDate string    `json:"lastModifyDate"`
	LastModifyUnit Insurance `json:"lastModifyUnit"`
	Ast            float32   `json:"ast"`
	Hgb            float32   `json:"hgb"`
	HealthCheck    string    `json:"healthCheck"`
	EcgText        string    `json:"ecgText"`
	RightUp        int       `json:"rightUp"`
	Alt            float32   `json:"alt"`
	RecLeftEye     string    `json:"recLeftEye"`
	CreateDate     string    `json:"createDate"`
	Denture        Insurance `json:"denture"`
	Wbc            float32   `json:"wbc"`
	LeftDown       int       `json:"leftDown"`
	Malb           string    `json:"malb"`
	Dka            string    `json:"dka"`
	Hdl            float32   `json:"hdl"`
	UrineOther     string    `json:"urineOther"`
	Platelet       float32   `json:"platelet"`
	B              Insurance `json:"b"`
	Hba1c          string    `json:"hba1c"`
	Lip            Insurance `json:"lip"`
	Ecg            Insurance `json:"ecg"`
	Natremia       string    `json:"natremia"`
	Alb            string    `json:"alb"`
	Fob            string    `json:"fob"`
	Bun            float32   `json:"bun"`
	LeftUp         int       `json:"leftUp"`
	RightDown      int       `json:"rightDown"`
	Pharyngeal     Insurance `json:"pharyngeal"`
	FuOther        string    `json:"fuOther"`
	X              Insurance `json:"x"`
	BText          string    `json:"bText"`
	Ps             string    `json:"ps"`
}

type HealthCheckDetailIsData struct {
	ChemicalsPro              string    `json:"chemicalsPro"`
	ExerciseStyle             string    `json:"exerciseStyle"`
	Occupational              Insurance `json:"occupational"`
	Jobs                      string    `json:"jobs"`
	Ray                       string    `json:"ray"`
	OtherProDesc              string    `json:"otherProDesc"`
	Other                     string    `json:"other"`
	LastModifyUnit            Insurance `json:"lastModifyUnit"`
	LastModifyUser            Insurance `json:"lastModifyUser"`
	RayPro                    string    `json:"rayPro"`
	HealthCheck               string    `json:"healthCheck"`
	Insistexercisetime        int64     `json:"insistexercisetime"`
	OtherPro                  string    `json:"otherPro"`
	Chemicals                 string    `json:"chemicals"`
	StopDrinkingTime          string    `json:"stopDrinkingTime"`
	RayProDesc                string    `json:"rayProDesc"`
	DrinkingFrequency         Insurance `json:"drinkingFrequency"`
	WhetherDrink              string    `json:"whetherDrink"`
	WorkTime                  string    `json:"workTime"`
	PhysicalFactorPro         string    `json:"physicalFactorPro"`
	DietaryHabit              Insurance `json:"dietaryHabit"`
	CreateDate                string    `json:"createDate"`
	GeginToDrinkTime          string    `json:"geginToDrinkTime"`
	AlcoholConsumption        string    `json:"alcoholConsumption"`
	DustProDesc               string    `json:"dustProDesc"`
	PhysicalFactorProDesc     string    `json:"physicalFactorProDesc"`
	LifestySituation          string    `json:"lifestySituation"`
	PhysicalExerciseFrequency Insurance `json:"physicalExerciseFrequency"`
	StopSmokeTime             string    `json:"stopSmokeTime"`
	BeginSmokeTime            string    `json:"beginSmokeTime"`
	WehtherSmoke              Insurance `json:"wehtherSmoke"`
	MainDrinkingVvarieties    string    `json:"mainDrinkingVvarieties"`
	CreateUser                Insurance `json:"createUser"`
	ChemicalsProDesc          string    `json:"chemicalsProDesc"`
	CreateUnit                Insurance `json:"createUnit"`
	Dust                      string    `json:"dust"`
	PhysicalFactor            string    `json:"physicalFactor"`
	DrinkOther                string    `json:"drinkOther"`
	IsDrink                   string    `json:"isDrink"`
	LastModifyDate            string    `json:"last_ModifyDate"`
	DustPro                   string    `json:"dustPro"`
	Smokes                    string    `json:"smokes"`
	EveryPhysicalExerciseTime int64     `json:"everyPhysicalExerciseTime"`
}

func (s *Server) RequestHealthCheckDetail(healthCheck string) (HealthCheckDetailResp, error) {
	url := "http://32.33.1.123:8082/pkehr/*.jsonRequest"
	cli := http.Client{}
	reqData := NewHealthCheckDetailRequest(healthCheck)
	bit, _ := json.Marshal(&reqData)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bit))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Cookie", s.Cookie)
	resp, err := cli.Do(request)
	if err != nil {
		return HealthCheckDetailResp{}, err
	}
	defer resp.Body.Close()
	bits, _ := ioutil.ReadAll(resp.Body)
	var result HealthCheckDetailResp
	err = json.Unmarshal(bits, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
