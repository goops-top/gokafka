package modules

import (
	"fmt"
	_ "os"
	_ "os/signal"
	_ "syscall"
	"testing"
)

func TestProducerMsg(t *testing.T) {

	producerApi := NewProducerApi([]string{"172.29.203.62:9092"})

	defer producerApi.Close()

	c := producerApi.PutFromString("test-push", "test a message")

	fmt.Println(c)
}
