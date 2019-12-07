package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserIdsRequestBody struct {
	EmpiId    string `json:"empiId"`
	IdsLoader string `json:"idsLoader"` //
}

type UserIdsRequestJson struct {
	Body UserIdsRequestBody `json:"body"`
	RequestJson
}

func NewUserIdsRequestJson(empiId string) UserIdsRequestJson {
	return UserIdsRequestJson{
		Body: UserIdsRequestBody{
			EmpiId:    empiId,
			IdsLoader: "chis.oldPeopleRecordIdLoader,chis.poorPeopleRecordIdLoader,chis.healthCheckIdLoader",
		},
		RequestJson: RequestJson{
			Method:    "execute",
			ServiceId: "chis.idsLoader",
		},
	}
	//empiId: 32012219401129042600000000000000
}

type UserIdsResp struct {
	Ids  UserIds `json:"ids"`
	Code int     `json:"code"`
}

type UserIds struct {
	EmpiID                        string      `json:"empiId"`
	PhrIDStatus                   string      `json:"phrId.status"`
	FamilyID                      string      `json:"familyId"`
	PhrID                         string      `json:"phrId"`
	EHRPoorPeopleVisitVisitID     interface{} `json:"EHR_PoorPeopleVisit.visitId"`
	MDCOldPeopleRecordPhrID       string      `json:"MDC_OldPeopleRecord.phrId"`
	HCHealthCheckPhrID            interface{} `json:"HC_HealthCheck.phrId"`
	MDCOldPeopleRecordPhrIDStatus string      `json:"MDC_OldPeopleRecord.phrId.status"`
}

func (s *Server) RequestUserIds(empiId string) (UserIdsResp, error) {
	url := "http://32.33.1.123:8082/pkehr/*.jsonRequest"
	cli := http.Client{}
	reqData := NewUserIdsRequestJson(empiId)
	bit, _ := json.Marshal(&reqData)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bit))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Cookie", s.Cookie)
	resp, err := cli.Do(request)
	if err != nil {
		return UserIdsResp{}, err
	}
	defer resp.Body.Close()
	bits, _ := ioutil.ReadAll(resp.Body)
	var result UserIdsResp
	err = json.Unmarshal(bits, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
