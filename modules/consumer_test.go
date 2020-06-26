package modules

import (
	_ "fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestConsumerMsg(t *testing.T) {

	consumerApi := NewConsumerApi([]string{"172.29.203.62:9092"})

	defer consumerApi.Close()

	c := consumerApi.ConsumerMsgFromTopics([]string{"test-push"})

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		log.Warnln("terminating: via signal")
	}
	c()

}
