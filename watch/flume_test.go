package watch

import (
	"fmt"
	"testing"
)

func TestFuckFlumeProcess(t *testing.T) {
	x,err := FuckFlumeProcess()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(x)
	for _,f := range x{
		c := NewConf(f.ConfigName)
		topic := c.ParseTopic()
		group := c.ParseGroup()
		path := c.ParsePath()
		logExample := c.ParseLogExample()
		fmt.Println(f.Name,f.Port, topic, group, path, logExample)
	}
}