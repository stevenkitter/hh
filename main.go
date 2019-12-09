package main

import "log"

func main() {
	// 系统有个接口会校验是否需要加密，目前是无需要加密
	s := Server{}
	userId := "01144817"
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
	token := logResp.Body.Tokens[0]
	res, err := s.Login(token.ID, token.UserID, pwd)
	if err != nil {
		log.Fatalf("登陆失败 %v", err)
	}
	log.Printf("登陆成功, %d", res.Code)

	//导入人员信息
	// s.ImpUserData()

	//导入健康体检
	s.ImpHealth()

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
