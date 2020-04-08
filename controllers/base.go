package controllers

import (
	"github.com/hzde0128/data_receiver/common/errcode"
	. "github.com/hzde0128/data_receiver/common/logger"
	"github.com/hzde0128/data_receiver/models"

	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
}

func (b *baseController) GetMapBody(data *map[string]interface{}) error {
	return json.Unmarshal(b.Ctx.Input.RequestBody, &data)
}

func (b *baseController) Response(code errcode.ECode, datas ...interface{}) {

	var data interface{}
	if datas == nil {
		data = nil
	} else {
		data = datas[0]
	}

	b.Data["json"] = map[string]interface{}{
		"code": code,
		"msg":  errcode.ResponseCode[code],
		"data": data,
	}

	b.ServeJSON()
}

func (b *baseController) DisposeData(dataType models.RealTimeDataType) {
	var data map[string]interface{}

	err := b.GetMapBody(&data)
	if err != nil {
		Log.Error("GetMapBody from request error: %s", err)
		b.Response(errcode.REQUEST_ERROR)
		return
	}

	device_id, exists := data["device_id"]
	if exists == false {
		Log.Error("key value error")
		b.Response(errcode.REQUEST_ERROR)
		return
	}

	key := fmt.Sprint(device_id)
	err = models.PushRealTimeData(dataType, key, data)
	if err != nil {
		Log.Error(err.Error())
		b.Response(errcode.SYSTEM_ERROR)
		return
	}

	b.Response(errcode.SUCCESS)
}
