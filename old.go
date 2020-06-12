package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// impOldToTmp 导入老年人中医体质
//
func (s *Server) impOldToTmp(dateStr string) (err error) {
	dest, err := s.FindUserRecordFromGongwei(dateStr)
	if err != nil {
		return
	}
	for _, item := range dest {
		// 根据用户身份证获取记录列表 如果没有当日的记录则新建
		// 如果有则修改
		record, er := s.FindUserRecords(item.IdCard)
		if er != nil {
			err = er
			return
		}
		if IsRecordDateExist(record, dateStr) {
			// 编辑
			rc := RecordByDateStr(record, dateStr)
			item.RecordID = rc.ID
			er = s.NewUserRecord(item)
			if er != nil {
				err = er
				return
			}
			log.Printf("中医体质导入成功 身份证： %s", item.IdCard)
			continue
		}
		// create
		er = s.NewUserRecord(item)
		if er != nil {
			fmt.Sprintf("导入有误 %s", er)
			continue
		}
		log.Printf("中医体质导入成功 身份证： %s", item.IdCard)

	}
	return
}

func IsRecordDateExist(res *FindUserRecordResult, dateStr string) bool {
	for _, item := range res.Body {
		if strings.Contains(item.ReportDate, dateStr) {
			return true
		}
	}
	return false
}
func RecordByDateStr(res *FindUserRecordResult, dateStr string) *FindUserRecordResultBody {
	for _, item := range res.Body {
		if strings.Contains(item.ReportDate, dateStr) {
			return &item
		}
	}
	return nil
}

// FindUserRecordRequest 请求格式
type FindUserRecordRequest struct {
	ServiceId     string        `json:"serviceId"`
	Method        string        `json:"method"`
	Schema        string        `json:"schema"`
	Cnd           []interface{} `json:"cnd"`
	PageSize      int           `json:"pageSize"`
	PageNo        int           `json:"pageNo"`
	ServiceAction string        `json:"serviceAction"`
}

// FindUserRecord 老年人体质记录
// {"serviceId":"chis.simpleQuery",
// "method":"execute",
// "schema":"chis.application.ohr.schemas.MDC_ChineseMedicineManageListView",
// "cnd":["and",["eq",["$","a.status"],["s","0"]],["like",["$","b.idCard"],["s","%320122195002074443%"]]],
// "pageSize":25,"pageNo":1,"serviceAction":""}
func (s *Server) NewRequest(idCard string) FindUserRecordRequest {
	return FindUserRecordRequest{
		Cnd: []interface{}{
			"and",
			[]interface{}{
				"eq", []interface{}{"$", "a.status"}, []interface{}{"s", "0"},
			},
			[]interface{}{
				"like",
				[]interface{}{"$", "b.idCard"},
				[]interface{}{"s", "%" + idCard + "%"},
			},
		},
		ServiceId: "chis.simpleQuery",
		Method:    "execute",
		Schema:    "chis.application.ohr.schemas.MDC_ChineseMedicineManageListView",
		PageSize:  25,
		PageNo:    1,
	}
}

