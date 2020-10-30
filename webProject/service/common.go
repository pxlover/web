package service

import (
	"webProject/response"
)

func CommonShowValues() (*response.ShowValues, error) {
	var values = new(response.ShowValues)
	values.Table = make([]response.Item, 1)
	values.Table[0].Name = "我爱老婆啵啵啵!"
	return values, nil
}