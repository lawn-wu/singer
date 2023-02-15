package main

import (
	"fmt"
	"meituan-api/domain"
	"meituan-api/medicine"
	"meituan-api/request"
)

var medicineClient = medicine.NewClient(domain.Config{
	AppId:             "Your AppId",
	AppSecret:         "Your secret",
	AppPoiAccessToken: "",
	AppPoiCode:        "",
})

func main() {
	data, err := medicineClient.PoiMGet(request.PoiMGetParam{AppPoiCodes: "Your AppPoiCode"})
	fmt.Println(data, err)
}
