package controllers

import (
	"github.com/hzde0128/data_receiver/common/errcode"
)

type TestController struct {
	baseController
}

func (t *TestController) Get() {
	t.Response(errcode.SUCCESS)
}
