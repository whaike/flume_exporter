package web

import (
	"encoding/base64"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/woozhijun/flume_exporter/watch"
	"net/http"
)

type Flumes struct {
	Topic string `json:"topic"`
	Group string `json:"group"`
	Path string `json:"path"`
	LogExample string `json:"logExample"`
}

func ConfigParse(w http.ResponseWriter, req *http.Request){
	result := make([]Flumes, 0)
	x,err := watch.FuckFlumeProcess()
	if err != nil {
		log.Error(err)
		w.Write([]byte("{}"))
	}
	for _,f := range x{
		c := watch.NewConf(f.ConfigName)
		topic := c.ParseTopic()
		group := c.ParseGroup()
		path := c.ParsePath()
		logExample := c.ParseLogExample()
		//fmt.Println(f.Name,f.Port, topic, group, path, logExample)
		result = append(result, Flumes{
			Topic:     topic,
			Group:     group,
			Path:      path,
			LogExample: logExample,
		})
	}
	w.Header().Set("content-type","text/json")
	msg,_ := json.Marshal(result)
	w.Write(msg)
}

type R struct {
	Name string
	Port string
	ConfigName string
	ConfigDetail string
}

func ConfigStr(w http.ResponseWriter, req *http.Request){
	result := make([]R, 0)
	x,err := watch.FuckFlumeProcess()
	if err != nil {
		log.Error(err)
		w.Write([]byte("{}"))
	}
	for _,f := range x{
		c := watch.NewConf(f.ConfigName)
		result = append(result, R{
			Name:         f.Name,
			Port:         f.Port,
			ConfigName:   f.ConfigName,
			ConfigDetail: base64.StdEncoding.EncodeToString([]byte(c.Detail())),
		})
	}
	w.Header().Set("content-type","text/json")
	msg,_ := json.Marshal(result)
	w.Write(msg)
}