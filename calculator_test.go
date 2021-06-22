package main

import (
	"fmt"
	"log"
	"testing"

	"calculator/service"
)

func Test_StringVer(t *testing.T)  {
	exp1 := "100++100-20*20"
	exp2 := "100+100 -20*20"
	if service.StringVer(exp1)==false && service.StringVer(exp2){
		fmt.Println("Test_StringVer pass")
	}else {
		log.Fatal("Test_StringVer error")
	}
}


func Test_GetResult(t *testing.T){
	exp := "3*3+3*2-4"
	prefix := service.MixToPost(exp)
	result ,err := service.GetResult(prefix)
	if result == 11  && err == nil{
		fmt.Println("Test_GetResult pass")
	} else if err != nil {
		log.Fatal("Test_GetResult error")
	}
}

func Test_MixToPost(t *testing.T){
	postfix := []string{"3", "3", "*" ,"3", "2", "*", "+", "4", "-"}
	exp := "3*3+3*2-4"
	prefixRet := service.MixToPost(exp)
	i1 := 0
	i2 := 0
	for ; i1 < len(postfix) && i2 < len(prefixRet) ; {
		if postfix[i1] != prefixRet[i2] {
			fmt.Println(i1, i2)
			log.Fatal("Test_MixToPost error")
		}
		i1 ++
		i2 ++
	}
	fmt.Println("Test_MixToPost pass")
}