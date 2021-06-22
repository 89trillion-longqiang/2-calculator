package handle

import (
	"strconv"
	"strings"

	"calculator/service"
)

func HandleGetCalculator(uri string)map[string]string{
	retMap := make(map[string]string,2)
	newExp := strings.Replace(uri,"/C1/calculator?exp=","",1)///将获取的url切分，仅保留表达式部分
	if service.StringVer(newExp) == false{				///验证字符串是否合法
		retMap["condition"] = "error"
		retMap["result"]  = "StringVerification error"
		return retMap
	}

	postfix := service.MixToPost(newExp)  ///转化为后缀表达式
	if postfix == nil {
		retMap["condition"] = "error"
		retMap["result"]  = "StringVerification error"
		return retMap
	}

	result ,err := service.GetResult(postfix)

	if err == nil{
		retMap["condition"] = "pass"
		retMap["result"]= strconv.Itoa(result)
	}else {
		retMap["condition"] = "error"
		retMap["result"]  = err.Error()
	}
	return retMap
}


