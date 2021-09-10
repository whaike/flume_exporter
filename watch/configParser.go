package watch

import (
	"io/ioutil"
	"regexp"
	"strings"
)

type Conf struct {
	name string
	detail string
}

func (c Conf) Detail() string{
	return c.detail
}

func (c Conf) ParseTopic() string {
	reg := regexp.MustCompile("kafka.topics =(?P<topic>.+)")
	ps := reg.FindAllStringSubmatch(c.detail, -1)
	for _,p := range ps{
		if p[0] != ""{
			return strings.TrimSpace(p[1])
		}
		break
	}
	return ""
}

func (c Conf) ParseGroup() string {
	reg := regexp.MustCompile("kafka.consumer.group.id =(?P<group>.+)")
	ps := reg.FindAllStringSubmatch(c.detail, -1)
	for _,p := range ps{
		if p[0] != ""{
			return strings.TrimSpace(p[1])
		}
		break
	}
	return ""
}

func (c Conf) ParsePath() string {
	reg := regexp.MustCompile("hdfs.path =(?P<path>.+)")
	ps := reg.FindAllStringSubmatch(c.detail, -1)
	for _,p := range ps{
		if p[0] != ""{
			return strings.TrimSpace(p[1])
		}
		break
	}
	return ""
}

func (c Conf) ParseLogExample() string {
	prefix := ""
	suffix := ""
	reg := regexp.MustCompile("hdfs.filePrefix =(.+)")
	ps := reg.FindAllStringSubmatch(c.detail, -1)
	for _,p := range ps{
		if p[0] != ""{
			prefix = strings.TrimSpace(p[1])
			break
		}
		break
	}
	reg2 := regexp.MustCompile("hdfs.fileSuffix =(.+)")
	ps2 := reg2.FindAllStringSubmatch(c.detail, -1)
	for _,p := range ps2{
		if p[0] != ""{
			suffix = strings.TrimSpace(p[1])
			break
		}
		break
	}
	if prefix=="" || suffix== ""{
		return ""
	}
	return prefix+".xxx."+suffix
}


func NewConf(name string) Conf {
	detail,_ := ioutil.ReadFile(name)
	return Conf{
		name:   name,
		detail:  string(detail),
	}
}