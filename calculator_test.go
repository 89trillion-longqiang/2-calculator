package main

import (
	"calculator/service"
	"fmt"
	"log"
	"testing"
)

func Test_Calculate(t *testing.T) {

	exp := "3+3+3"
	ret, err := service.Calculate(exp)

	if ret == 9 && err == nil {
		fmt.Println("Test_Calculate pass")
	} else {
		log.Fatal("Test_Calculate error")
	}
}
