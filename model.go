package main

import "fmt"

type StringPair struct {
	Key   string
	Value string
}

type TagInfo struct {
	Start      int
	End        int
	CurrentTag []StringPair
	InjectTag  []StringPair
}

func (t TagInfo) ToTag() string {
	existMap := map[string]bool{}
	tag := formatKeyValuePair(t.InjectTag, existMap) + formatKeyValuePair(t.CurrentTag, existMap)
	return tag[1:]
}

func formatKeyValuePair(tags []StringPair, existMap map[string]bool) string {
	tag := ""
	for _, v := range tags {
		if _, ok := existMap[v.Key]; !ok {
			tag += fmt.Sprintf(" %v:%v", v.Key, v.Value)
			existMap[v.Key] = true
		}
	}

	return tag
}
