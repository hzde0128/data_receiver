package models

import (
	"encoding/json"
	"github.com/hzde0128/data_receiver/common/kafkaclient"
	. "github.com/hzde0128/data_receiver/common/logger"
	"github.com/hzde0128/data_receiver/initializer/runconfig"
)

type RealTimeDataType int

// 实时数据类型
const (
	REALTIME_TYPE_ALARM        RealTimeDataType = iota // 实时告警
	REALTIME_TYPE_PERSON_TRIP                          // 人行记录
	REALTIME_TYPE_VEHICLE_TRIP                         // 车行记录
)

var RealTimeTypeDescMap = map[RealTimeDataType]string{
	REALTIME_TYPE_ALARM:        "alarm",
	REALTIME_TYPE_PERSON_TRIP:  "person_trip",
	REALTIME_TYPE_VEHICLE_TRIP: "car_trip",
}

var RealTimeTopicsMap map[RealTimeDataType]string

type RealTimeData struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

var kafkaProduct *kafkaclient.HashProductGroup

func Initialize(zookeeper []string) error {
	RealTimeTopicsMap = map[RealTimeDataType]string{
		REALTIME_TYPE_ALARM:        runconfig.Config.Kafka.AlarmTopics,
		REALTIME_TYPE_PERSON_TRIP:  runconfig.Config.Kafka.PersonTripTopics,
		REALTIME_TYPE_VEHICLE_TRIP: runconfig.Config.Kafka.VehicleTripTopics,
	}

	var err error
	kafkaProduct, err = kafkaclient.NewHashProductGroup(zookeeper)
	if err != nil {
		Log.Error("%s", err)
	}

	return err
}

func PushRealTimeData(dataType RealTimeDataType, key string, value map[string]interface{}) error {
	var realTimeData RealTimeData
	realTimeData.Type = RealTimeTypeDescMap[dataType]
	realTimeData.Data = value

	topics := RealTimeTopicsMap[dataType]

	var err error
	payload, err := json.Marshal(realTimeData)
	if err != nil {
		Log.Error("json.Marshal request error: %s", err)
		return err
	}

	Log.Info("push to topic %s, data: %s", topics, payload)
	err = kafkaProduct.PushMsg(topics, key, payload)
	if err != nil {
		Log.Error("push msg %s error: %s", key, err)
	} else {
		Log.Info("push msg %s success", key)
	}

	return err
}
