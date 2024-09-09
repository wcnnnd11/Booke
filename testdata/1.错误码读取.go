package main

import (
	"GVB_server/models/res"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

const flie = "models/res/error_code.json"

type ErrMap map[res.ErrorCode]string

func main() {
	byteData, err := os.ReadFile(flie)
	if err != nil {
		logrus.Error(err)
		return
	}
	var errMap = ErrMap{}
	err = json.Unmarshal(byteData, &errMap)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(errMap)
	fmt.Println(errMap[res.SettingsError])
}
