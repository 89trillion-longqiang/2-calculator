package service

import "fmt"

func StringVerification(s string)bool{
	a := s[len(s)-1]
	if a >= '0' && a <= '9' {
		return true
	}else {
		return false
	}
}

func getSum(nums []int) (sum int){

	for _, v := range nums { // 将nums栈中的数字逐个累加
		sum += v
	}
	return sum
}

func Calculate(s string) (ret int ,err error){
	if !StringVerification(s){
		err = fmt.Errorf("%s","exp error")
		return 0 ,err
	}

	var nums []int // 栈。解析过的数字入栈
	curNum := 0    // 当前遇到的数字
	prevOp := '+'  // 上一个遇到的运算符
	s = s + "+"    // 追加一个运算符，不然会少做最后一次运算

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' { // 如果是数字
			curNum = curNum * 10 + int(s[i] - '0') // 更新curNum
		} else if s[i] == ' ' { // 如果遇到空格，跳过
			continue
		} else { // 如果遇到加减乘除
			if prevOp == '+' { // 如果上一个遇到的运算符是加号
				nums = append(nums, curNum) // 将curNum入栈
			} else if prevOp == '-' {
				nums = append(nums, -1*curNum) // 将负的curNum入栈
			} else if prevOp == '*' {
				nums[len(nums)-1] = nums[len(nums)-1] * curNum // 更新栈顶为乘过的值
			} else if prevOp == '/' {
				if curNum == 0 {
					err = fmt.Errorf("%s","The divisor is 0")
					return 0,err
				}
				nums[len(nums)-1] = nums[len(nums)-1] / curNum // 更新栈顶为除过的值
			}else {
				err = fmt.Errorf("%s","exp error")
				return 0 ,err
			}
			prevOp = rune(s[i]) // 记录当前遇到的运算符
			curNum = 0 // curNum 重置
		}
	}

	return getSum(nums),err
}

