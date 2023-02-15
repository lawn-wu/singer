package util

import (
	"meituan-api/domain"
	"strconv"
	"time"
)

func ConvertSystemParamToSortedMap(systemParam *domain.SystemParam) (resultMap domain.SortedMaps) {
	resultMap = append(resultMap, domain.SortedMap{
		Key:   "timestamp",
		Value: strconv.Itoa(int(time.Now().Unix())),
	})
	resultMap = append(resultMap, domain.SortedMap{
		Key:   "app_id",
		Value: systemParam.AppId,
	})
	resultMap = append(resultMap, domain.SortedMap{
		Key:   "appSecret",
		Value: systemParam.AppSecret,
	})

	return
}
