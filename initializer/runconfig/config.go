package runconfig

import (
	"os"
	"strconv"

	"github.com/astaxie/beego"
)

type logParam struct {
	FilePath string
	Level    int
}

type kafkaParam struct {
	Zookeepers string

	AlarmTopics       string
	PersonTripTopics  string
	VehicleTripTopics string
}

type configParam struct {
	Log logParam // 日志相关参数

	Kafka kafkaParam
}

var Config = &configParam{}

func (c *configParam) Parse() error {
	c.Log.FilePath = beego.AppConfig.DefaultString("log::filepath", "./log/data_receive_gateway.log")
	c.Log.Level = beego.AppConfig.DefaultInt("log::loglevel", 7)

	server_listen_http_addr := os.Getenv("SERVER_ADDR")
	server_listen_http_port := os.Getenv("SERVER_PORT")
	if server_listen_http_addr != "" {
		beego.BConfig.Listen.HTTPAddr = server_listen_http_addr
	}
	if len(server_listen_http_port) != 0 {
		beego.BConfig.Listen.HTTPPort, _ = strconv.Atoi(server_listen_http_port)
	}

	zks := os.Getenv("ZK_SERVER_ADDRESS")
	if len(zks) == 0 {
		zks = beego.AppConfig.DefaultString("kafka::zookeepers", "localhost:2181")
	}

	alarmTopics := os.Getenv("ALARM_TOPICS")
	if len(alarmTopics) == 0 {
		alarmTopics = beego.AppConfig.String("kafka::alarmTopics")
	}

	personTripTopics := os.Getenv("PERSON_TRIP_TOPICS")
	if len(personTripTopics) == 0 {
		personTripTopics = beego.AppConfig.String("kafka::personTripTopics")
	}

	vehicleTripTopics := os.Getenv("VEHICLE_TRIP_TOPICS")
	if len(vehicleTripTopics) == 0 {
		vehicleTripTopics = beego.AppConfig.String("kafka::vehicleTripTopics")
	}

	c.Kafka.Zookeepers = zks
	c.Kafka.AlarmTopics = alarmTopics
	c.Kafka.PersonTripTopics = personTripTopics
	c.Kafka.VehicleTripTopics = vehicleTripTopics

	return nil
}
