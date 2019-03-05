package rules

import (
	"strconv"
)

// 根据前17位数字获取校验码
func getValidateNumber(number string) string {
	weight := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}        // 十七位数字本体码权重
	validate := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"} // mod11,对应校验码字符值
	sum := 0
	for i := 0; i < len(number); i++ {
		temp, _ := strconv.Atoi(string(number[i]))
		temp = temp * weight[i]
		sum = sum + temp
	}
	mode := sum % 11
	return validate[mode]
}

// 身份证证前6位区位编码
func ValidateAreaCode(areaCode string) bool {
	var legal bool
	if len(areaCode) != 6 {
		legal = false
	} else {
		shortAreaCode := MapOfChinaAreaData()["shortAreaCode"]
		for _, ele := range shortAreaCode { // 目前为3661次循环, 可以排序做个二分查找，不过没必要
			if ele == areaCode {
				legal = true
			}
		}
	}
	return legal
}

// 15位身份证转18位
// TODO 需要校验15位正确后再转18位
func IdCardFormat18(oldIdCard string) (string, bool) {
	var legal bool
	var newIdCard string
	if len(oldIdCard) == 15 {
		temp := oldIdCard[0:6] + "19" + oldIdCard[6:]
		newIdCard = temp + getValidateNumber(temp)
		legal = true
	} else {
		newIdCard = ""
		legal = false
	}
	return newIdCard, legal
}

// 校验身份证号码
func ValidateChinaIdCard(idCardVal string) bool {
	var legal bool
	length := len(idCardVal)
	switch length {
	case 15:
		// TODO 15位身份证校验
		legal = false
	case 18:
		areaCode := idCardVal[0:6]
		if !ValidateAreaCode(areaCode) {
			legal = false
		} else {
			birthDate := idCardVal[6:14]
			if !ValidateStandardDate(birthDate) {
				legal = false
			} else {
				if getValidateNumber(idCardVal[0:17]) == idCardVal[17:] {
					legal = true
				}
			}
		}
	default:
		legal = false
	}
	return legal
}



