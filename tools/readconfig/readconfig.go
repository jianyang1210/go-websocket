package readconfig

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var ConfigData config.Configer

func InitConfig() (err error) {
	lasTwoPath := map[string]bool{
		"readconfig":    true,
		"send2client":   true,
		"bind2group":    true,
		"send2group":    true,
		"getonlinelist": true,
		"register":      true,
	}

	path, _ := os.Getwd()
	if strings.Contains(path, "servers") {
		path += "/.."
	} else {
		for key := range lasTwoPath {
			if strings.Contains(path, key) {
				path += "/../.."
				break
			}
		}
	}
	ConfigData, err = config.NewConfig("ini", path+"/configs/config.ini")
	if err != nil {
		return err
	}

	cluster, err := ConfigData.Bool("common::cluster")
	if err != nil {
		return err
	}

	//如果设置了集群
	if cluster {
		//todo
		logrus.Debug(fmt.Sprintf("cluster setting %v", cluster))
	}
	return nil
}
