package routers

import (
	"github.com/hzde0128/data_receiver/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/v1.0/data_receiver/alarm_report", &controllers.AlarmController{}, "post:Post")
	beego.Router("/api/v1.0/data_receiver/person_trip_report", &controllers.PersonTripController{}, "post:Post")
	beego.Router("/api/v1.0/data_receiver/vehicle_trip_report", &controllers.VehicleTripController{}, "post:Post")

	beego.Router("/test", &controllers.TestController{})
}
