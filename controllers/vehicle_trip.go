package controllers

import (
	"github.com/hzde0128/data_receiver/models"
)

type VehicleTripController struct {
	baseController
}

func (v *VehicleTripController) Post() {
	v.DisposeData(models.REALTIME_TYPE_VEHICLE_TRIP)
}
