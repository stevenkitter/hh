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

func (s *Server) ScanFiles() error {
	var users []*User
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {

		if strings.Contains(path, "名单.xlsx") {
			users, err = s.FileToData(path)
			if err != nil {
				return err
			}
		}
		if strings.Contains(path, "结果表.xlsx") {
			for _, user := range users {
				err := s.Result(path, user)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	//TODO 根据用户录入信息
	for _, u := range users {
		s.AddUserInfo(*u)
	}
	return err
}

func (s *Server) AddUserInfo(user User) {
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
		s.NewData(user, userInfo.Body.EmpiId, userIds.Ids.PhrID)
	}

	for _, item := range res.Body {
		// 根据新建时间
		today := fmt.Sprintf("%d-%02d-%02d", time.Now().Year(), time.Now().Month(), time.Now().Day())
		// 有这个点的时间
		if strings.Contains(item.CheckDate, today) {
			_, err := s.RequestHealthCheckDetail(item.HealthCheck)
			if err != nil {
				log.Printf("request err %v", err)
			}
			//changeRes, err := s.ChangeHealthCheckRequest(resDetail.Body, user)
			//if err != nil {
			//	log.Printf("request err %v", err)
			//}
			//log.Printf("resDetail %d", changeRes.Code)
		} else {
			s.NewData(user, userInfo.Body.EmpiId, userIds.Ids.PhrID)
		}
	}
	log.Printf("request ok %v", res)
}

func (s *Server) NewData(user User, empiId, phrId string) {
	/*
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
	*/
	//changeRes, err := s.ChangeHealthCheckRequest(detail, user)
	//if err != nil {
	//	log.Printf("request err %v", err)
	//}
	//log.Printf("resDetail %d", changeRes.Code)
}

/**
 *  TODO 安装git
 *  Set GO111MODULE on
 *  Set GOPROXY https://goproxy.io
 *  go mod tidy
 */
func (s *Server) FileToData(path string) ([]*User, error) {
	users := make([]*User, 0)
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
		user := &User{}
		user.Code = row.Cells[0].String()
		user.HealthCode = row.Cells[1].String()
		user.Name = row.Cells[2].String()
		user.Sex = row.Cells[3].String()
		user.Birthday = row.Cells[4].String()
		user.IdCard = row.Cells[6].String()
		users = append(users, user)
	}
	return users, nil
}

func (s *Server) Result(path string, user *User) error {
	xl, err := xlsx.OpenFile(path)
	if err != nil {
		return err
	}
	if len(xl.Sheets) <= 0 {
		return errors.New("sheet number is wrong")
	}
	manSheet := xl.Sheets[0]
	for index, row := range manSheet.Rows {
		if index == 0 {
			continue
		}
		code := row.Cells[0].String()
		if code == user.Code {
			user.Constriction = row.Cells[1].String()
			user.Diastolic = row.Cells[2].String()
			user.Constriction_L = row.Cells[3].String()
			user.Diastolic_L = row.Cells[4].String()
			user.Height = row.Cells[5].String()
			user.Weight = row.Cells[6].String()
			user.Waistline = row.Cells[7].String()
			user.WehtherSmoke = row.Cells[8].String()
			user.IsDrink = row.Cells[9].String()
			user.Pulse = row.Cells[10].String()
			user.HeartMurmurDesc = row.Cells[11].String()
			if strings.Contains(row.Cells[12].String(), "正常") {
				user.HeartMurmur = "0"
			} else {
				user.HeartMurmur = "1"
			}
			user.BText = row.Cells[13].String()
			if strings.Contains(row.Cells[13].String(), "正常") {
				user.B = "0"
			} else {
				user.B = "1"
			}
			user.Hgb = row.Cells[21].String()
			user.Wbc = row.Cells[24].String()
			user.Platelet = row.Cells[22].String()
			user.Alt = row.Cells[25].String()
			user.Ast = row.Cells[26].String()
			user.Tbil = row.Cells[32].String()
			user.Cr = row.Cells[28].String()
			user.Bun = row.Cells[34].String()
			user.Tc = row.Cells[27].String()
			user.Tg = row.Cells[33].String()
			user.Ldl = row.Cells[31].String()
			user.Hdl = row.Cells[30].String()
		}

	}
	return nil
}
