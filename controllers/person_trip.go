package controllers

import (
	"github.com/hzde0128/data_receiver/models"
)

type PersonTripController struct {
	baseController
}

func (a *PersonTripController) Post() {
	a.DisposeData(models.REALTIME_TYPE_PERSON_TRIP)
}
