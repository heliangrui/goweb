package mqttConfig

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"go-web/core/config"
	"go-web/core/logs"
	"go.uber.org/zap"
	"time"
)

var client *MQTT.Client

func InitMQTT() {
	l := logs.Log
	//读取配置文件
	mqttConfig := config.GetMqttConfig()
	//mqtt连接
	options := MQTT.NewClientOptions()
	options.SetUsername(mqttConfig.Username)
	options.SetPassword(mqttConfig.Password)
	options.AddBroker(mqttConfig.Broker)
	options.SetClientID(options.ClientID)
	options.SetDefaultPublishHandler(messagePubHandler)
	options.SetAutoReconnect(true)
	options.SetMaxReconnectInterval(time.Duration(10) * time.Second)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectLostHandler
	options.SetKeepAlive(time.Duration(60) * time.Second)
	myClient := MQTT.NewClient(options)
	l.Info("开始连接mqtt")
	token := myClient.Connect()
	for token.Wait() && token.Error() != nil {
		l.Info("mqtt连接错误！正在重连")
		time.Sleep(time.Second * 5)
		token = myClient.Connect()
	}
	client = &myClient
	//sub(myClient)
	//Pub(myClient)
}
func sub(client MQTT.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}

func Pub(topic string, object interface{}) {
	marshal, err := json.Marshal(&object)
	if err != nil {
		logs.Log.Info("mqtt发送失败:发送主题为:topic,json转换错误！,错误信息为：object", zap.String("topic", topic), zap.String("object", err.Error()))
		return
	}
	logs.Log.Info("mqtt:发送主题为:topic,内容为：object", zap.String("topic", topic), zap.String("object", string(marshal)))
	c := *client
	token := c.Publish(topic, 0, false, marshal)
	token.Wait()
}

// 订阅回调
var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	logs.Log.Info("mqtt:接受主题为：topic，内容为：message;", zap.String("topic", msg.Topic()), zap.String("message", string(msg.Payload())))
}

// 状态从未连接/断开更改为已连接
var connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
	logs.Log.Info("mqtt连接成功")
}

// 意外断开回调
var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
	logs.Log.Info("连接失败：" + err.Error())
}
