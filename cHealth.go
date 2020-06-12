package main

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (s *Server) AddUserHealthInfo(user ChangeHealthCheckRequestBody) {
	userNewInfo, err := s.RequestUserNewInfo(user.IdCard)
	if err != nil || len(userNewInfo.Body) == 0 {
		log.Printf("request err %v", err)
	}
	//获取用户empiId
	userInfo, err := s.RequestUserInfo(userNewInfo.Body[0].EmpiID)
	if err != nil {
		log.Printf("request err %v", err)
	}

	// 获取phrid
	userIds, err := s.RequestUserIds(userInfo.Body.EmpiId)
	if err != nil {
		log.Printf("request err %v", err)
	}

	// 获取体检记录
	res, err := s.RequestHealthCheckList(userInfo.Body.EmpiId)
	if err != nil {
		log.Printf("request err %v", err)
		return
	}

	// 没有任何体检记录
	if len(res.Body) == 0 {
		// 新建
		s.NewHealthData(user, userInfo.Body.EmpiId, userIds.Ids.PhrID)
		return
	}
	cd := strings.Split(user.HcData.CheckDate, "T")
	if len(cd) == 0 {
		log.Printf("健康日期有误")
		return
	}
	today := cd[0]
	item, tag := s.HealthCheckItemContain(res.Body, today)
	log.Printf("正在导入 %v", userInfo.Body.IdCard)
	if tag {
		resDetail, err := s.RequestHealthCheckDetail(item.HealthCheck)
		if err != nil {
			log.Printf("request err %v", err)
		}
		changeRes, err := s.ChangeHealthCheckRequest(resDetail.Body, user)
		if err != nil {
			log.Printf("request err %v", err)
		}
		log.Printf("resDetail %d", changeRes.Code)
	} else {
		s.NewHealthData(user, userInfo.Body.EmpiId, userIds.Ids.PhrID)
	}

	log.Printf("request ok %v", res)
}

func (s *Server) HealthCheckItemContain(list []HealthCheckItem, date string) (HealthCheckItem, bool) {
	for _, item := range list {
		if strings.Contains(item.CheckDate, date) {
			return item, true
		}
	}
	return HealthCheckItem{}, false
}

func (s *Server) NewHealthData(user ChangeHealthCheckRequestBody, empiId, phrId string) {
	detail := HealthCheckDetailBody{
		IhList: nil,
		NiList: nil,
		HhList: make([]MsListItem, 0),
		HaData: HealthCheckDetailHaData{},
		MsList: nil,
		HcData: HealthCheckDetailHcData{
			EmpiId: empiId,
			PhrId:  phrId,
		},
		ExaData: HealthCheckDetailExaData{},
		AeData:  HealthCheckDetailAeData{},
		LsData:  HealthCheckDetailIsData{},
	}
	changeRes, err := s.ChangeHealthCheckRequest(detail, user)
	if err != nil {
		log.Printf("request err %v", err)
	}
	log.Printf("resDetail %d", changeRes.Code)
}

