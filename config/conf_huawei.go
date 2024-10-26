package config

type HuaWei struct {
	Enable    bool    `json:"enable" yaml:"enable"` // 是否启用华为云存储
	AccessKey string  `json:"access_key" yaml:"access_key"`
	SecretKey string  `json:"secret_key" yaml:"secret_key"`
	Bucket    string  `json:"bucket" yaml:"bucket"`       // 存储桶的名字
	EndPoint  string  `json:"end_point" yaml:"end_point"` // 终端节点
	Size      float64 `json:"size" yaml:"size"`           // 存储的大小限制，单位是MB
	Prefix    string  `json:"prefix" yaml:"prefix"`
}
