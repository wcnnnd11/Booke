package desens

import "strings"

func DesensitizationEmail(email string) string {
	// 11111@qq.com 1****@qq.com
	eList := strings.Split(email, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}

func DesensitizationTel(tel string) string {
	// 158 5252 6585
	//158 **** 6585
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "****" + tel[7:]
}
