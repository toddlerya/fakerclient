package main

import (
	"fakerclient/rules"
	"fmt"
)

func main() {
	//rules.ReadGobData()
	//ss := rules.MapOfChinaAreaData()
	//fmt.Println("areaCode length", len(ss["areaCode"]))
	//fmt.Println("shortAreaCode length", len(ss["shortAreaCode"]))
	//fmt.Println("zipCode length", len(ss["zipCode"]))
	//fmt.Println(rules.ValidateAreaCode("970881"))
	fmt.Println(rules.ValidateStandardDate("20199395"))
	fmt.Println(rules.ValidateChinaIdCard("370881199206080314"))
}
