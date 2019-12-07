package main

type User struct {
	Code       string `json:"code"`        //条码
	HealthCode string `json:"health_code"` //健康证号
	Name       string `json:"name"`        // 姓名
	Sex        string `json:"sex"`         //性别
	Birthday   string `json:"birthday"`    //名族
	IdCard     string `json:"id_card"`     //身份证
	Phone      string `json:"phone"`       //电话
	Org        string `json:"org"`         //机构
	Address    string `json:"address"`     //地址

	/**
	 *  结果表
	 */
	Constriction    string `json:"constriction"`   // 右侧收缩压
	Diastolic       string `json:"diastolic"`      // 右侧舒张压
	Constriction_L  string `json:"constriction_L"` // 左侧收缩压
	Diastolic_L     string `json:"diastolic_L"`    // 左侧舒张压
	Height          string `json:"height"`
	Weight          string `json:"weight"`
	Waistline       string `json:"waistline"`       //腰围/臀围
	WehtherSmoke    string `json:"wehtherSmoke"`    //吸烟状况
	IsDrink         string `json:"isDrink"`         //吸烟状况
	Pulse           string `json:"pulse"`           //心率
	HeartMurmur     string `json:"heartMurmur"`     //心电图
	HeartMurmurDesc string `json:"heartMurmurDesc"` // 心电图
	B               string `json:"b"`               // B 超
	BText           string `json:"bText"`

	Hgb      string `json:"hgb"`      //血红蛋白
	Wbc      string `json:"wbc"`      // 24
	Platelet string `json:"platelet"` //22
	Alt      string `json:"alt"`      //25
	Ast      string `json:"ast"`      // 26
	Tbil     string `json:"tbil"`     //32
	Cr       string `json:"cr"`       // 28
	Bun      string `json:"bun"`      //34
	Tc       string `json:"tc"`       //27
	Tg       string `json:"tg"`       // 33
	Ldl      string `json:"ldl"`      // 31
	Hdl      string `json:"hdl"`      // 30

}
