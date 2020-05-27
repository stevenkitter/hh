package main

import (
	"time"
)

type Examination struct {
	Base
	UserId  string
	BarCode *string `sql:"comment:'体检码'" json:"barCode" gorm:"unique"`

	VisitTime uint64 `sql:"comment:'随访日期'" json:"visitTime"`
	DoctorId  string `sql:"comment:'医生id'" json:"doctorId"`

	Symptoms []*ExaminationRecordSymptom `json:"symptoms"` //

	Temperature float64 `sql:"comment:'体温'" json:"temperature"`

	PulseRate float64 `sql:"comment:'脉率'" json:"pulseRate"`

	RespiratoryRate float64 `sql:"comment:'呼吸频率'" json:"respiratoryRate"`

	LeftShrink float64 `sql:"comment:'左侧收缩'" json:"leftShrink"`
	LeftExpand float64 `sql:"comment:'左侧舒张'" json:"leftExpand"`

	RightShrink float64 `sql:"comment:'右侧收缩'" json:"rightShrink"`
	RightExpand float64 `sql:"comment:'右侧舒张'" json:"rightExpand"`

	Height    float64 `sql:"comment:'身高cm'" json:"height" json:"height"`
	Weight    float64 `sql:"comment:'当前体重'" json:"weight" json:"weight"`
	Waistline float64 `sql:"comment:'腰围'" json:"waistline"`

	HealthStatusSelf int `sql:"comment:'老年人健康状态自我评估'" json:"healthStatusSelf"`
	SelfCareAbility  int `sql:"comment:'自理能力评估'" json:"selfCareAbility"`

	CognitiveFunction *ExaminationRecordSymptom `sql:"comment:'认知功能'" json:"cognitiveFunction"`
	AffectiveState    *ExaminationRecordSymptom `sql:"comment:'情感状态'" json:"affectiveState"`

	// 生活方式
	ExerciseRate   int     `sql:"comment:'锻炼频率'" json:"exerciseRate"`
	EveryTime      float64 `sql:"comment:'每次锻炼时间'" json:"everyTime"`
	InsistTime     float64 `sql:"comment:'坚持多长了'" json:"insistTime"`
	ExerciseMethod string  `sql:"comment:'锻炼方式'" json:"exerciseMethod"`

	DietaryHabit int `sql:"comment:'饮食习惯'" json:"dietaryHabit"`

	// TODO
	DietaryHabitStr []*ExaminationRecordSymptom `sql:"comment:'饮食习惯多条'" json:"dietaryHabitStr"`

	SmokeStatus   int     `sql:"comment:'吸烟状况'" json:"smokeStatus"`
	DaySmoke      float64 `sql:"comment:'日吸烟量'" json:"daySmoke"`
	StartSmokeAge int     `sql:"comment:'开始吸烟年龄'" json:"startSmokeAge"`
	StopSmoke     int     `sql:"comment:'戒烟年龄'" json:"stopSmoke"`

	DrinkStatus int     `sql:"comment:'吸烟状态'" json:"drinkStatus"`
	DayDrink    float64 `sql:"comment:'日平均饮酒'" json:"dayDrink"`

	IsStopDrink *ExaminationRecordSymptom `sql:"comment:'是否饮酒'" json:"isStopDrink"`

	StartDrinkAge int                         `sql:"comment:'开始饮酒年龄'" json:"startDrinkAge"`
	IsDrunkenness int                         `sql:"comment:'是否醉酒'" json:"isDrunkenness"`
	DrinkTypes    []*ExaminationRecordSymptom `sql:"comment:'酒类型'" json:"drinkTypes"`

	OccupationalDiseaseTouch *ExaminationOccupationalDisease `sql:"comment:'职业病接触史'" json:"occupationalDiseaseTouch"`

	// 脏器功能
	Mouth int `sql:"comment:'口'" json:"mouth"`

	Tooth    []*ExaminationTooth `sql:"comment:'牙齿'" json:"tooth"`
	Throat   int                 `sql:"comment:'咽'" json:"throat"`
	LeftEye  float64             `sql:"comment:'左眼'" json:"leftEye"`
	RightEye float64             `sql:"comment:'右眼'" json:"rightEye"`

	LeftCorrect  float64 `sql:"comment:'左眼'" json:"leftCorrect"`
	RightCorrect float64 `sql:"comment:'右眼'" json:"rightCorrect"`

	Hearing       int `sql:"comment:'听力'" json:"hearing"`
	SportFunction int `sql:"comment:'运动功能'" json:"sportFunction"`

	// 查体
	EyeBottom *ExaminationRecordSymptom `sql:"comment:'眼底'" json:"eyeBottom"`
	Skin      *ExaminationRecordSymptom `sql:"comment:'皮肤'" json:"skin"`
	Solid     *ExaminationRecordSymptom `sql:"comment:'巩膜'" json:"solid"`
	Lymphaden *ExaminationRecordSymptom `sql:"comment:'淋巴结'" json:"lymphaden"`

	// 肺部
	TubChest   int                       `sql:"comment:'桶装胸'" json:"tubChest"`
	Breath     *ExaminationRecordSymptom `sql:"comment:'呼吸音'" json:"breath"`
	LuoYin     *ExaminationRecordSymptom `sql:"comment:'罗音'" json:"luoYin"`
	HeartRate  float64                   `sql:"comment:'心率'" json:"heartRate"`
	HeartBeat  int                       `sql:"comment:'心律'" json:"heartBeat"`
	Noise      *ExaminationRecordSymptom `sql:"comment:'杂音'" json:"noise"`
	Tenderness *ExaminationRecordSymptom `sql:"comment:'压痛'" json:"tenderness"`

	EnclosedMass   *ExaminationRecordSymptom   `sql:"comment:'包块'" json:"enclosedMass"`
	LiverBig       *ExaminationRecordSymptom   `sql:"comment:'肝大'" json:"liverBig"`
	SpleenBig      *ExaminationRecordSymptom   `sql:"comment:'脾大'" json:"spleenBig"`
	MoveMuddyVoice *ExaminationRecordSymptom   `sql:"comment:'移动式浊音'" json:"moveMuddyVoice"`
	EdemaLowerLimb int                         `sql:"comment:'下肢水肿'" json:"edemaLowerLimb"`
	Dorsalis       int                         `sql:"comment:'足背脉动'" json:"dorsalis"`
	Anus           *ExaminationRecordSymptom   `sql:"comment:'肛门'" json:"anus"`
	Breast         []*ExaminationRecordSymptom `sql:"comment:'乳腺'" json:"breast"`

	Vulva  *ExaminationRecordSymptom `sql:"comment:'外阴'" json:"vulva"`
	Vagina *ExaminationRecordSymptom `sql:"comment:'阴道'" json:"vagina"`

	Cervical    *ExaminationRecordSymptom `sql:"comment:'宫颈'" json:"cervical"`
	UterineBody *ExaminationRecordSymptom `sql:"comment:'宫体'" json:"uterineBody"`
	Enclosure   *ExaminationRecordSymptom `sql:"comment:'附件'" json:"enclosure"`

	PhysicalExamination string `sql:"comment:'查体'" json:"physicalExamination"`

	Hemoglobin float64 `sql:"comment:'血红蛋白'" json:"hemoglobin"`
	WhiteBlood float64 `sql:"comment:'白细胞'" json:"whiteBlood"`
	Platelet   float64 `sql:"comment:'血小板'" json:"platelet"`
	BloodOther string  `sql:"comment:'其他'" json:"bloodOther"`

	UrineProtein float64 `sql:"comment:'尿蛋白'" json:"urineProtein"`
	UrineSugar   float64 `sql:"comment:'尿酮体'" json:"urineSugar"`
	UrineKetone  float64 `sql:"comment:'尿酮体'" json:"urineKetone"`
	UrineBlood   float64 `sql:"comment:'尿潜血'" json:"urineBlood"`
	UrineOther   string  `sql:"comment:'其他'" json:"urineOther"`

	FastingBloodGlucose0 float64 `sql:"comment:'空腹血糖0'" json:"fastingBloodGlucose0"`
	FastingBloodGlucose1 float64 `sql:"comment:'空腹血糖1'" json:"fastingBloodGlucose1"`

	MicroalBuminuria float64 `sql:"comment:'尿微量白蛋白'" json:"microalBuminuria"`
	StoolDiving      int     `sql:"comment:'大便潜血'" json:"stoolDiving"`

	GlycosylatedHemoglobin float64 `sql:"comment:'糖化血红蛋白'" json:"glycosylatedHemoglobin"`
	Hepatitis              int     `sql:"comment:'乙型肝炎'" json:"hepatitis"`

	LiverFunction0 float64 `sql:"comment:'血清谷丙转安酶'" json:"liverFunction0"`
	LiverFunction1 float64 `sql:"comment:'血清谷草'" json:"liverFunction1"`
	LiverFunction2 float64 `sql:"comment:'白蛋白'" json:"liverFunction2"`
	LiverFunction3 float64 `sql:"comment:'总胆红素'" json:"liverFunction3"`
	LiverFunction4 float64 `sql:"comment:'结合胆红素'" json:"liverFunction4"`

	RenalFunction0 float64 `sql:"comment:'血清肌酐'" json:"renalFunction0"`
	RenalFunction1 float64 `sql:"comment:'血尿素氮'" json:"renalFunction1"`
	RenalFunction2 float64 `sql:"comment:'血钾浓度'" json:"renalFunction2"`
	RenalFunction3 float64 `sql:"comment:'血钠浓度'" json:"renalFunction3"`

	BloodFat0 float64 `sql:"comment:'总胆固醇'" json:"bloodFat0"`
	BloodFat1 float64 `sql:"comment:'甘油三脂'" json:"bloodFat1"`

	BloodLowCholesterol  string `sql:"comment:'血清低密度脂'" json:"bloodLowCholesterol"`
	BloodHighCholesterol string `sql:"comment:'血清高密度脂'" json:"bloodHighCholesterol"`

	Electrocardiogram string `sql:"comment:'心电图'" json:"electrocardiogram"`
	ChestX            string `sql:"comment:'胸部X线片'" json:"chestX"`
	BMode             string `sql:"comment:'B超'" json:"bMode"`
	CervicalPic       string `sql:"comment:'宫颈涂片'" json:"cervicalPic"`
	OtherCheck        string `sql:"comment:'其他辅助检查'" json:"otherCheck"`

	// 中医体制
	Medicine0 int `sql:"comment:'和平质'" json:"medicine0"`
	Medicine1 int `sql:"comment:'气虚质'" json:"medicine1"`
	Medicine2 int `sql:"comment:'阳虚质'" json:"medicine2"`
	Medicine3 int `sql:"comment:'阴虚质'" json:"medicine3"`
	Medicine4 int `sql:"comment:'痰湿质'" json:"medicine4"`
	Medicine5 int `sql:"comment:'湿热质'" json:"medicine5"`
	Medicine6 int `sql:"comment:'血瘀质'" json:"medicine6"`
	Medicine7 int `sql:"comment:'气郁质'" json:"medicine7"`
	Medicine8 int `sql:"comment:'特秉质'" json:"medicine8"`

	CerebrovascularDisease []*ExaminationRecordSymptom `sql:"comment:'脑血管疾病'" json:"cerebrovascularDisease"`
	KidneyDisease          []*ExaminationRecordSymptom `sql:"comment:'肾脏疾病'" json:"kidneyDisease"`
	HeartDisease           []*ExaminationRecordSymptom `sql:"comment:'心脏疾病'" json:"heartDisease"`

	BloodDisease         []*ExaminationRecordSymptom `sql:"comment:'血管疾病'" json:"bloodDisease"`
	EyeDisease           []*ExaminationRecordSymptom `sql:"comment:'眼部疾病'" json:"eyeDisease"`
	NervousSystemDisease *ExaminationRecordSymptom   `sql:"comment:'神经系统疾病'" json:"nervousSystemDisease"`
	OtherSystemDisease   *ExaminationRecordSymptom   `sql:"comment:'其他系统疾病'" json:"otherSystemDisease"`

	Hospitalization []*ExaminationHospitalization `sql:"comment:'住院史'" json:"hospitalization"`
	HospitalBed     []*ExaminationHospitalization `sql:"comment:'病床史'" json:"hospitalBed"`

	DrugSituation []*ExaminationDrugSituation `sql:"comment:'用药情况'" json:"drugSituation"`
	Vaccination   []*ExaminationVaccination   `sql:"comment:'接种'" json:"vaccination"`

	HealthAssessmentGuidance []*ExaminationRecordSymptom `sql:"comment:'健康评价指导'" json:"healthAssessmentGuidance"` //
	HealthGuidance           []*ExaminationRecordSymptom `sql:"comment:'健康指导'" json:"healthGuidance"`
	RiskFactors              []*ExaminationRecordSymptom `sql:"comment:'危险因素'" json:"riskFactors"`
}

