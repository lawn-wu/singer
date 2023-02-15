package domain

import "net/url"

type SortedMaps []SortedMap

type SortedMap struct {
	Key   string
	Value string
}

func (r SortedMaps) Len() int { return len(r) }
func (r SortedMaps) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r SortedMaps) Less(i, j int) bool {
	return r[i].Key < r[j].Key
}

func (r SortedMaps) Get(key string) (value string) {
	for _, sortedMap := range r {
		if sortedMap.Key == key {
			value = sortedMap.Value
			break
		}
	}

	return
}

func (r SortedMaps) UrlValues() (data url.Values) {
	for _, sortedMap := range r {
		data.Set(sortedMap.Key, sortedMap.Value)
	}

	return
}
