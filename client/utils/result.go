package utils

import (
	"sync"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
var (
	r          *Result
	once       sync.Once
	ResultUtil = getResult()
)
func getResult() *Result {
	once.Do(func() {
		r = &Result{}
	})
	return r
}
func (r *Result) Success(data interface{}) *Result {
	r.Code = 0
	r.Msg = "success"
	r.Data = data
	return r
}
func (r *Result) Failure(msg string) *Result {
	r.Code = -1
	r.Msg = msg
	r.Data = nil
	return r
}
