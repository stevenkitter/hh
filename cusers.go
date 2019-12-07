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
		if strings.Contains(path, "个人基本信息表.xlsx") {
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
		user.NationCode = "01"
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
		user.RegionCode = "3201110010100010001"
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