func (s *Server) HealthExcelToCUsers() ([]ChangeHealthCheckRequestBody, error) {
	var users []ChangeHealthCheckRequestBody
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "健康检查表-导入模板.xlsx") {
			users, err = s.HealthExcelPathToData(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Server) HealthExcelPathToData(path string) ([]ChangeHealthCheckRequestBody, error) {
	users := make([]ChangeHealthCheckRequestBody, 0)
	xl, err := xlsx.OpenFile(path)
	if err != nil {
		return nil, err
	}
	if len(xl.Sheets) <= 0 {
		return nil, errors.New("sheet number is wrong")
	}
	manSheet := xl.Sheets[0]
	for index, row := range manSheet.Rows {
		if index == 0 {
			continue
		}
		user := ChangeHealthCheckRequestBody{
			HcData:             HcRequestData{},
			LsData:             LsRequestData{},
			ExaData:            ExaRequestData{},
			AeData:             AeRequestData{},
			HaData:             HaRequestData{},
			InhospitalListData: nil,
			MedicineListData:   nil,
			NiListData:         nil,
		}
		user.IdCard = row.Cells[0].String()
		createDate := time.Now().Format("2006-01-02 15:04:05")
		user.HcData.CreateDate = createDate
		user.HcData.LastModifyDate = createDate
		user.LsData.CreateDate = createDate
		user.LsData.LastModifyDate = createDate
		user.ExaData.CreateDate = createDate
		user.ExaData.LastModifyDate = createDate
		user.AeData.CreateDate = createDate
		user.AeData.LastModifyDate = createDate
		user.HaData.CreateDate = createDate
		user.HaData.LastModifyDate = createDate

		t, err := time.Parse("20060102", row.Cells[2].String())
		if err != nil {
			log.Printf("idcard is %s 检查日期有误", user.IdCard)
			continue
		}

		user.HcData.CheckDate = fmt.Sprintf("%d-%02d-%02dT00:00:00", t.Year(), t.Month(), t.Day())
		user.HcData.CheckWay = row.Cells[3].String() // 体检类型
		user.HcData.Symptom = "0" + row.Cells[4].String()
		user.HcData.Temperature = row.Cells[5].String()
		user.HcData.Breathe = row.Cells[7].String()
		user.HcData.Pulse = row.Cells[6].String()
		user.HcData.Constriction = row.Cells[8].String()
		user.HcData.Diastolic = row.Cells[9].String()
		user.HcData.ConstrictionL = row.Cells[10].String()
		user.HcData.DiastolicL = row.Cells[11].String()
		user.HcData.Height = row.Cells[12].String()
		user.HcData.Weight = row.Cells[13].String()
		user.HcData.Waistline = row.Cells[14].String()
		user.HcData.Bmi = Bmi(user.HcData.Height, user.HcData.Weight)
		if row.Cells[15].String() == "0" {
			user.HcData.HealthStatus = ""
		} else {
			user.HcData.HealthStatus = row.Cells[15].String()
		}
		if row.Cells[16].String() == "0" {
			user.HcData.SelfCare = ""
		} else {
			user.HcData.SelfCare = row.Cells[16].String()
		}
		if row.Cells[17].String() == "0" {
			user.HcData.Cognitive = ""
			user.HcData.CognitiveZf = ""
		} else {
			user.HcData.Cognitive = row.Cells[17].String()
		}
		if row.Cells[18].String() == "0" {
			user.HcData.Emotion = ""
		} else {
			user.HcData.EmotionZf = row.Cells[18].String()
		}

		// LsData
		user.LsData.PhysicalExerciseFrequency = row.Cells[19].String()
		user.LsData.EveryPhysicalExerciseTime = row.Cells[20].String()
		user.LsData.Insistexercisetime = row.Cells[21].String()
		user.LsData.ExerciseStyle = row.Cells[22].String()
		user.LsData.DietaryHabit = row.Cells[23].String()
		user.LsData.WehtherSmoke = row.Cells[24].String()
		user.LsData.Smokes = row.Cells[25].String()
		user.LsData.BeginSmokeTime = row.Cells[26].String()
		user.LsData.StopSmokeTime = row.Cells[27].String()
		user.LsData.DrinkingFrequency = row.Cells[28].String()
		user.LsData.AlcoholConsumption = row.Cells[29].String()
		user.LsData.WhetherDrink = row.Cells[30].String()
		user.LsData.GeginToDrinkTime = row.Cells[31].String()
		user.LsData.StopDrinkingTime = row.Cells[32].String()
		user.LsData.MainDrinkingVvarieties = ""
		user.LsData.DrinkOther = row.Cells[33].String()
		user.LsData.Occupational = row.Cells[34].String()

		// aeData
		user.AeData.Lip = row.Cells[35].String()
		user.AeData.Denture = row.Cells[36].String()
		user.AeData.LeftUp = row.Cells[37].String()
		user.AeData.LeftDown = row.Cells[38].String()
		user.AeData.RightUp = row.Cells[39].String()
		user.AeData.RightDown = row.Cells[40].String()

		// 老37
		number := 41
		user.AeData.Pharyngeal = row.Cells[number].String()
		number += 1
		user.AeData.LeftEye = row.Cells[number].String()
		number += 1
		user.AeData.RightEye = row.Cells[number].String()
		number += 1
		user.AeData.RecLeftEye = row.Cells[number].String()
		number += 1
		user.AeData.RecRightEye = row.Cells[number].String()
		number += 1
		user.AeData.Hearing = row.Cells[number].String()
		number += 1
		user.AeData.Motion = row.Cells[number].String()

		//exaData
		number += 1
		user.ExaData.Fundus = row.Cells[number].String()
		number += 1
		user.ExaData.Skin = row.Cells[number].String()
		number += 1
		user.ExaData.Sclera = row.Cells[number].String()
		number += 1
		user.ExaData.Lymphnodes = row.Cells[number].String()
		number += 1
		user.ExaData.BarrelChest = row.Cells[number].String()
		number += 1
		user.ExaData.BreathSound = row.Cells[number].String()
		number += 1
		user.ExaData.Rales = row.Cells[number].String()
		number += 1
		user.ExaData.HeartRate = row.Cells[number].String()
		number += 1
		user.ExaData.Rhythm = row.Cells[number].String()
		number += 1
		user.ExaData.HeartMurmur = row.Cells[number].String()
		number += 1
		user.ExaData.AbdominAltend = row.Cells[number].String()
		number += 1
		user.ExaData.AdbominAlmass = row.Cells[number].String()
		number += 1
		user.ExaData.LiverBig = row.Cells[number].String()
		number += 1
		user.ExaData.Splenomegaly = row.Cells[number].String()
		number += 1
		user.ExaData.Dullness = row.Cells[number].String()
		number += 1
		user.ExaData.Edema = row.Cells[number].String()
		number += 1
		user.ExaData.FootPulse = row.Cells[number].String()
		number += 1
		user.ExaData.Dre = row.Cells[number].String()
		number += 1
		user.ExaData.Breast = row.Cells[number].String()
		number += 1
		user.ExaData.Vulva = row.Cells[number].String()
		number += 1
		user.ExaData.Vaginal = row.Cells[number].String()
		number += 1
		user.ExaData.Cervix = row.Cells[number].String()
		number += 1
		user.ExaData.Palace = row.Cells[number].String()
		number += 1
		user.ExaData.Attachment = row.Cells[number].String()
		number += 1
		user.ExaData.Tjother = row.Cells[number].String()

		//aeData
		number += 1
		user.AeData.Hgb = row.Cells[number].String()
		number += 1
		user.AeData.Wbc = row.Cells[number].String()
		number += 1
		user.AeData.Platelet = row.Cells[number].String()
		number += 1
		user.AeData.BloodOther = row.Cells[number].String()
		number += 1
		user.AeData.Proteinuria = row.Cells[number].String()
		number += 1
		user.AeData.Glu = row.Cells[number].String()
		number += 1
		user.AeData.Dka = row.Cells[number].String()
		number += 1
		user.AeData.Oc = row.Cells[number].String()
		number += 1
		user.AeData.UrineOther = row.Cells[number].String()
		number += 1
		user.AeData.Fbs = row.Cells[number].String() //78
		user.AeData.Fbs2 = ""
		number += 1
		user.AeData.Ecg = row.Cells[number].String()
		number += 1
		user.AeData.Malb = row.Cells[number].String()
		number += 1
		user.AeData.Fob = row.Cells[number].String()
		number += 1
		user.AeData.Hba1C = row.Cells[number].String()
		number += 1
		user.AeData.Hbsag = row.Cells[number].String()
		number += 1
		user.AeData.Alt = row.Cells[number].String()
		number += 1
		user.AeData.Ast = row.Cells[number].String()
		number += 1
		user.AeData.Alb = row.Cells[number].String()
		number += 1
		user.AeData.Tbil = row.Cells[number].String()
		number += 1
		user.AeData.Dbil = row.Cells[number].String()
		number += 1
		user.AeData.Cr = row.Cells[number].String()
		number += 1
		user.AeData.Bun = row.Cells[number].String()
		number += 1
		user.AeData.Kalemia = row.Cells[number].String()
		number += 1
		user.AeData.Natremia = row.Cells[number].String()
		number += 1
		user.AeData.Tc = row.Cells[number].String()
		number += 1
		user.AeData.Tg = row.Cells[number].String()
		number += 1
		user.AeData.Ldl = row.Cells[number].String()
		number += 1
		user.AeData.Hdl = row.Cells[number].String()
		number += 1
		user.AeData.X = row.Cells[number].String()
		number += 1
		user.AeData.B = row.Cells[number].String()
		number += 1
		user.AeData.Ps = row.Cells[number].String()
		number += 1
		user.AeData.FuOther = row.Cells[number].String()

		// hcData
		number += 1
		user.HcData.CerebrovascularDiseases = row.Cells[number].String()
		number += 1
		user.HcData.HeartDisease = row.Cells[number].String()
		number += 1
		user.HcData.KidneyDiseases = row.Cells[number].String()
		number += 1
		user.HcData.VascularDisease = row.Cells[number].String()
		number += 1
		user.HcData.EyeDiseases = row.Cells[number].String()
		number += 1
		user.HcData.NeurologicalDiseases = row.Cells[number].String()
		number += 1
		user.HcData.OtherDiseasesone = row.Cells[number].String() // 111
		number += 1
		if row.Cells[number].String() == "2" {
			user.HcData.InhospitalFlag = "n"
		} else {
			user.HcData.InhospitalFlag = "y"
		}
		number += 1
		if row.Cells[number].String() == "2" {
			user.HcData.InfamilybedFlag = "n"
		} else {
			user.HcData.InfamilybedFlag = "y"
		}
		number += 1
		if row.Cells[number].String() == "2" {
			user.HcData.MedicineFlag = "n"
		} else {
			user.HcData.MedicineFlag = "y"
		}
		number += 1
		if row.Cells[number].String() == "2" {
			user.HcData.NonimmuneFlag = "n"
		} else {
			user.HcData.NonimmuneFlag = "y"
		}

		// HaData
		number += 1
		user.HaData.Abnormality = row.Cells[number].String() // 116
		number += 1
		user.HaData.Abnormality1 = row.Cells[number].String() // 117
		number += 1
		user.HaData.Abnormality2 = row.Cells[number].String() // 118
		number += 1
		user.HaData.Abnormality3 = row.Cells[number].String() // 119
		number += 1
		user.HaData.Abnormality4 = row.Cells[number].String() // 119
		number += 1

		user.HaData.Mana = row.Cells[number].String()
		number += 1
		user.HaData.RiskfactorsControl = row.Cells[number].String()
		number += 1
		user.HaData.TargetWeight = row.Cells[number].String()
		number += 1
		user.HaData.Vaccine = row.Cells[number].String()
		number += 1
		user.HaData.PjOther = row.Cells[number].String()

		users = append(users, user)
	}
	return users, nil
}