type ExaminationRecordSymptom struct {
	Base
	RecordId string `sql:"comment:'主表记录id'" json:"recordId"`
	Tag      uint32 `sql:"comment:'标志0无症状1头痛头晕2恶心呕吐3眼花耳鸣45678 9其他'" json:"tag"`
	Value    string `sql:"comment:'文字'" json:"value"`
	Type     string `sql:"comment:'类型'"`
}

type ExaminationOccupationalDisease struct {
	Base

	RecordId                string `sql:"comment:'主表记录id'" json:"recordId"`
	Profession              string `sql:"comment:'工种'" json:"profession"` //
	WorkTime                int    `sql:"comment:'从业时间'" json:"workTime"` //
	Dust                    string `sql:"comment:'粉尘'" json:"dust"`
	DustSafeguardProcedures string `sql:"comment:'防护措施'" json:"dustSafeguardProcedures"`

	Radiate                    string `sql:"comment:'放射物质'" json:"radiate"`
	RadiateSafeguardProcedures string `sql:"comment:'放射物质 防护措施'" json:"radiateSafeguardProcedures"`

	Physics                    string `sql:"comment:'物理因素'" json:"physics"`
	PhysicsSafeguardProcedures string `sql:"comment:'物理因素 防护措施'" json:"physicsSafeguardProcedures"`

	Chemistry                    string `sql:"comment:'化学物质'" json:"chemistry"`
	ChemistrySafeguardProcedures string `sql:"comment:'化学物质 防护措施'" json:"chemistrySafeguardProcedures"`

	Other                    string `sql:"comment:'其他'" json:"other"`
	OtherSafeguardProcedures string `sql:"comment:'其他 防护措施'" json:"otherSafeguardProcedures"`
}

