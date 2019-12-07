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
	//获取用户empiId
	userInfo, err := s.RequestUserInfo(user.IdCard)
	if err != nil {
		log.Printf("request err %v", err)
	}

	// 获取phrid
	userIds, err := s.RequestUserIds(userInfo.Body.EmpiId)
	if err != nil {
		log.Printf("request err %v", err)
	}

	// 获取体检记录
	res, err := s.RequestHealthCheckList(user.IdCard)
	if err != nil {
		log.Printf("request err %v", err)
		return
	}

	// 没有任何体检记录
	if len(res.Body) == 0 {
		// 新建
		s.NewHealthData(user, userInfo.Body.EmpiId, userIds.Ids.PhrID)
	}

	for _, item := range res.Body {
		// 根据新建时间
		cd := strings.Split(user.HcData.CheckDate, "T")
		if len(cd) == 0 {
			log.Printf("健康日期有误")
			return
		}
		today := cd[0]
		// 有这个点的时间
		if strings.Contains(item.CheckDate, today) {
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
	}
	log.Printf("request ok %v", res)
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
		if strings.Contains(path, "个人基本信息表.xlsx") {
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
		user.HcData.Symptom = row.Cells[4].String()
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
		user.AeData.LeftUp = ""
		user.AeData.LeftDown = ""
		user.AeData.RightUp = ""
		user.AeData.RightDown = ""
		user.AeData.Pharyngeal = row.Cells[37].String()
		user.AeData.LeftEye = row.Cells[38].String()
		user.AeData.RightEye = row.Cells[39].String()
		user.AeData.RecLeftEye = row.Cells[40].String()
		user.AeData.RecRightEye = row.Cells[41].String()
		user.AeData.Hearing = row.Cells[42].String()
		user.AeData.Motion = row.Cells[43].String()

		//exaData
		user.ExaData.Fundus = row.Cells[44].String()
		user.ExaData.Skin = row.Cells[45].String()
		user.ExaData.Sclera = row.Cells[46].String()
		user.ExaData.Lymphnodes = row.Cells[47].String()
		user.ExaData.BarrelChest = row.Cells[48].String()
		user.ExaData.BreathSound = row.Cells[49].String()
		user.ExaData.Rales = row.Cells[50].String()
		user.ExaData.HeartRate = row.Cells[51].String()
		user.ExaData.Rhythm = row.Cells[52].String()
		user.ExaData.HeartMurmur = row.Cells[53].String()
		user.ExaData.AbdominAltend = row.Cells[54].String()
		user.ExaData.AdbominAlmass = row.Cells[55].String()
		user.ExaData.LiverBig = row.Cells[56].String()
		user.ExaData.Splenomegaly = row.Cells[57].String()
		user.ExaData.Dullness = row.Cells[58].String()
		user.ExaData.Edema = row.Cells[59].String()
		user.ExaData.FootPulse = row.Cells[60].String()
		user.ExaData.Dre = row.Cells[61].String()
		user.ExaData.Breast = row.Cells[62].String()
		user.ExaData.Vulva = row.Cells[63].String()
		user.ExaData.Vaginal = row.Cells[64].String()
		user.ExaData.Cervix = row.Cells[65].String()
		user.ExaData.Palace = row.Cells[66].String()
		user.ExaData.Attachment = row.Cells[67].String()
		user.ExaData.Tjother = row.Cells[68].String()

		//aeData
		user.AeData.Hgb = row.Cells[69].String()
		user.AeData.Wbc = row.Cells[70].String()
		user.AeData.Platelet = row.Cells[71].String()
		user.AeData.BloodOther = row.Cells[72].String()
		user.AeData.Proteinuria = row.Cells[73].String()
		user.AeData.Glu = row.Cells[74].String()
		user.AeData.Dka = row.Cells[75].String()
		user.AeData.Oc = row.Cells[76].String()
		user.AeData.UrineOther = row.Cells[77].String()
		user.AeData.Fbs = row.Cells[78].String()
		user.AeData.Fbs2 = ""
		user.AeData.Ecg = row.Cells[79].String()
		user.AeData.Malb = row.Cells[80].String()
		user.AeData.Fob = row.Cells[81].String()
		user.AeData.Hba1C = row.Cells[82].String()
		user.AeData.Hbsag = row.Cells[83].String()
		user.AeData.Alt = row.Cells[84].String()
		user.AeData.Ast = row.Cells[85].String()
		user.AeData.Alb = row.Cells[86].String()
		user.AeData.Tbil = row.Cells[87].String()
		user.AeData.Dbil = row.Cells[88].String()
		user.AeData.Cr = row.Cells[89].String()
		user.AeData.Bun = row.Cells[90].String()
		user.AeData.Kalemia = row.Cells[91].String()
		user.AeData.Natremia = row.Cells[92].String()
		user.AeData.Tc = row.Cells[93].String()
		user.AeData.Tg = row.Cells[94].String()
		user.AeData.Ldl = row.Cells[95].String()
		user.AeData.Hdl = row.Cells[96].String()
		user.AeData.X = row.Cells[97].String()
		user.AeData.B = row.Cells[98].String()
		user.AeData.Ps = row.Cells[99].String()
		user.AeData.FuOther = row.Cells[100].String()

		// hcData
		user.HcData.CerebrovascularDiseases = row.Cells[101].String()
		user.HcData.HeartDisease = row.Cells[102].String()
		user.HcData.KidneyDiseases = row.Cells[103].String()
		user.HcData.VascularDisease = row.Cells[104].String()
		user.HcData.EyeDiseases = row.Cells[105].String()
		user.HcData.NeurologicalDiseases = row.Cells[106].String()
		user.HcData.OtherDiseasesone = row.Cells[107].String()
		if row.Cells[108].String() == "2" {
			user.HcData.InhospitalFlag = "n"
		} else {
			user.HcData.InhospitalFlag = "y"
		}
		if row.Cells[109].String() == "2" {
			user.HcData.InfamilybedFlag = "n"
		} else {
			user.HcData.InfamilybedFlag = "y"
		}
		if row.Cells[110].String() == "2" {
			user.HcData.MedicineFlag = "n"
		} else {
			user.HcData.MedicineFlag = "y"
		}
		if row.Cells[111].String() == "2" {
			user.HcData.NonimmuneFlag = "n"
		} else {
			user.HcData.NonimmuneFlag = "y"
		}

		// HaData
		user.HaData.Abnormality = row.Cells[112].String()
		user.HaData.Mana = row.Cells[113].String()
		user.HaData.RiskfactorsControl = row.Cells[114].String()

		users = append(users, user)
	}
	return users, nil
}
