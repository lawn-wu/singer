package utils

import (
	"math"
	"net"
	"singer/global"
	"strconv"
	"strings"
)

// InterfaceToInt
// @Description: 空接口类型转int
// @param t1
// @return int
func InterfaceToInt(t1 interface{}) int {
	var t2 int
	switch t1.(type) {
	case uint:
		t2 = int(t1.(uint))
	case int8:
		t2 = int(t1.(int8))
	case uint8:
		t2 = int(t1.(uint8))
	case int16:
		t2 = int(t1.(int16))
	case uint16:
		t2 = int(t1.(uint16))
	case int32:
		t2 = int(t1.(int32))
	case uint32:
		t2 = int(t1.(uint32))
	case int64:
		t2 = int(t1.(int64))
	case uint64:
		t2 = int(t1.(uint64))
	case float32:
		t2 = int(t1.(float32))
	case float64:
		t2 = int(t1.(float64))
	case string:
		t2, _ = strconv.Atoi(t1.(string))
	default:
		t2 = t1.(int)
	}
	return t2
}

// Unset [T int | string]
// @Description: 删除切片元素（根据key）
// @param slice
// @param index
// @return []T
func Unset[T int | string](slice []T, index int) []T {
	slice = append(slice[:index], slice[index+1:]...) // 删除中间1个元素
	return slice
}

// UnsetForValue [T int | string]
// @Description: 删除切片元素（根据value）
// @param slice
// @param value
// @return []T
func UnsetForValue[T int | string](slice []T, value T) []T {
	// 查找元素处于切片位置
	index := SliceSearch(slice, value)
	if index > -1 {
		Unset(slice, index)
	}
	return slice
}

// SliceDiff [T int | string]
// @Description: 查找两个切片的差集
// @param set1
// @param set2
// @return []T
func SliceDiff[T int | string](set1 []T, set2 []T) []T {
	diff := make([]T, 0)
	for _, v := range set1 {
		if !InSlice(set2, v) {
			diff = append(diff, v)
		}
	}

	for _, v := range set2 {
		if !InSlice(set1, v) {
			diff = append(diff, v)
		}
	}
	return diff
}

// InSlice [T int | string]
// @Description: 查找元素是否存在于切片
// @param array
// @param item
// @return bool
func InSlice[T int | string](array []T, item T) bool {
	return SliceSearch(array, item) >= 0
}

// SliceSearch [T int | string]
// @Description: 查找元素处于切片位置
// @param array
// @param item
// @return index
func SliceSearch[T int | string](array []T, item T) (index int) {
	index = -1
	for idx, value := range array {
		if value == item {
			index = idx
			break
		}
	}
	return
}

// GetIP
// @Description: 获取ip
// @return string
func GetIP() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addresses {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}

// MinInt
// @Description: 比较两个int大小，取较小
// @param x
// @param y
// @return int
func MinInt(x, y int) int {
	min := math.Min(float64(x), float64(y))
	return int(min)
}

// MaxInt
// @Description: 比较两个int大小，取较大
// @param x
// @param y
// @return int
func MaxInt(x, y int) int {
	max := math.Max(float64(x), float64(y))
	return int(max)
}

// PathFormat
// @Description: 格式化资源文件地址
// @param path
// @return string
func PathFormat(path string) string {
	if path == "" {
		return path
	}
	if strings.HasPrefix(path, "http") {
		return path
	}
	if path[0] != '/' {
		path = "/" + path
	}
	return global.Config.System.Domain + path
}

// StringSearch
// @Description: 查找某个可分割的规律字符串中是否存在相应字符串
// @param s
// @param substr
// @param sep
// @return bool
func StringSearch(s string, substr string, sep string) bool {
	if sep == "" {
		return strings.Contains(s, substr)
	}

	rawArray := strings.Split(s, sep)
	return InSlice(rawArray, substr)
}
