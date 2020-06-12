package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Base struct {
	ID        string     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"-"`
	DeletedAt *time.Time `sql:"index;comment:'软删除时间'" json:"-"`
}

type UserModel struct {
	Base
	AreaCode *string `json:"areaCode" sql:"comment:'社区code'"`
	AreaName string  `json:"areaName" sql:"-"`

	Name          string     `json:"name" sql:"comment:'姓名'"`
	Sex           string     `json:"sex"`
	IdCard        string     `json:"idCard" sql:"comment:'身份证'"`
	Phone         string     `json:"phone" sql:"comment:'电话'"`
	Address       string     `json:"address" sql:"comment:'现住地址'"`
	Disease       string     `json:"disease"`
	Blood         string     `json:"blood"`
	Note          string     `json:"note"`
	Diseases      []*Disease `json:"diseases"`
	Pic           string     `json:"pic" sql:"comment:'照片'"`
	PicPath       string     `gorm:"-" json:"picPath"`
	Ethnic        string     `json:"ethnic" sql:"comment:'名族'"`
	Census        uint8      `json:"census" sql:"comment:'户籍类型 1户籍 2非户籍'"`
	CensusAddress string     `json:"censusAddress" sql:"comment:'户籍地址'"`

	//更多信息
	Employer      string `json:"employer" sql:"comment:'工作单位'"`
	ContactName   string `json:"contactName" sql:"comment:'联系人'"`
	ContactPhone  string `json:"contactPhone" sql:"comment:'联系电话'"`
	BloodType     uint8  `json:"bloodType" sql:"comment:'血型1A 2B 3O 4AB 5不详'"`
	BloodRh       uint8  `json:"bloodRh" sql:"comment:'1阳 2阴 3不详'"`
	EducationType uint8  `json:"educationType" sql:"comment:'1研究生 2大学本科 3专科 4中专 5技工学校 6高中 7初中 8小学 9文盲 10不详'"`
	JobType       uint8  `json:"jobType" sql:"comment:'1国家 2专业 3办事 4商业 5农 6生产 7军人 8不便 9无'"`
	MarriageType  uint8  `json:"marriageType" sql:"comment:'1未婚 2已婚 3丧欧 4离婚 5未说明'"`

	MedicinePays       []*UserMedicinePayType `json:"medicinePays"`
	MedicineAllergies  []*UserMedicineAllergy `json:"medicineAllergies"`
	UserExposes        []*UserExpose          `json:"userExposes"`
	UserBeforeDiseases []*UserBeforeDisease   `json:"userBeforeDiseases"`

	//
	Operation   string `json:"operation" sql:"comment:'手术1无2有'"`
	Traumatic   string `json:"traumatic" sql:"comment:'外伤1无2有'"`
	Transfusion string `json:"transfusion" sql:"comment:'输血1无2有'"`

	FamilyFathers  []*FamilyDisease `json:"familyFathers"`
	FamilyMothers  []*FamilyDisease `json:"familyMothers"`
	FamilyBrothers []*FamilyDisease `json:"familyBrothers"`
	FamilyChildren []*FamilyDisease `json:"familyChildren"`

	//
	Inheritance string `json:"inheritance" sql:"遗传1无2有"`

	DiseaseConditions []*DiseaseCondition `json:"diseaseConditions"`

	KitchenDevice uint8 `json:"kitchenDevice" sql:"comment:'标志1无2油烟机3换气扇4烟囱'"`
	Fuel          uint8 `json:"fuel" sql:"comment:'标志1液化气2煤3天然气4沼气5材火'"`
	Water         uint8 `json:"water" sql:"comment:'标志1自来水2净化3井水4湖河水5塘水6其他'"`
	Washroom      uint8 `json:"washroom" sql:"comment:'标志1卫生2马桶3露天4一格5简易'"`
	Beast         uint8 `json:"beast" sql:"comment:'标志1无2单设3室内4室外'"`
}

type Disease struct {
	Base
	Name      string `sql:"comment:'疾病名称'" json:"name"`
	ShortName string `sql:"comment:'缩写'" json:"shortName"`
}

func (a *Disease) BeforeCreate(scope *gorm.Scope) (err error) {
	return scope.SetColumn("id", ksuid.New().String())
}

func (a *User) BeforeCreate(scope *gorm.Scope) (err error) {
	return scope.SetColumn("id", ksuid.New().String())
}

//医疗费用支付方式 1对多
type UserMedicinePayType struct {
	Base
	UserId string `sql:"comment:'用户id'" json:"userId"`
	Tag    uint32 `sql:"comment:'标志1城镇职工2城镇居民3新型4贫困5商业6全公费7全自费8其他'" json:"tag"`
	Value  string `sql:"comment:'文字'" json:"value"`
}

func (a *UserMedicinePayType) BeforeCreate(scope *gorm.Scope) (err error) {
	return scope.SetColumn("id", ksuid.New().String())
}

//药物过敏 1对多
type UserMedicineAllergy struct {
	Base
	UserId string `sql:"comment:'用户id'" json:"userId"`
	Tag    uint32 `sql:"comment:'标志1无2青霉素3磺胺4链霉素5其他'" json:"tag"`
	Value  string `sql:"comment:'文字'" json:"value"`
}

func (a *UserMedicineAllergy) BeforeCreate(scope *gorm.Scope) (err error) {
	return scope.SetColumn("id", ksuid.New().String())
}

//药物过敏 1对多
type UserExpose struct {
	Base
	UserId string `sql:"comment:'用户id'" json:"userId"`
	Tag    uint32 `sql:"comment:'标志1无2化学品3毒物4射线5其他'" json:"tag"`
	Value  string `sql:"comment:'文字'" json:"value"`
}

func (a *UserExpose) BeforeCreate(scope *gorm.Scope) (err error) {
	return scope.SetColumn("id", ksuid.New().String())
}

//疾病史 1对多
type UserBeforeDisease struct {
	Base
	UserId string `sql:"comment:'用户id'" json:"userId"`
	Tag    uint32 `sql:"comment:'标志1无2高血压3糖尿病4冠心病5慢性6恶性7闹卒8严重9肺结核10肝炎11其他12职业病13其他'" json:"tag"`
	Name   string `sql:"comment:'病名'" json:"name"`
	Time   uint64 `sql:"comment:'确诊时间'" json:"time"`
}

func (a *UserBeforeDisease) BeforeCreate(scope *gorm.Scope) (err error) {
	return scope.SetColumn("id", ksuid.New().String())
}

type FamilyDisease struct {
	Base
	UserId string `sql:"comment:'用户id'" json:"userId"`
	Tag    uint32 `sql:"comment:'标志1无2高血压3糖尿病4冠心病5慢性6恶性7闹卒8严重9肺结核10肝炎11先天畸形12其他'" json:"tag"`
	Value  string `sql:"comment:'文字'" json:"value"`
	Type   uint8  `sql:"comment:'1父亲2母亲3兄弟姐妹4子女'" json:"type"`
}

type DiseaseConditions struct {
	Base
	UserId string `sql:"comment:'用户id'" json:"userId"`
	Tag    uint32 `sql:"comment:'标志1无2视力残疾3听力残疾4言语残疾5肢体残疾6智力残疾7精神残疾8其他'" json:"tag"`
	Value  string `sql:"comment:'文字'" json:"value"`
}

func (a *FamilyDisease) BeforeCreate(scope *gorm.Scope) (err error) {
	return scope.SetColumn("id", ksuid.New().String())
}

type DiseaseCondition struct {
	Base
	UserId string `sql:"comment:'用户id'" json:"userId"`
	Tag    uint32 `sql:"comment:'标志1无2视力3听力4言语5肢体6智力7精神8其他'" json:"tag"`
	Value  string `sql:"comment:'文字'" json:"value"`
}

func (a *DiseaseCondition) BeforeCreate(scope *gorm.Scope) (err error) {
	return scope.SetColumn("id", ksuid.New().String())
}

// PullUserToUS 抓去的数据进数据库
// 请求接口
func (s *Server) PullUserToUS(e, r string) error {

	detail, err := s.UserInfoDetail(e, r)
	if err != nil {
		return err
	}

	code := "320111004036"

	registeredPermanent, _ := strconv.Atoi(detail.Body.RegisteredPermanent.Key)

	var bloodType string = detail.Body.BloodTypeCode.Key

	bloodTypeint, _ := strconv.Atoi(bloodType)

	var rhBloodType string = detail.Body.RhBloodCode.Key
	rhBloodTypeint, err := strconv.Atoi(rhBloodType)

	educationCode := detail.Body.EducationCode.Text
	var educationKey uint8 = 1

	switch educationCode {
	case "文盲或半文盲":
		educationKey = 9
	case "小学":
		educationKey = 8
	case "初中":
		educationKey = 7
	case "高中/技校/中专":
		educationKey = 6

	case "大学专科及以上":
		educationKey = 2
	case "不详":
		educationKey = 10
	}

	jobCode := detail.Body.WorkCode.Text
	var jobKey uint8 = 1
	switch jobCode {
	case "国家机关、党群组织、企业、事业单位负责人":
		jobKey = 1
	case "专业技术人员":
		jobKey = 2
	case "办事人员和有关人员":
		jobKey = 3
	case "商业、服务业人员":
		jobKey = 4
	case "农、林、牧、渔、水利业生产人员":
		jobKey = 5
	case "生产、运输设备操作人员及有关人员":
		jobKey = 6
	case "军人":
		jobKey = 7
	case "不便分类的其他从业人员":
		jobKey = 8
	}

	marriageKey, _ := strconv.Atoi(detail.Body.MaritalStatusCode.Key)

	medicinePayTag, _ := strconv.Atoi(detail.Body.IncomeSource.Key)
	insuranceCode := make([]*UserMedicinePayType, 0)
	insuranceCode = append(insuranceCode, &UserMedicinePayType{
		Tag:   uint32(medicinePayTag),
		Value: detail.Body.IncomeSource.Text,
	})

	//diseasetext_check_gm
	medicineAllergies := make([]*UserMedicineAllergy, 0)
	var medicineAllergiesTag uint32 = 0
	switch detail.Body.DiseasetextCheckGm.Text {
	case "无药物过敏史":
		medicineAllergiesTag = 1
	case "青霉素":
		medicineAllergiesTag = 2
	case "磺胺":
		medicineAllergiesTag = 3
	case "链霉素":
		medicineAllergiesTag = 4
	default:
		medicineAllergiesTag = 5
	}
	medicineAllergies = append(medicineAllergies, &UserMedicineAllergy{
		Tag:   medicineAllergiesTag,
		Value: detail.Body.DiseasetextCheckGm.Text,
	})
	diseasetextCheckBl := make([]*UserExpose, 0)
	var diseasetextCheckBlTag uint32 = 0
	switch detail.Body.DiseasetextCheckBl.Text {
	case "无暴露史":
		diseasetextCheckBlTag = 1
	case "化学品":
		diseasetextCheckBlTag = 2
	case "毒物":
		diseasetextCheckBlTag = 3
	case "射线":
		diseasetextCheckBlTag = 4
	default:
		diseasetextCheckBlTag = 5
	}
	diseasetextCheckBl = append(diseasetextCheckBl, &UserExpose{
		Tag:   diseasetextCheckBlTag,
		Value: detail.Body.DiseasetextCheckBl.Text,
	})
	familyFathers := make([]*FamilyDisease, 0)
	//爸爸
	diseasetextCheckfq := strings.Split(detail.Body.DiseasetextCheckFq.Text, ",")
	for _, d := range diseasetextCheckfq {
		var fatherTag uint32 = 1
		var fatherjb string
		fatherjb = d
		switch d {
		case "高血压":
			fatherTag = 2
		case "糖尿病":
			fatherTag = 3
		case "冠心病":
			fatherTag = 4
		case "慢性阻塞性肺疾病":
			fatherTag = 5
		case "恶性肿瘤":
			fatherTag = 6
		case "脑卒中":
			fatherTag = 7
		case "重性精神疾病":
			fatherTag = 8
		case "结核病":
			fatherTag = 9
		case "肝炎":
			fatherTag = 10
		case "先天畸形":
			fatherTag = 11
		case "其他":
			fatherTag = 12
			fatherjb = detail.Body.QtFq1
		}
		fqu := &FamilyDisease{
			Tag:   fatherTag,
			Value: fatherjb,
			Type:  1,
		}
		familyFathers = append(familyFathers, fqu)

	}
	familyMonthers := make([]*FamilyDisease, 0)
	diseasetextCheckmq := strings.Split(detail.Body.DiseasetextCheckMQ.Text, ",")
	for _, d := range diseasetextCheckmq {
		var fatherTag uint32 = 1
		var motherjb string
		motherjb = d
		switch d {
		case "高血压":
			fatherTag = 2
		case "糖尿病":
			fatherTag = 3
		case "冠心病":
			fatherTag = 4
		case "慢性阻塞性肺疾病":
			fatherTag = 5
		case "恶性肿瘤":
			fatherTag = 6
		case "脑卒中":
			fatherTag = 7
		case "重性精神疾病":
			fatherTag = 8
		case "结核病":
			fatherTag = 9
		case "肝炎":
			fatherTag = 10
		case "先天畸形":
			fatherTag = 11
		case "其他":
			fatherTag = 12
			motherjb = detail.Body.QtMq1
		}
		mqu := &FamilyDisease{
			Tag:   fatherTag,
			Value: motherjb,
			Type:  2,
		}
		familyMonthers = append(familyMonthers, mqu)
	}

	familyBrothers := make([]*FamilyDisease, 0)
	diseasetextCheckXd := strings.Split(detail.Body.DiseasetextCheckXDJM.Text, ",")
	for _, d := range diseasetextCheckXd {
		var fatherTag uint32 = 1
		var motherjb string
		motherjb = d
		switch d {
		case "高血压":
			fatherTag = 2
		case "糖尿病":
			fatherTag = 3
		case "冠心病":
			fatherTag = 4
		case "慢性阻塞性肺疾病":
			fatherTag = 5
		case "恶性肿瘤":
			fatherTag = 6
		case "脑卒中":
			fatherTag = 7
		case "重性精神疾病":
			fatherTag = 8
		case "结核病":
			fatherTag = 9
		case "肝炎":
			fatherTag = 10
		case "先天畸形":
			fatherTag = 11
		case "其他":
			fatherTag = 12
			motherjb = detail.Body.QtXdjm1
		}
		mqu := &FamilyDisease{
			Tag:   fatherTag,
			Value: motherjb,
			Type:  2,
		}
		familyBrothers = append(familyBrothers, mqu)

	}
	familyChildren := make([]*FamilyDisease, 0)
	diseasetextCheckZn := strings.Split(detail.Body.DiseasetextCheckZN.Text, ",")
	for _, d := range diseasetextCheckZn {
		var fatherTag uint32 = 1
		var motherjb string
		motherjb = d
		switch d {
		case "高血压":
			fatherTag = 2
		case "糖尿病":
			fatherTag = 3
		case "冠心病":
			fatherTag = 4
		case "慢性阻塞性肺疾病":
			fatherTag = 5
		case "恶性肿瘤":
			fatherTag = 6
		case "脑卒中":
			fatherTag = 7
		case "重性精神疾病":
			fatherTag = 8
		case "结核病":
			fatherTag = 9
		case "肝炎":
			fatherTag = 10
		case "先天畸形":
			fatherTag = 11
		case "其他":
			fatherTag = 12
			motherjb = detail.Body.QtXdjm1
		}
		mqu := &FamilyDisease{
			Tag:   fatherTag,
			Value: motherjb,
			Type:  2,
		}
		familyChildren = append(familyChildren, mqu)

	}

	diseaseConditions := make([]*DiseaseCondition, 0)
	diseasetextCheckCj := strings.Split(detail.Body.DiseasetextCheckCJ.Text, ",")
	for _, d := range diseasetextCheckCj {
		var jbname string
		var cjtype uint32 = 0
		jbname = d
		//无2345678
		switch d {
		case "无残疾":
			cjtype = 1
		case "视力残疾":
			cjtype = 2
		case "听力残疾":
			cjtype = 3
		case "言语残疾":
			cjtype = 4
		case "肢体残疾":
			cjtype = 5
		case "智力残疾":
			cjtype = 6
		case "精神残疾":
			cjtype = 7
		case "其他":
			cjtype = 8
			jbname = detail.Body.cjqkQtcj1
		default:
			cjtype = 0
		}
		cjqk := &DiseaseCondition{
			Tag:   cjtype,
			Value: jbname,
		}
		diseaseConditions = append(diseaseConditions, cjqk)
	}

	userBeforeDiseases := make([]*UserBeforeDisease, 0)
	diseasetextCheckJb := strings.Split(detail.Body.DiseasetextCheckJb.Text, ",")
	timeTemplate3 := "2006-01-02"
	var stamp uint64 = 0
	var comfirmdate string
	for _, d := range diseasetextCheckJb {
		var jbname string
		switch d {
		case "高血压":
			comfirmdate = detail.Body.ConfirmdateGxy
			var stamp1, _ = time.ParseInLocation(timeTemplate3, comfirmdate, time.Local)
			stamp = uint64(stamp1.Unix())
			jbname = d
		case "糖尿病":
			comfirmdate = detail.Body.ConfirmdateTnb
			var stamp1, _ = time.ParseInLocation(timeTemplate3, comfirmdate, time.Local)
			stamp = uint64(stamp1.Unix())
			jbname = d
		case "冠心病":
			comfirmdate = detail.Body.ConfirmdateGxb
			var stamp1, _ = time.ParseInLocation(timeTemplate3, comfirmdate, time.Local)
			stamp = uint64(stamp1.Unix())
			jbname = d
		case "慢性阻塞性肺疾病":
			comfirmdate = detail.Body.ConfirmdateMxzsxfjb
			var stamp1, _ = time.ParseInLocation(timeTemplate3, comfirmdate, time.Local)
			stamp = uint64(stamp1.Unix())
			jbname = d
		case "恶性肿瘤":
			comfirmdate = detail.Body.ConfirmdateExzl
			var stamp1, _ = time.ParseInLocation(timeTemplate3, comfirmdate, time.Local)
			stamp = uint64(stamp1.Unix())
			jbname = d
		case "脑卒中":
			comfirmdate = detail.Body.ConfirmdateNzz
			var stamp1, _ = time.ParseInLocation(timeTemplate3, comfirmdate, time.Local)
			stamp = uint64(stamp1.Unix())
			jbname = d
		case "重性精神疾病":
			comfirmdate = detail.Body.ConfirmdateZxjsjb
			var stamp1, _ = time.ParseInLocation(timeTemplate3, comfirmdate, time.Local)
			stamp = uint64(stamp1.Unix())
			jbname = d
		case "结核病":
			comfirmdate = detail.Body.ConfirmdateJhb
			var stamp1, _ = time.ParseInLocation(timeTemplate3, comfirmdate, time.Local)
			stamp = uint64(stamp1.Unix())
			jbname = d
		case "肝脏疾病":
			comfirmdate = detail.Body.ConfirmdateGzjb
			var stamp1, _ = time.ParseInLocation(timeTemplate3, comfirmdate, time.Local)
			stamp = uint64(stamp1.Unix())
			jbname = d
		case "先天畸形":
			comfirmdate = detail.Body.ConfirmdateXtjx
			var stamp1, _ = time.ParseInLocation(timeTemplate3, comfirmdate, time.Local)
			stamp = uint64(stamp1.Unix())
			jbname = d
		case "贫血":
			comfirmdate = detail.Body.ConfirmdatePx
			var stamp1, _ = time.ParseInLocation(timeTemplate3, comfirmdate, time.Local)
			stamp = uint64(stamp1.Unix())
			jbname = d
		case "肾脏疾病":
			comfirmdate = detail.Body.ConfirmdateSzjb
			var stamp1, _ = time.ParseInLocation(timeTemplate3, comfirmdate, time.Local)
			stamp = uint64(stamp1.Unix())
			jbname = d
		case "职业病":
			comfirmdate = detail.Body.ConfirmdateZyb
			var stamp1, _ = time.ParseInLocation(timeTemplate3, comfirmdate, time.Local)
			stamp = uint64(stamp1.Unix())
			jbname = detail.Body.DiseasetextZyb
		case "其他法定传染病":
			comfirmdate = detail.Body.ConfirmdateZyb
			var stamp1, _ = time.ParseInLocation(timeTemplate3, comfirmdate, time.Local)
			stamp = uint64(stamp1.Unix())
			jbname = detail.Body.DiseasetextQtfdcrb

		default:
			stamp = 0

		}
		u := &UserBeforeDisease{
			Tag:  UserBeforeDiseaseNameToTag(d),
			Name: jbname,
			Time: stamp,
		}

		userBeforeDiseases = append(userBeforeDiseases, u)

	}
	operation := detail.Body.DiseasetextSs.Text
	if operation == "无手术史" {
		operation = "无"
	}
	traumatic := detail.Body.DiseasetextWs.Text
	if traumatic == "无外伤史" {
		traumatic = "无"
	}
	transfusion := detail.Body.DiseasetextSx.Text
	if transfusion == "无输血史" {
		transfusion = "无"
	}
	inheritance := detail.Body.DiseasetextRedioYCBS.Text
	if inheritance == "无遗传病史" {
		inheritance = "无"
	}
	kitchenDevice, _ := strconv.Atoi(detail.Body.ShhjCheckCFPFSS.Key)
	fuel, _ := strconv.Atoi(detail.Body.ShhjCheckRLLX.Key)

	water, _ := strconv.Atoi(detail.Body.ShhjCheckYS.Key)

	var washroom uint8 = 0
	switch detail.Body.ShhjCheckCS.Text {
	case "卫生厕所":
		washroom = 1
	case "马桶":
		washroom = 2
	case "露天粪坑":
		washroom = 3
	case "一格或二格粪池式":
		washroom = 4
	case "简易棚厕":
		washroom = 5
	default:
		washroom = 6
	}

	var beastTag uint8 = 0
	switch detail.Body.ShhjCheckQCL.Text {
	case "无":
		beastTag = 1
	case "单设":
		beastTag = 2
	case "室内":
		beastTag = 3
	case "室外":
		beastTag = 4
	default:
		beastTag = 5
	}
	var u = UserModel{
		AreaCode:           &code,
		AreaName:           "", // 不需要复制
		Name:               detail.Body.PersonName,
		Sex:                detail.Body.SexCode.Text,
		IdCard:             detail.Body.IDCard,
		Phone:              detail.Body.PhoneNumber,
		Address:            detail.Body.Address,
		Disease:            "",
		Blood:              detail.Body.BloodTypeCode.Text,
		Note:               "",
		Diseases:           nil,
		Pic:                "",
		PicPath:            "",
		Ethnic:             detail.Body.NationCode.Text,
		Census:             uint8(registeredPermanent),
		CensusAddress:      "",
		Employer:           detail.Body.WorkPlace,
		ContactName:        detail.Body.Contact,
		ContactPhone:       detail.Body.ContactPhone,
		BloodType:          uint8(bloodTypeint),
		BloodRh:            uint8(rhBloodTypeint),
		EducationType:      educationKey,
		JobType:            jobKey,
		MarriageType:       uint8(marriageKey),
		MedicinePays:       insuranceCode,
		MedicineAllergies:  medicineAllergies,
		UserExposes:        diseasetextCheckBl,
		UserBeforeDiseases: userBeforeDiseases,
		Operation:          operation,
		Traumatic:          traumatic,
		Transfusion:        transfusion,
		FamilyFathers:      familyFathers,
		FamilyMothers:      familyMonthers,
		FamilyBrothers:     familyBrothers,
		FamilyChildren:     familyChildren,
		Inheritance:        inheritance,
		DiseaseConditions:  diseaseConditions,
		KitchenDevice:      uint8(kitchenDevice),
		Fuel:               uint8(fuel),
		Water:              uint8(water),
		Washroom:           washroom,
		Beast:              beastTag,
	}
	dd, _ := json.Marshal(u)
	url := "https://gongwei-api.julu666.com/app/user/info"
	cli := http.Client{}
	request, _ := http.NewRequest("POST", url, bytes.NewReader(dd))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	to := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOiIxUVlseEk0cDZDd2xKR2hCM3lWY2lFVGdMNkQiLCJhcmVhQ29kZSI6IjMyMDAwMDAwMDAwMCIsImV4cCI6MTU5MjAxNzA1NSwiaXNzIjoiZ29uZ3dlaSJ9.6UqaDVt6o46-0-x7t1XY-KsJAzWHIq5w5eZSDmnOcKM"
	request.Header.Add("Authorization", to) // 需要录入token
	resp, err := cli.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var result Response
	err = json.Unmarshal(d, &result)
	if err != nil {
		return err
	}
	if result.Code != 200 {
		return errors.New(result.Msg)
	}
	return nil
}

type Response struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func UserBeforeDiseaseNameToTag(name string) uint32 {
	switch name {
	case "无疾病史":
		return 1
	case "高血压":
		return 2
	case "糖尿病":
		return 3
	case "冠心病":
		return 4
	case "慢性阻塞性肺疾病":
		return 5
	case "恶性肿瘤":
		return 6
	case "脑卒中":
		return 7
	case "重性精神疾病":
		return 8
	case "结核病":
		return 9
	case "肝脏疾病":
		return 10
	case "其他法定传染病":
		return 11
	case "职业病":
		return 12
	case "其他":
		return 13
	}
	return 1
}
