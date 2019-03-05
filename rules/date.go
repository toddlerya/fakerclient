package rules

import "github.com/gogf/gf/g/util/gvalid"

// 校验标准格式的日期字符串yyyymmdd
func ValidateStandardDate(dateVal string) bool {
	var legal bool
	rule := "date:20190305"
	if m := gvalid.Check(dateVal, rule, nil); m != nil {
		legal = false
	} else {
		legal = true
	}
	return legal
}
