package main

import (
	"bytes"
	"encoding/json"
	"errors"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 系统有个接口会校验是否需要加密，目前是无需要加密
	s := Server{}
	db, _ := ConnectMysql("47.100.127.10", "Ecc8MDPsuEA9AWLk")
	s.DB = db

	t, _ := GongweiToken()
	s.GongweiToken = t

	userId := "09283221"
	pwd := "1"
	logResp, err := s.DoctorInfo(userId, pwd)
	if err != nil {
		log.Fatalf("登陆错误，请联系管理员")
	}
	for _, token := range logResp.Body.Tokens {
		log.Printf("医生的角色是： %s", token.Role.Name)
		log.Printf("医生的结构是：%s", token.ManageUnit.Name)
	}
	if len(logResp.Body.Tokens) == 0 {
		log.Fatalf("没有医生的信息无法操作")
	}
	var token DoctorInfoToken
	for _, i := range logResp.Body.Tokens {
		if i.RoleName == "责任医生" {
			token = i
		}
	}
	//token := logResp.Body.Tokens[0]
	res, err := s.Login(token.ID, token.UserID, pwd)
	if err != nil {
		log.Fatalf("登陆失败 %v", err)
	}
	log.Printf("登陆成功, %d", res.Code)

	// 导入人员信息
	// s.ImpUserData()
	//导入健康体检
	// s.ImpHealth()
	// 江浦数据进公卫系统
	// s.ExportUserInfoToGongWei()
	// 公卫健康体检进江浦系统
	err = s.impOldToTmp("2020-04-26")
	if err != nil {
		log.Printf("导入错误 %v", err)
	}
}

// ExportUserInfoToGongWei 导出到共卫系统
func (s *Server) ExportUserInfoToGongWei() error {
	theyCode := "320111001010"
	var start = 0
	list, err := s.GongweiUserInfo(theyCode+"%", start)
	if err != nil {
		return err
	}
	for _, i := range list.Body {
		log.Printf(i.PersonName)
		err := s.PullUserToUS(i.EmpiID, i.RegionCode)

		if err != nil {
			return err
		}
	}
	for len(list.Body) > 0 {
		start++
		list, err = s.GongweiUserInfo(theyCode+"%", start)
		if err != nil {
			return err
		}
		for _, i := range list.Body {
			log.Printf(i.PersonName)
			err := s.PullUserToUS(i.EmpiID, i.RegionCode)

			if err != nil {
				return err
			}
		}
	}
	return nil
}

//  导入人员信息
func (s *Server) ImpUserData() {

	users, err := s.ExcelToCUsers()
	if err != nil {
		log.Fatalf("Excel数据读取错误 %v", err)
	}
	for _, u := range users {
		person, err := s.PersonalInfo(u.IDCard)
		if err != nil {
			log.Printf("此用户信息获取错误 %v", err)
			continue
		}
		if person.Body.EmpiID == "" {
			log.Printf("此用户信息不存在，新建中...")
			resp, err := s.ChangeUserInfoRequest(u, person.Body)
			if err != nil {
				log.Printf("录入错误 身份证为：%s", u.IDCard)
				continue
			}
			if resp.Body.EmpiID == "" {
				log.Printf("录入错误 身份证为：%s", u.IDCard)
				continue
			}
			log.Printf("此用户信息已新建完成")
		} else {
			log.Printf("此用户信息已存在，修改中...")
			resp, err := s.ChangeUserInfoRequest(u, person.Body)
			if err != nil {
				log.Printf("录入错误 身份证为：%s", u.IDCard)
				continue
			}
			if resp.Body.EmpiID == "" {
				log.Printf("录入错误 身份证为：%s", u.IDCard)
				continue
			}
			log.Printf("此用户信息已修改完成")

		}
	}
}

// 导入健康信息
func (s *Server) ImpHealth() {
	users, err := s.HealthExcelToCUsers()
	if err != nil {
		log.Fatalf("Excel数据读取错误 %v", err)
	}
	for _, u := range users {
		s.AddUserHealthInfo(u)
	}
}

// 公卫系统体检进江浦
// 求打中文 擦
func (s *Server) GongweiHealthImp() {
	sql := `SELECT id FROM examinations WHERE DATE_SUB(CURDATE(), INTERVAL 7 DAY) <= DATE(updated_at) `
	rows, err := s.DB.Raw(sql).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	ids := make([]string, 0)
	for rows.Next() {
		var dest struct {
			ID string
		}
		s.DB.ScanRows(rows, &dest)
		ids = append(ids, dest.ID)
	}
	for _, d := range ids {
		detail, err := s.GetExaminationDetail(d)
		if err != nil {
			return
		}
		// 关键点 转模型
		log.Printf("detail %s", detail.ID)
	}

}

func (s *Server) GetExaminationDetail(id string) (health *Examination, err error) {
	path := "https://gongwei-api.julu666.com/app/examination/detail"
	cli := http.Client{}
	request, err := http.NewRequest("GET", path+"?id="+id, nil)
	if err != nil {
		return
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	request.Header.Add("Authorization", s.GongweiToken)
	response, err := cli.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	var em Examination
	err = json.Unmarshal(data, &em)
	if err != nil {
		return
	}
	health = &em
	return
}

// gongwei 系统的token
type LoginResponse struct {
	Code uint32    `json:"code"`
	Msg  string    `json:"msg"`
	Data TokenData `json:"data"`
}
type TokenData struct {
	Token string `json:"token"`
}

type GongweiLoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func GongweiToken() (t string, err error) {
	path := "https://gongwei-api.julu666.com/app/login"
	cli := http.Client{}
	dd := GongweiLoginRequest{Account: "14751601462", Password: "4297f44b13955235245b2497399d7a93"}
	data, _ := json.Marshal(dd)
	request, _ := http.NewRequest("POST", path, bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	resp, err := cli.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var result LoginResponse
	err = json.Unmarshal(d, &result)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = errors.New(result.Msg)
		return
	}
	t = result.Data.Token
	return
}
