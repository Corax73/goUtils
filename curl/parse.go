package curl

import (
	"regexp"
	"slices"
	"strings"

	goutils "github.com/Corax73/goUtils"
)

func ParseCurlString(str string) *Request {
	var reqData Request
	reqData.Headers = make(map[string]string)
	headersSlice := make([]string, 0)
	parsedData := make(map[string]string)
	strSlice := strings.Split(str, " ")
	usedIndexes := make([]int, 0)
	for i, v := range strSlice {
		if !slices.Contains(usedIndexes, i) && slices.Contains(validKeys, v) {
			if slices.Contains(requestMethods, v) {
				parsedData[v] = ""
				continue
			}

			if slices.Contains(headerKeys, v) {
				match, _ := regexp.MatchString(":$", strSlice[i+1])
				if match {
					headersSlice = append(headersSlice, goutils.ConcatSlice([]string{strSlice[i+1], strSlice[i+2]}))
					usedIndexes = append(usedIndexes, i+2)
				} else {
					headersSlice = append(headersSlice, strSlice[i+1])
				}
				continue
			}

			match, _ := regexp.MatchString(":$", strSlice[i+1])
			if match {
				parsedData[v] = goutils.ConcatSlice([]string{strSlice[i+1], strSlice[i+2]})
				usedIndexes = append(usedIndexes, i+2)
			} else {
				parsedData[v] = strSlice[i+1]
			}
			usedIndexes = append(usedIndexes, i, i+1)
		}
		if !slices.Contains(usedIndexes, i) {
			match, _ := regexp.MatchString(`^(http:\/\/|https:\/\/|\/|\/\/)`, v)
			if match {
				reqData.Url = v
			}
		}
	}
	for k, v := range parsedData {
		if slices.Contains(requestMethods, k) {
			if slices.Contains(requestMethods[:2], k) {
				reqData.Method = "GET"
			} else {
				reqData.Method = "HEAD"
			}
		}
		if k == "-X" || k == "--request" {
			reqData.Method = v
		}
	}
	for _, v := range headersSlice {
		strHeader := strings.Split(v, ":")
		reqData.Headers[strHeader[0]] = strHeader[1]
	}
	return &reqData
}
