package config

type Upload struct {
	Size int    `yaml:"size" json:"size"` //图上上传的大小
	Path string `yaml:"path" json:"path"` //图上上传的目录
}
