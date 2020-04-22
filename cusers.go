package main

import (
	"errors"
	"github.com/tealeg/xlsx"
	"os"
	"path/filepath"
	"strings"
)

func (s *Server) ExcelToCUsers() ([]ChangeUserInfoRequestBody, error) {
	var users []ChangeUserInfoRequestBody
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "个人基本信息表-导入模板.xlsx") {
			users, err = s.ExcelPathToData(path)
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

func (s *Server) ExcelPathToData(path string) ([]ChangeUserInfoRequestBody, error) {
	users := make([]ChangeUserInfoRequestBody, 0)
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
		user := ChangeUserInfoRequestBody{}
		user.IDCard = row.Cells[0].String()
		user.Birthday = IdCardToBirthDay(user.IDCard)
		user.DefinePhrid = row.Cells[1].String()
		user.Address = row.Cells[2].String()
		user.PersonName = row.Cells[3].String()
		user.SexCode = row.Cells[4].String()
		user.WorkPlace = row.Cells[5].String()
		user.MobileNumber = row.Cells[6].String()
		user.Contact = row.Cells[7].String()
		user.ContactPhone = row.Cells[8].String()

		if row.Cells[9].String() == "2" {
			user.SignFlag = "n"
		} else {
			user.SignFlag = "y" //签约标志
		}

		if row.Cells[10].String() == "2" {
			user.IsAgrRegister = "n"
		} else {
			user.IsAgrRegister = "y" //农村户籍
		}
		user.IncomeSource = row.Cells[11].String()
		user.RegisteredPermanent = row.Cells[12].String()
		user.NationCode = NationCodeOfString(row.Cells[13].String()) //"01"

		user.BloodTypeCode = row.Cells[14].String()
		user.RhBloodCode = row.Cells[15].String()
		if row.Cells[16].String() == "2" {
			user.KnowFlag = "n"
		} else {
			user.KnowFlag = "y" //居民知晓
		}
		if row.Cells[17].String() == "2" {
			user.IsPovertyAlleviation = "n"
		} else {
			user.IsPovertyAlleviation = "y" //扶贫对象
		}
		user.EducationCode = row.Cells[18].String()
		user.WorkCode = "Y"
		user.MaritalStatusCode = row.Cells[20].String()
		user.InsuranceCode = "0" + row.Cells[21].String()
		user.DiseasetextCheckGm = "010" + row.Cells[22].String() // 药物过敏史22
		user.DiseasetextCheckBl = "120" + row.Cells[23].String()
		user.DiseasetextRadioJb = "0201"
		user.DiseasetextSs = "0301"
		user.DiseasetextSs0 = "无手术史"
		user.DiseasetextWs = "0601"
		user.DiseasetextWs1 = "无外伤史"
		user.DiseasetextSx = "0401"
		user.DiseasetextSx0 = "无输血史"
		user.DiseasetextCheckFq = "0701"
		user.DiseasetextCheckMQ = "0801"
		user.DiseasetextCheckXDJM = "0901"
		user.DiseasetextCheckZN = "1001"
		user.DiseasetextRedioYCBS = "0501"
		user.DiseasetextYCBS = "无遗传病史"
		user.DiseasetextCheckCJ = "110" + row.Cells[33].String()
		user.ShhjCheckCFPFSS = row.Cells[34].String()
		user.ShhjCheckRLLX = row.Cells[35].String()
		user.ShhjCheckYS = row.Cells[36].String()
		user.ShhjCheckCS = row.Cells[37].String()
		user.ShhjCheckQCL = row.Cells[38].String()
		user.DeadFlag = "n" //row.Cells[39].String()
		if row.Cells[40].String() == "2" {
			user.PersonalizedFamilyDoctorSigned = "n"
		} else {
			user.PersonalizedFamilyDoctorSigned = "y" //扶贫对象
		}
		if row.Cells[41].String() == "2" {
			user.FamilyDoctorSigned = "n"
		} else {
			user.FamilyDoctorSigned = "y" //扶贫对象
		}
		user.RegionCode = "320111001011003" //"3201110010100010001"

		user.ManaUnitID = "320111001"
		user.ManaDoctorID = "01144817"
		if row.Cells[43].String() == "2" {
			user.MasterFlag = "n"
		} else {
			user.MasterFlag = "y"
		}

		users = append(users, user)
	}
	return users, nil
}

func NationCodeOfString(name string) string {
	switch name {
	case "汉族":
		return "01"
	case "蒙古族":
		return "02"
	case "回族":
		return "03"
	case "藏族":
		return "04"
	case "维吾尔族":
		return "05"
	case "苗族":
		return "06"
	case "彝族":
		return "07"
	case "壮族":
		return "08"
	case "布依族":
		return "09"
	case "朝鲜族":
		return "10"
	case "满族":
		return "11"
	case "侗族":
		return "12"
	case "瑶族":
		return "13"
	case "白族":
		return "14"
	case "土家族":
		return "15"
	case "哈尼族":
		return "16"
	case "哈萨克族":
		return "17"
	case "傣族":
		return "18"
	case "黎族":
		return "19"
	case "傈僳族":
		return "20"
	case "佤族":
		return "21"
	case "畲族":
		return "22"
	case "高山族":
		return "23"
	case "拉祜族":
		return "24"
	case "水族":
		return "25"
	case "东乡族":
		return "26"
	case "纳西族":
		return "27"
	case "景颇族":
		return "28"
	case "柯尔克孜族":
		return "29"
	case "土族":
		return "30"
	case "达斡尔族":
		return "31"
	case "仫佬族":
		return "32"
	case "羌族":
		return "33"
	case "布朗族":
		return "34"
	case "撒拉族":
		return "35"
	case "毛南族":
		return "36"
	case "仡佬族":
		return "37"
	case "锡伯族":
		return "38"
	case "阿昌族":
		return "39"
	case "普米族":
		return "40"
	case "塔吉克族":
		return "41"
	case "怒族":
		return "42"
	case "乌孜别克族":
		return "43"
	case "俄罗斯族":
		return "44"
	case "鄂温克族":
		return "45"
	case "德昂族":
		return "46"
	case "保安族":
		return "47"
	case "裕固族":
		return "48"
	case "京族":
		return "49"
	case "塔塔尔族":
		return "50"
	case "独龙族":
		return "51"
	case "鄂伦春族":
		return "52"
	case "赫哲族":
		return "53"
	case "门巴族":
		return "54"
	case "珞巴族":
		return "55"
	case "基诺族":
		return "56"
	case "其他":
		return "57"
	}
	return "01"
}