type FindUserRecordResultBody struct {
	Birthday           string `json:"birthday"`
	CreateUserText     string `json:"createUser_text"`
	SexCodeText        string `json:"sexCode_text"`
	LastModifyUnit     string `json:"lastModifyUnit"`
	ManaUnitID         string `json:"manaUnitId"`
	ReportUser         string `json:"reportUser"`
	LastModifyUser     string `json:"lastModifyUser"`
	CreateUnitText     string `json:"createUnit_text"`
	PersonName         string `json:"personName"`
	BodyType           string `json:"bodyType"`
	ID                 string `json:"id"`
	ManaDoctorIDText   string `json:"manaDoctorId_text"`
	ManaDoctorID       string `json:"manaDoctorId"`
	ReportDate         string `json:"reportDate"`
	CreateDate         string `json:"createDate"`
	SexCode            string `json:"sexCode"`
	RegionCode         string `json:"regionCode"`
	EmpiID             string `json:"empiId"`
	IDCard             string `json:"idCard"`
	RegionCodeText     string `json:"regionCode_text"`
	ManaUnitIDText     string `json:"manaUnitId_text"`
	LastModifyUnitText string `json:"lastModifyUnit_text"`
	Status             string `json:"status"`
	ReportUserText     string `json:"reportUser_text"`
	BodyTypeText       string `json:"bodyType_text"`
	StatusText         string `json:"status_text"`
	CreateUser         string `json:"createUser"`
	CreateUnit         string `json:"createUnit"`
	LastModifyUserText string `json:"lastModifyUser_text"`
	PhrID              string `json:"phrId"`
	LastModifyDate     string `json:"lastModifyDate"`
	MobileNumber       string `json:"mobileNumber"`
}
type FindUserRecordResult struct {
	Body       []FindUserRecordResultBody `json:"body"`
	TotalCount int                        `json:"totalCount"`
	PageNo     int                        `json:"pageNo"`
	PageSize   int                        `json:"pageSize"`
	Code       int                        `json:"code"`
	Msg        string                     `json:"msg"`
}

func (s *Server) FindUserRecords(idCard string) (res *FindUserRecordResult, err error) {
	reqData := s.NewRequest(idCard)
	url := "http://32.33.1.123:8082/pkehr/*.jsonRequest?start=0&limit=25"
	cli := http.Client{}
	bit, _ := json.Marshal(&reqData)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bit))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Cookie", s.Cookie)
	resp, err := cli.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bits, _ := ioutil.ReadAll(resp.Body)
	var result FindUserRecordResult
	err = json.Unmarshal(bits, &result)
	if err != nil {
		return
	}
	res = &result
	return
}

