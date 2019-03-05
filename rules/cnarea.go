package rules

import (
	"encoding/gob"
	"log"
	"os"
)

const dataPath = "data/cnarea_2017.data"

type cnareaData struct {
	Level, AreaCode, ZipCode, CityCode, Name, ShortName, MergerName, Lng, Lat string
}


func ReadGobData() []cnareaData {
	var readData []cnareaData
	file, err := os.Open(dataPath)
	if err != nil {
		log.Fatalf("can't open cnarea data file: %s", err)
	}
	dec := gob.NewDecoder(file)
	err = dec.Decode(&readData)
	if err != nil {
		log.Fatal(err)
	}
	return readData
}

// 通过map主键唯一的特性过滤重复元素, 空间换时间
func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{}  // 存放不重复主键
	for _, e := range slc{
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l{  // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

//该函数总共初始化两个变量，一个长度为0的slice，一个空map。由于slice传参是按引用传递，没有创建额外的变量。
//只是用了一个for循环，代码更简洁易懂。
//利用了map的多返回值特性。
//空struct不占内存空间，可谓巧妙。
func removeDuplicateElement(slice []string) []string {
	result := make([]string, 0, len(slice))
	temp := map[string]struct{}{}
	for _, item := range slice {
		if _, ok := temp[item]; !ok {  // 不存在则添加到map, 放入结果slice
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func MapOfChinaAreaData() map[string][]string {
	var allDataMap = make(map[string][]string)
	var areaCodeSlice []string
	var shortAreaCodeSlice []string
	var zipCodeSlice []string
	for _, row := range ReadGobData() {
		areaCode := row.AreaCode
		if len(areaCode) > 6 {
			areaCodeSlice = append(areaCodeSlice, areaCode)
			shortAreaCode := areaCode[0:6]
			shortAreaCodeSlice = append(shortAreaCodeSlice, shortAreaCode)
		}
		zipCode := row.ZipCode
		zipCodeSlice = append(zipCodeSlice, zipCode)
	}
	allDataMap["areaCode"] = removeDuplicateElement(areaCodeSlice)
	allDataMap["shortAreaCode"] = removeDuplicateElement(shortAreaCodeSlice)
	allDataMap["zipCode"] = removeDuplicateElement(zipCodeSlice)
	return allDataMap
}