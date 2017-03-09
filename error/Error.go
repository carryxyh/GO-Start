package main

import (
	"errors"
	"fmt"
)

type MineError struct {
	msg    string
	offset int64
}

func (e *MineError) Error() string {
	return e.msg
}

type Customerror struct {
	infoa string
	infob string
	Err   error
}

func (cerr Customerror) Error() string {
	errorinfo := fmt.Sprintf("infoa : %s , infob : %s , original err info : %s ", cerr.infoa, cerr.infob, cerr.Err.Error())
	return errorinfo
}

//函数返回自定义错误时，返回值推荐设置为error类型，而非自定义错误类型，特别需要注意的是不应预声明自定义错误类型的变量。
func Decode() *MineError { // 错误，将可能导致上层调用者err!=nil的判断永远为true。
	var err *MineError
	if 条件 {
		err = &MineError{}
	}
	return err // 错误，err永远等于非nil，导致上层调用者err!=nil的判断始终为true
}

func main() {

}
