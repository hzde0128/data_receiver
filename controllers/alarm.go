package controllers

import (
	"github.com/hzde0128/data_receiver/models"
)

type AlarmController struct {
	baseController
}

func (a *AlarmController) Post() {
	a.DisposeData(models.REALTIME_TYPE_ALARM)
}
