package main

import (
	"errors"
	"fmt"
)

type CustomerError struct { //自定义错误，实现error接口
	code int
	msg  string
	err  error
}

func (c *CustomerError) Error() string {
	return fmt.Sprintf("code:%d,msg:%s,err_type:%s", c.code, c.msg, c.err.Error())
}

var SyncError = errors.New("the sync goroutine is wrong")
var ExchangeError = errors.New("this exchange is wrong")
var BuildError = errors.New("build occur error")

func MakeRunError(code int, msg string) error {
	if code == 0 {
		return &CustomerError{code, msg, SyncError}
	} else if code == 1 {
		return &CustomerError{code, msg, ExchangeError}
	} else {
		return &CustomerError{code, msg, BuildError}
	}
}

func main() {
	err := MakeRunError(0, "some code worng").Error()
	fmt.Println(err)

	err = MakeRunError(1, "if code worng").Error()
	fmt.Println(err)

	err = MakeRunError(2, "else code worng").Error()
	fmt.Println(err)
}