type CreateUserRecordRequest struct {
	ServiceId     string                      `json:"serviceId"`
	Method        string                      `json:"method"`
	Op            string                      `json:"op"`
	Schema        string                      `json:"schema"`
	ServiceAction string                      `json:"serviceAction"`
	Body          CreateUserRecordRequestBody `json:"body"`
}
type CreateUserRecordRequestBody struct {
	ID                string      `json:"id"`
	PhrID             string      `json:"phrId"`
	EmpiID            string      `json:"empiId"`
	EnergyFull        int         `json:"energyFull"`
	EasyWeary         int         `json:"easyWeary"`
	EasyPant          int         `json:"easyPant"`
	VoiceWeak         int         `json:"voiceWeak"`
	Moodiness         int         `json:"moodiness"`
	Nervous           int         `json:"nervous"`
	Loneliness        int         `json:"loneliness"`
	EasyScare         int         `json:"easyScare"`
	Overweight        int         `json:"overweight"`
	EyeDry            int         `json:"eyeDry"`
	FootFearCold      int         `json:"footFearCold"`
	BackFearCold      int         `json:"backFearCold"`
	FearCold          int         `json:"fearCold"`
	Cold              int         `json:"cold"`
	Rhinobyon         int         `json:"rhinobyon"`
	MouthGreasy       int         `json:"mouthGreasy"`
	Allergy           int         `json:"allergy"`
	SkinUrticaria     int         `json:"skinUrticaria"`
	SkinBleeding      int         `json:"skinBleeding"`
	SkinRed           int         `json:"skinRed"`
	SkinDry           int         `json:"skinDry"`
	LimbsNumb         int         `json:"limbsNumb"`
	FaceGreasy        int         `json:"faceGreasy"`
	FaceDim           int         `json:"faceDim"`
	SkinEczema        int         `json:"skinEczema"`
	MouthDry          int         `json:"mouthDry"`
	BitterTaste       int         `json:"bitterTaste"`
	BellyLarge        int         `json:"bellyLarge"`
	FearCool          int         `json:"fearCool"`
	StoolStiction     int         `json:"stoolStiction"`
	StoolDry          int         `json:"stoolDry"`
	FurStodgily       int         `json:"furStodgily"`
	StasisPurple      int         `json:"stasisPurple"`
	Score1            string      `json:"score1"`
	PhysiqueIdentify1 interface{} `json:"physiqueIdentify1"`
	HealthGuide1      string      `json:"healthGuide1"`
	Other1            string      `json:"other1"`
	Score2            string      `json:"score2"`
	PhysiqueIdentify2 interface{} `json:"physiqueIdentify2"`
	HealthGuide2      string      `json:"healthGuide2"`
	Other2            string      `json:"other2"`
	Score3            string      `json:"score3"`
	PhysiqueIdentify3 interface{} `json:"physiqueIdentify3"`
	HealthGuide3      string      `json:"healthGuide3"`
	Other3            string      `json:"other3"`
	Score4            string      `json:"score4"`
	PhysiqueIdentify4 interface{} `json:"physiqueIdentify4"`
	HealthGuide4      string      `json:"healthGuide4"`
	Other4            string      `json:"other4"`
	Score5            string      `json:"score5"`
	PhysiqueIdentify5 interface{} `json:"physiqueIdentify5"`
	HealthGuide5      string      `json:"healthGuide5"`
	Other5            string      `json:"other5"`
	Score6            string      `json:"score6"`
	PhysiqueIdentify6 interface{} `json:"physiqueIdentify6"`
	HealthGuide6      string      `json:"healthGuide6"`
	Other6            string      `json:"other6"`
	Score7            string      `json:"score7"`
	PhysiqueIdentify7 interface{} `json:"physiqueIdentify7"`
	HealthGuide7      string      `json:"healthGuide7"`
	Other7            string      `json:"other7"`
	Score8            string      `json:"score8"`
	PhysiqueIdentify8 interface{} `json:"physiqueIdentify8"`
	HealthGuide8      string      `json:"healthGuide8"`
	Other8            string      `json:"other8"`
	Score9            string      `json:"score9"`
	PhysiqueIdentify9 interface{} `json:"physiqueIdentify9"`
	HealthGuide9      string      `json:"healthGuide9"`
	Other9            string      `json:"other9"`
	ReportDate        string      `json:"reportDate"`
	ReportUser        string      `json:"reportUser"`
	Status            string      `json:"status"`
	BodyType          string      `json:"bodyType"`
	CreateUnit        string      `json:"createUnit"`
	CreateUser        string      `json:"createUser"`
	CreateDate        string      `json:"createDate"`
	LastModifyUser    string      `json:"lastModifyUser"`
	LastModifyDate    string      `json:"lastModifyDate"`
	LastModifyUnit    string      `json:"lastModifyUnit"`
}

type PhysiqueIdentify struct {
	Key  string `json:"key"`
	Text string `json:"text"`
}

// "serviceId":"chis.chineseMedicineManageService",
//    "method":"execute",
//    "op":"create",
//    "schema":"chis.application.ohr.schemas.MDC_ChineseMedicineManage",
// 		saveChineseMedicineManage
//
type UserRecordAnswerGongwei struct {
	AnswerContent string
	QuestionID    string
	Sort          int
}
type CorAnswerResponse struct {
	Judges     []*CorAnswerItem `json:"judges"`
	Result     []string         `json:"result"`
	UserResult []string         `json:"userResult"`
	OperatorID string           `json:"operatorId"`
}

type CorAnswerItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value uint32 `json:"value"`
	Tag   int32  `json:"tag"` // 1是 2倾向是 0不是
}

