package core

import (
	"GVB_server/config"
	"GVB_server/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"log"
)

const ConfigFile = "settings.yaml"

// InitConf  选取yaml文件配置
func InitConf() {

	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal:%v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)

	if err != nil {

		return err
	}

	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {

		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
