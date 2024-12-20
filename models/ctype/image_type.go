package ctype

import "encoding/json"

type ImageType int

const (
	Local  ImageType = 1 // 本地
	QiNiu  ImageType = 2 // 七牛
	ALi    ImageType = 3 //阿里
	HuaWei ImageType = 4 //华为
)

func (s ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageType) String() string {
	var str string
	switch s {
	case Local:
		str = "本地"
	case QiNiu:
		str = "七牛"
	case ALi:
		str = "阿里"
	case HuaWei:
		str = "华为"
	default:
		str = "其他"
	}
	return str
}
