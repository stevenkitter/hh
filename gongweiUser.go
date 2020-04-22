package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"io/ioutil"
	"net/http"
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

	code := "320111004021"

	u := UserModel{
		AreaCode:           &code,
		AreaName:           "", // 不需要复制
		Name:               detail.Body.PersonName,
		Sex:                "",
		IdCard:             detail.Body.IDCard,
		Phone:              detail.Body.PhoneNumber,
		Address:            "",
		Disease:            "",
		Blood:              "",
		Note:               "",
		Diseases:           nil,
		Pic:                "",
		PicPath:            "",
		Ethnic:             "",
		Census:             0,
		CensusAddress:      "",
		Employer:           "",
		ContactName:        "",
		ContactPhone:       "",
		BloodType:          0,
		BloodRh:            0,
		EducationType:      0,
		JobType:            0,
		MarriageType:       0,
		MedicinePays:       nil,
		MedicineAllergies:  nil,
		UserExposes:        nil,
		UserBeforeDiseases: nil,
		Operation:          "",
		Traumatic:          "",
		Transfusion:        "",
		FamilyFathers:      nil,
		FamilyMothers:      nil,
		FamilyBrothers:     nil,
		FamilyChildren:     nil,
		Inheritance:        "",
		DiseaseConditions:  nil,
		KitchenDevice:      0,
		Fuel:               0,
		Water:              0,
		Washroom:           0,
		Beast:              0,
	}
	dd, _ := json.Marshal(u)
	url := "https://gongwei-api.julu666.com/app/user/info"
	cli := http.Client{}
	request, _ := http.NewRequest("POST", url, bytes.NewReader(dd))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOiIxUWUzaHBabFdDakZYMnM3dDB1RGlBQW5MRjAiLCJhcmVhQ29kZSI6IiIsImV4cCI6MTU4NzcwNDI3MiwiaXNzIjoiZ29uZ3dlaSJ9.zDoh2uOeWch-QDZEY7qZkdv1YN53n-EwGs2Lmpm6Mxc") // 需要录入token
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