func (s *Server) NewCreateUserRecordRequest(us *UserRecordGongwei) (res *CreateUserRecordRequest, err error) {
	res = &CreateUserRecordRequest{
		ServiceId:     "chis.chineseMedicineManageService",
		Method:        "execute",
		Op:            "create",
		Schema:        "chis.application.ohr.schemas.MDC_ChineseMedicineManage",
		ServiceAction: "saveChineseMedicineManage",
	}

	if us.RecordID != "" {
		res.Op = "update"
	}

	emInfo, err := s.UserEmInfo(us.IdCard)
	if err != nil {
		return
	}

	body := CreateUserRecordRequestBody{
		EmpiID:         emInfo.EmpiData.EmpiID,
		PhrID:          emInfo.PhrID,
		ReportUser:     s.ManaDoctorID,
		Status:         emInfo.EmpiData.Status,
		BodyType:       "9",
		CreateUnit:     emInfo.EmpiData.CreateUnit,
		CreateUser:     emInfo.EmpiData.CreateUser,
		LastModifyUser: emInfo.EmpiData.LastModifyUser,
		LastModifyUnit: emInfo.EmpiData.LastModifyUnit,
		ID:             us.RecordID,
	}
	if us.CreatedAt != nil {
		t := fmt.Sprintf("%d-%02d-%02d", us.CreatedAt.Year(), us.CreatedAt.Month(), us.CreatedAt.Day())
		body.ReportDate = fmt.Sprintf("%sT00:00:00", t)

		body.CreateDate = us.CreatedAt.Format("2006-01-02 15:04:05")
		body.LastModifyDate = us.CreatedAt.Format("2006-01-02 15:04:05")
	}

	sql := `SELECT r.question_id, r.answer_content, cq.sort FROM records r 
				INNER JOIN corporeity_questions cq ON cq.id = r.question_id
				WHERE user_record_id = ?`
	d := make([]*UserRecordAnswerGongwei, 0)
	if err = s.DB.Raw(sql, us.ID).Scan(&d).Error; err != nil {
		return
	}

	for _, item := range d {
		if item.Sort == 1 {
			body.EnergyFull = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 2 {
			body.EasyWeary = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 3 {
			body.EasyPant = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 4 {
			body.VoiceWeak = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 5 {
			body.Moodiness = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 6 {
			body.Nervous = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 7 {
			body.Loneliness = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 8 {
			body.EasyScare = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 9 {
			body.Overweight = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 10 {
			body.EyeDry = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 11 {
			body.FootFearCold = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 12 {
			body.BackFearCold = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 13 {
			body.FearCold = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 14 {
			body.Cold = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 15 {
			body.Rhinobyon = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 16 {
			body.MouthGreasy = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 17 {
			body.Allergy = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 18 {
			body.SkinUrticaria = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 19 {
			body.SkinBleeding = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 20 {
			body.SkinRed = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 21 {
			body.SkinDry = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 22 {
			body.LimbsNumb = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 23 {
			body.FaceGreasy = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 24 {
			body.FaceDim = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 25 {
			body.SkinEczema = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 26 {
			body.MouthDry = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 27 {
			body.BitterTaste = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 28 {
			body.BellyLarge = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 29 {
			body.FearCool = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 30 {
			body.StoolStiction = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 31 {
			body.StoolDry = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 32 {
			body.FurStodgily = ContentToValue(item.AnswerContent)
		}
		if item.Sort == 33 {
			body.StasisPurple = ContentToValue(item.AnswerContent)
		}

	}
	// 系统内存的答案

	path := "https://gongwei-api.julu666.com/app/chart/calculate?recordId=" + us.ID
	cli := http.Client{}
	request, _ := http.NewRequest("GET", path, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Cookie", s.GongweiToken)
	resp, err := cli.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	da, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var rsp CorAnswerResponse
	err = json.Unmarshal(da, &rsp)
	if err != nil {
		return
	}

	bodyTypes := make([]string, 0)
	// 系统内存的体质辨识
	for index, im := range rsp.Judges {
		if im.Tag != 0 {
			bodyTypes = append(bodyTypes, fmt.Sprintf("%d", index+1))
		}

		if im.Name == "气虚质" {
			body.Score1 = fmt.Sprintf("%d", im.Value)
			if im.Tag == 0 {
				body.PhysiqueIdentify1 = PhysiqueIdentify{}
			} else {
				body.PhysiqueIdentify1 = im.Tag
			}
		}
		if im.Name == "阳虚质" {
			body.Score2 = fmt.Sprintf("%d", im.Value)
			if im.Tag == 0 {
				body.PhysiqueIdentify2 = PhysiqueIdentify{}
			} else {
				body.PhysiqueIdentify2 = im.Tag
			}
		}
		if im.Name == "阴虚质" {
			body.Score3 = fmt.Sprintf("%d", im.Value)
			if im.Tag == 0 {
				body.PhysiqueIdentify3 = PhysiqueIdentify{}
			} else {
				body.PhysiqueIdentify3 = im.Tag
			}
		}
		if im.Name == "痰虚质" {
			body.Score4 = fmt.Sprintf("%d", im.Value)
			if im.Tag == 0 {
				body.PhysiqueIdentify4 = PhysiqueIdentify{}
			} else {
				body.PhysiqueIdentify4 = im.Tag
			}
		}
		if im.Name == "湿热质" {
			body.Score5 = fmt.Sprintf("%d", im.Value)
			if im.Tag == 0 {
				body.PhysiqueIdentify5 = PhysiqueIdentify{}
			} else {
				body.PhysiqueIdentify5 = im.Tag
			}
		}
		if im.Name == "血瘀质" {
			body.Score6 = fmt.Sprintf("%d", im.Value)
			if im.Tag == 0 {
				body.PhysiqueIdentify6 = PhysiqueIdentify{}
			} else {
				body.PhysiqueIdentify6 = im.Tag
			}
		}
		if im.Name == "气郁质" {
			body.Score7 = fmt.Sprintf("%d", im.Value)
			if im.Tag == 0 {
				body.PhysiqueIdentify7 = PhysiqueIdentify{}
			} else {
				body.PhysiqueIdentify7 = im.Tag
			}
		}
		if im.Name == "特禀质" {
			body.Score8 = fmt.Sprintf("%d", im.Value)
			if im.Tag == 0 {
				body.PhysiqueIdentify8 = PhysiqueIdentify{}
			} else {
				body.PhysiqueIdentify8 = im.Tag
			}
		}
		if im.Name == "平和质" {
			body.Score9 = fmt.Sprintf("%d", im.Value)
			if im.Tag == 0 {
				body.PhysiqueIdentify9 = PhysiqueIdentify{}
			} else {
				body.PhysiqueIdentify9 = im.Tag
			}
		}
	}
	res.Body = body
	return
}
func ContentToValue(content string) int {
	switch content {
	case "没有":
		return 1
	case "很少":
		return 2
	case "有时":
		return 3
	case "经常":
		return 4
	case "总是":
		return 5
	}
	return 0
}

func (s *Server) NewUserRecord(us *UserRecordGongwei) (err error) {
	reqData, err := s.NewCreateUserRecordRequest(us)
	if err != nil {
		return
	}
	path := "http://32.33.1.123:8082/pkehr/*.jsonRequest"
	cli := http.Client{}
	bit, _ := json.Marshal(&reqData)
	request, _ := http.NewRequest("POST", path, bytes.NewReader(bit))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Cookie", s.Cookie)
	resp, err := cli.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bits, _ := ioutil.ReadAll(resp.Body)
	var rsp SuccessResponse
	err = json.Unmarshal(bits, &rsp)
	if err != nil {
		return
	}
	if rsp.Code != 200 {
		err = errors.New("导入失败")
	}
	return
}

type SuccessResponse struct {
	Body struct {
		LastModifyUnit struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"lastModifyUnit"`
		LastModifyUser struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"lastModifyUser"`
		LastModifyDate string `json:"lastModifyDate"`
	} `json:"body"`
	Code int `json:"code"`
}

func (s *Server) ChangeUserRecord(us *UserRecordGongwei) (err error) {

	return
}

type UserRecordGongwei struct {
	ID        string
	UserID    string
	IdCard    string
	CreatedAt *time.Time

	//
	RecordID string
}

// FindUserRecordFromGongwei 找寻共卫系统的老年人中医体质
// dateStr 日期格式2020-05-25
func (s *Server) FindUserRecordFromGongwei(dateStr string) (res []*UserRecordGongwei, err error) {
	start := dateStr + " 00:00:00"
	end := dateStr + " 23:59:59"

	recordSQL := `SELECT ur.id, user_id, u.id_card, ur.created_at FROM user_records ur
					INNER JOIN users u ON u.id = ur.user_id
					WHERE ur.created_at BETWEEN ? AND ?`
	dest := make([]*UserRecordGongwei, 0)
	err = s.DB.Raw(recordSQL, start, end).Scan(&dest).Error
	if err != nil {
		return
	}
	res = dest
	return
}

type EmpInfo struct {
	Body    interface{} `json:"body"`
	Control struct {
		EHRPoorPeopleVisitControl struct {
			Create bool `json:"create"`
		} `json:"EHR_PoorPeopleVisit_control"`
		EHRHealthRecordControl struct {
			Update bool `json:"update"`
		} `json:"EHR_HealthRecord_control"`
		HCHealthCheckControl struct {
			Update bool `json:"update"`
		} `json:"HC_HealthCheck_control"`
		MDCOldPeopleRecordControl struct {
			Update bool `json:"update"`
			Create bool `json:"create"`
		} `json:"MDC_OldPeopleRecord_control"`
	} `json:"control"`
	HealthCheckType    string `json:"healthCheckType"`
	PostnatalVisitType string `json:"postnatalVisitType"`
	Ids                struct {
		EmpiID                        string      `json:"empiId"`
		PhrIDStatus                   string      `json:"phrId.status"`
		FamilyID                      string      `json:"familyId"`
		PhrID                         string      `json:"phrId"`
		EHRPoorPeopleVisitVisitID     interface{} `json:"EHR_PoorPeopleVisit.visitId"`
		Brid                          string      `json:"brid"`
		MDCOldPeopleRecordPhrID       string      `json:"MDC_OldPeopleRecord.phrId"`
		HCHealthCheckPhrID            interface{} `json:"HC_HealthCheck.phrId"`
		MDCOldPeopleRecordPhrIDStatus string      `json:"MDC_OldPeopleRecord.phrId.status"`
	} `json:"ids"`
	DebilityShowType string      `json:"debilityShowType"`
	ZYH              interface{} `json:"ZYH"`
	Code             int         `json:"code"`
	Msg              string      `json:"msg"`
	EmpiData         struct {
		CreateUserText          string      `json:"createUser_text"`
		InsuranceCodeText       string      `json:"insuranceCode_text"`
		LastModifyUser          string      `json:"lastModifyUser"`
		ZYH                     interface{} `json:"ZYH"`
		BloodTypeCode           string      `json:"bloodTypeCode"`
		PersonName              string      `json:"personName"`
		NationCode              string      `json:"nationCode"`
		WorkPlace               string      `json:"workPlace"`
		BloodTypeCodeText       string      `json:"bloodTypeCode_text"`
		EducationCodeText       string      `json:"educationCode_text"`
		MZHM                    string      `json:"MZHM"`
		RegisteredPermanentText string      `json:"registeredPermanent_text"`
		PhoneNumber             string      `json:"phoneNumber"`
		Age                     int         `json:"age"`
		ManaDoctorIDText        string      `json:"manaDoctorId_text"`
		ManaDoctorID            string      `json:"manaDoctorId"`
		EmpiID                  string      `json:"empiId"`
		Status                  string      `json:"status"`
		Zlls                    string      `json:"zlls"`
		BRID                    int         `json:"BRID"`
		LifeCycle               string      `json:"lifeCycle"`
		WorkCode                string      `json:"workCode"`
		CreateUser              string      `json:"createUser"`
		EducationCode           string      `json:"educationCode"`
		LifeCycleText           string      `json:"lifeCycle_text"`
		CreateUnit              string      `json:"createUnit"`
		BRXZ                    string      `json:"BRXZ"`
		PhrID                   string      `json:"phrId"`
		Temp                    int64       `json:"temp"`
		MaritalStatusCode       string      `json:"maritalStatusCode"`
		VersionNumber           string      `json:"versionNumber"`
		InsuranceCode           string      `json:"insuranceCode"`
		Birthday                string      `json:"birthday"`
		CreateTime              string      `json:"createTime"`
		SexCodeText             string      `json:"sexCode_text"`
		LastModifyUnit          string      `json:"lastModifyUnit"`
		CreateUnitText          string      `json:"createUnit_text"`
		Contact                 string      `json:"contact"`
		ZllsText                string      `json:"zlls_text"`
		RhBloodCode             string      `json:"rhBloodCode"`
		RegisteredPermanent     string      `json:"registeredPermanent"`
		NationalityCode         string      `json:"nationalityCode"`
		MaritalStatusCodeText   string      `json:"maritalStatusCode_text"`
		NationalityCodeText     string      `json:"nationalityCode_text"`
		WorkCodeText            string      `json:"workCode_text"`
		SexCode                 string      `json:"sexCode"`
		IDCard                  string      `json:"idCard"`
		ContactPhone            string      `json:"contactPhone"`
		LastModifyUnitText      string      `json:"lastModifyUnit_text"`
		NationCodeText          string      `json:"nationCode_text"`
		Photo                   string      `json:"photo"`
		Address                 string      `json:"address"`
		LastModifyUserText      string      `json:"lastModifyUser_text"`
		MobileNumber            string      `json:"mobileNumber"`
		RhBloodCodeText         string      `json:"rhBloodCode_text"`
		LastModifyTime          string      `json:"lastModifyTime"`
	} `json:"empiData"`
	AreaGridShowType   string `json:"areaGridShowType"`
	PhrID              string `json:"phrId"`
	Postnatal42DayType string `json:"postnatal42dayType"`
}

type UserEmInfoRequest struct {
	ServiceID string                `json:"serviceId"`
	Method    string                `json:"method"`
	Body      UserEmInfoRequestBody `json:"body"`
}
type UserEmInfoRequestBody struct {
	EmpiID    string `json:"empiId"`
	IdsLoader string `json:"idsLoader"`
}

func NewUserEmInfoRequest(empID string) UserEmInfoRequest {
	return UserEmInfoRequest{
		ServiceID: "chis.idsLoader",
		Method:    "execute",
		Body:      UserEmInfoRequestBody{EmpiID: empID, IdsLoader: "chis.oldPeopleRecordIdLoader,chis.poorPeopleRecordIdLoader,chis.healthCheckIdLoader"},
	}
}
func (s *Server) UserEmInfo(idCard string) (info *EmpInfo, err error) {
	userNewInfo, err := s.RequestUserNewInfo(idCard)
	if err != nil || len(userNewInfo.Body) == 0 {
		err = errors.New("获取用户信息失败")
		return
	}
	path := "http://32.33.1.123:8082/pkehr/*.jsonRequest"
	cli := http.Client{}
	reqData := NewUserEmInfoRequest(userNewInfo.Body[0].EmpiID)
	bit, _ := json.Marshal(&reqData)
	request, _ := http.NewRequest("POST", path, bytes.NewReader(bit))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Cookie", s.Cookie)
	resp, err := cli.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bits, _ := ioutil.ReadAll(resp.Body)
	var rsp EmpInfo
	err = json.Unmarshal(bits, &rsp)
	if err != nil {
		return
	}
	info = &rsp
	return
}
