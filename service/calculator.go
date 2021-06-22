package service

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"

	"calculator/module"
)
func StringVer(exp string) bool{  ///判断字符串是否符合要求
	flag := false
	for _, char := range exp {
		if char == ' ' {
			continue
		} else if !unicode.IsDigit(rune(char)) && flag==false {
			flag = true
		}else if !unicode.IsDigit(rune(char)) && flag == true{
			return false
		}else if unicode.IsDigit(rune(char)) && flag == true{
			flag = false
		}

	}
	if  !unicode.IsDigit(rune(exp[len(exp) -1])) {
		return false
	}
	return true
}

func GetResult(postfix []string) (int ,error){
	var curStack module.Stack
	for _, v := range postfix {
		curChar := v
		if i, ero := strconv.Atoi(v); ero == nil {
			curStack.Push(strconv.Itoa(i))
		} else {
			v1, _ := curStack.Pop()
			v2, _ := curStack.Pop()
			num1, _ := strconv.Atoi(v1)
			num2, _ := strconv.Atoi(v2)
			switch curChar {
			case "+":
				curStack.Push(strconv.Itoa(num2 + num1))
			case "-":
				curStack.Push(strconv.Itoa(num2 - num1))
			case "*":
				curStack.Push(strconv.Itoa(num2 * num1))
			case "/":
				curStack.Push(strconv.Itoa(num2 / num1))
			}
		}
	}
	ret, err := curStack.Top()
	if err != nil{
		fmt.Println(err.Error())
		return 0,errors.New("stack is out of index")
	}
	result, _ := strconv.Atoi(ret)
	return result,nil
}

func MixToPost(exp string) []string {
	var curStack module.Stack
	var prefix []string
	expLen := len(exp)
	for i := 0; i < expLen; i++ {
		char := string(exp[i])
		switch char {
		case " ":
			continue
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {
				digit += string(exp[j])
			}
			prefix = append(prefix, digit)
			i = j - 1
		default:
			for !curStack.IsEmpty() {
				top, _ := curStack.Top()
				if isMore(top, char) {
					break
				}
				prefix = append(prefix, top)
				curStack.Pop()
			}
			curStack.Push(char)
		}
	}
	for !curStack.IsEmpty() {
		data, _ := curStack.Pop()
		prefix = append(prefix, data)
	}
	return prefix
}

func isMore(top, char string) bool {
	switch top {
	case "+", "-":
		if char == "*" || char == "/" {
			return true
		}
	}
	return false
}