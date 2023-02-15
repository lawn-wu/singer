package util

import (
	"meituan-api/domain"
	"reflect"
)

func GetStructField(object any, tagName string) (applicationParams []domain.SortedMap) {
	typeOfSelf := reflect.TypeOf(object)

	// 遍历结构体所有成员
	for i := 0; i < typeOfSelf.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfSelf.Field(i)

		// 通过字段名, 找到字段类型信息
		if selfType, ok := typeOfSelf.FieldByName(fieldType.Name); ok {
			// 输出成员名和tag
			if fieldType.Tag != "" && selfType.Tag.Get(tagName) != "" {
				applicationParams = append(applicationParams, domain.SortedMap{
					Key:   selfType.Tag.Get(tagName),
					Value: reflect.ValueOf(object).FieldByName(fieldType.Name).String(),
				})
			}
		}

	}

	return
}
