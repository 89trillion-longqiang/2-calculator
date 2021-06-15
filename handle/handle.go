package handle

import (
	"strconv"
	"strings"

	"calculator/service"
)

func HandleGetCalculator(uri string)map[string]string{
	retMap := make(map[string]string,2)
	newExp := strings.Replace(uri,"/C1/calculator?exp=","",1)///将获取的url切分，仅保留表达式部分
	ret,err := service.Calculate(newExp)
	retString := strconv.Itoa(ret)
	if err == nil{
		retMap["condition"] = "pass"
		retMap["result"]= retString
	}else {
		retMap["condition"] = "error"
		retMap["result"]  = err.Error()
	}
	return retMap
}


