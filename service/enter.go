package service

import (
	"GVB_server/service/image_ser"
	"GVB_server/service/user_ser"
)

type ServiceGroup struct {
	ImageService image_ser.ImageService
	UserService  user_ser.UserService
}

var ServiceApp = new(ServiceGroup)