type ExaminationTooth struct {
	Base
	RecordId string  `sql:"comment:'主表记录id'" json:"recordId"`
	Position int     `sql:"comment:'位置 0正常 1缺 2锯 3义齿'" json:"position"`
	LTop     float64 `sql:"comment:''" json:"lTop"`
	LDown    float64 `sql:"comment:''" json:"lDown"`
	RTop     float64 `sql:"comment:''" json:"rTop"`
	RDown    float64 `sql:"comment:''" json:"rDown"`
}

type ExaminationHospitalization struct {
	Base
	RecordId     string    `sql:"comment:'主表记录id'" json:"recordId"`
	InDate       time.Time `sql:"comment:'入院'" json:"inDate"`
	OutTime      time.Time `json:"outTime"`
	Reason       string    `json:"reason"`
	Organization string    `sql:"comment:'医疗名称'" json:"organization"`
	DiseaseCode  string    `sql:"comment:'病案号'" json:"diseaseCode"`
	Type         string    `json:"type"`
}

type ExaminationDrugSituation struct {
	Base
	RecordId  string `sql:"comment:'主表记录id'" json:"recordId"`
	Name      string `sql:"comment:'药物名称'" json:"name"`
	Use       string `sql:"comment:'用法'" json:"use"`
	UseAmount string `sql:"comment:'用量'" json:"useAmount"`
	UseTime   string `sql:"comment:'用药时间'" json:"useTime"`
	UseRelyOn int    `sql:"comment:'用药依赖'" json:"useRelyOn"`
}

type ExaminationVaccination struct {
	Base
	RecordId     string    `sql:"comment:'主表记录id'" json:"recordId"`
	Name         string    `json:"name"`
	Day          time.Time `json:"day"`
	Organization string    `json:"organization"`
}
