package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HealthCheckRequest struct {
	Actions       string        `json:"actions"`
	Cnd           []interface{} `json:"cnd"`
	ServiceId     string        `json:"serviceId"`
	Method        string        `json:"method"`
	Schema        string        `json:"schema"`
	PageSize      int           `json:"pageSize"`
	PageNo        int           `json:"pageNo"`
	ServiceAction string
}

func NewHealthCheckRequest(pkey string) HealthCheckRequest {

	return HealthCheckRequest{
		Actions:       "update",
		ServiceId:     "chis.simpleQuery",
		Method:        "execute",
		Schema:        "chis.application.hc.schemas.HC_HealthCheck_list",
		PageNo:        1,
		PageSize:      25,
		ServiceAction: "",
		Cnd:           []interface{}{"eq", []string{"$", "a.empiId"}, []string{"s", pkey}},
	}
}

type HealthCheckListResp struct {
	TotalCount int               `json:"totalCount"`
	PageNo     int               `json:"pageNo"`
	PageSize   int               `json:"pageSize"`
	Code       int               `json:"code"`
	Msg        string            `json:"msg"`
	Body       []HealthCheckItem `json:"body"`
}
type HealthCheckItem struct {
	CreateUserText               string  `json:"createUser_text"`
	SelfCare                     string  `json:"selfCare"`
	CheckWay                     string  `json:"checkWay"`
	LastModifyUser               string  `json:"lastModifyUser"`
	CerebrovascularDiseases      string  `json:"cerebrovascularDiseases"`
	OtherheartDisease            string  `json:"otherheartDisease"`
	Diastolic                    int     `json:"diastolic"`
	HeartDisease                 string  `json:"heartDisease"`
	Temperature                  float32 `json:"temperature"`
	NeurologicalDiseasesText     string  `json:"neurologicalDiseases_text"`
	CerebrovascularDiseasesText  string  `json:"cerebrovascularDiseases_text"`
	Height                       float32 `json:"height"`
	ManaDoctorIdText             string  `json:"manaDoctorId_text"`
	CheckDate                    string  `json:"checkDate"`
	ManaDoctorId                 string  `json:"manaDoctorId"`
	EmotionText                  string  `json:"emotion_text"`
	KidneyDiseases               string  `json:"kidneyDiseases"`
	SymptomText                  string  `json:"symptom_text"`
	CheckWay_text                string  `json:"checkWay_text"`
	Cognitive_text               string  `json:"cognitive_text"`
	EmpiId                       string  `json:"empiId"`
	HealthStatus                 string  `json:"healthStatus"`
	ManaUnitId_text              string  `json:"manaUnitId_text"`
	Breathe                      int     `json:"breathe"`
	HealthStatus_text            string  `json:"healthStatus_text"`
	VascularDisease              string  `json:"VascularDisease"`
	OtherVascularDisease         string  `json:"otherVascularDisease"`
	Emotion                      string  `json:"emotion"`
	PersonalCode                 string  `json:"personalCode"`
	CreateUser                   string  `json:"createUser"`
	Cognitive                    string  `json:"cognitive"`
	OtherDiseasesone             string  `json:"otherDiseasesone"`
	CreateUnit                   string  `json:"createUnit"`
	PhrId                        string  `json:"phrId"`
	LastModifyDate               string  `json:"lastModifyDate"`
	OtherDiseasesoneDesc         string  `json:"otherDiseasesoneDesc"`
	NeurologicalDiseasesDesc     string  `json:"neurologicalDiseasesDesc"`
	Symptom                      string  `json:"symptom"`
	Weight                       float32 `json:"weight"`
	CognitiveZf                  string  `json:"cognitiveZf"`
	EmotionZf                    string  `json:"emotionZf"`
	Bmi                          float32 `json:"bmi"`
	LastModifyUnit               string  `json:"lastModifyUnit"`
	ManaUnitId                   string  `json:"manaUnitId"`
	NeurologicalDiseases         string  `json:"neurologicalDiseases"`
	HealthCheck                  string  `json:"healthCheck"`
	EyeDiseases                  string  `json:"eyeDiseases"`
	Pulse                        int     `json:"pulse"`
	CreateUnit_text              string  `json:"createUnit_text"`
	Diastolic_L                  int     `json:"diastolic_L"`
	OtherDiseasesone_text        string  `json:"otherDiseasesone_text"`
	EyeDiseases_text             string  `json:"eyeDiseases_text"`
	Waistline                    float32 `json:"waistline"`
	SymptomOt                    string  `json:"symptomOt"`
	OthercerebrovascularDiseases string  `json:"othercerebrovascularDiseases"`
	CreateDate                   string  `json:"createDate"`
	HeartDisease_text            string  `json:"heartDisease_text"`
	VascularDisease_text         string  `json:"VascularDisease_text"`
	Constriction                 int     `json:"constriction"`
	OthereyeDiseases             string  `json:"othereyeDiseases"`
	KidneyDiseases_text          string  `json:"kidneyDiseases_text"`
	OtherkidneyDiseases          string  `json:"otherkidneyDiseases"`
	Constriction_L               int     `json:"constriction_L"`
	SelfCare_text                string  `json:"selfCare_text"`
}

func (s *Server) RequestHealthCheckList(idcard string) (HealthCheckListResp, error) {
	url := "http://32.33.1.123:8082/pkehr/*.jsonRequest"
	cli := http.Client{}
	reqData := NewHealthCheckRequest(idcard)
	bit, _ := json.Marshal(&reqData)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bit))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Cookie", s.Cookie)
	resp, err := cli.Do(request)
	if err != nil {
		return HealthCheckListResp{}, err
	}
	defer resp.Body.Close()
	bits, _ := ioutil.ReadAll(resp.Body)
	var result HealthCheckListResp
	err = json.Unmarshal(bits, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
