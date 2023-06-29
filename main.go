package main

import (
	"context"
	dpfm_api_caller "data-platform-api-supply-chain-relationship-exconf-rmq-kube/DPFM_API_Caller"
	dpfm_api_output_formatter "data-platform-api-supply-chain-relationship-exconf-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-supply-chain-relationship-exconf-rmq-kube/config"
	"fmt"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

func main() {
	ctx := context.Background()
	l := logger.NewLogger()
	c := config.NewConf()
	db, err := database.NewMySQL(c.DB)
	if err != nil {
		l.Error(err)
		return
	}
	defer db.Close()

	rmq, err := rabbitmq.NewRabbitmqClient(c.RMQ.URL(), c.RMQ.QueueFrom(), "", nil, -1)
	if err != nil {
		l.Fatal(err.Error())
	}
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()
	for msg := range iter {
		go dataCallProcess(ctx, c, db, msg, rmq)
	}
}

func dataCallProcess(
	ctx context.Context,
	c *config.Conf,
	db *database.Mysql,
	rmqMsg rabbitmq.RabbitmqMessage,
	rmq *rabbitmq.RabbitmqClient,
) {
	defer rmqMsg.Success()
	l := logger.NewLogger()
	sessionId := getBodyHeader(rmqMsg.Data())
	l.AddHeaderInfo(map[string]interface{}{"runtime_session_id": sessionId})
	conf := dpfm_api_caller.NewExistenceConf(ctx, db, l)
	exist := conf.Conf(rmqMsg)
	//rmqMsg.Respond(exist)

	output, err := dpfm_api_output_formatter.NewOutput(rmqMsg, exist)
	if err != nil {
		l.Error(err)
		return
	}

	l.JsonParseOut(output)
	rmq.Send("data-platform-api-request-reads-cache-manager-receive-queue", output)
}

func getBodyHeader(data map[string]interface{}) string {
	id := fmt.Sprintf("%v", data["runtime_session_id"])
	return id
}
