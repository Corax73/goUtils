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
	reqData.Cookies = make(map[string]string)
	headersSlice := make([]string, 0)
	parsedData := make(map[string]string)
	strSlice := strings.Split(str, " ")
	usedIndexes := make([]int, 0)
	for i, v := range strSlice {
		v = strings.Trim(v, `"`)
		v = strings.Trim(v, `'`)
		if !slices.Contains(usedIndexes, i) {
			if slices.Contains(validKeys, v) {
				if slices.Contains(requestMethods, v) {
					parsedData[v] = ""
					continue
				}

				if slices.Contains(headerKeys, v) {
					match, _ := regexp.MatchString(":$", strSlice[i+1])
					if match {
						headersSlice = append(
							headersSlice,
							goutils.ConcatSlice([]string{
								strings.Trim(strings.Trim(strSlice[i+1], `"`), `'`),
								strings.Trim(strings.Trim(strSlice[i+2], `"`), `'`),
							}),
						)
						usedIndexes = append(usedIndexes, i+2)
					} else {
						headersSlice = append(headersSlice, strings.Trim(strings.Trim(strSlice[i+1], `"`), `'`))
					}
					continue
				}

				if slices.Contains(cookieKeys, v) {
					for j := i + 1; j < len(strSlice); j++ {
						strSliceCookie := strings.Split(strSlice[j], "=")
						if len(strSliceCookie) > 1 {
							key := strSliceCookie[0]
							key = strings.Trim(strings.Trim(key, `"`), `'`)
							val := strSliceCookie[1]
							val = strings.Trim(strings.Trim(strings.Trim(val, `"`), `'`), `;`)
							reqData.Cookies[key] = val
						}
						usedIndexes = append(usedIndexes, j)
						matchByDouble, _ := regexp.MatchString(`"$`, strSlice[j])
						matchBySingle, _ := regexp.MatchString(`'$`, strSlice[j])
						if matchByDouble || matchBySingle {
							break
						}
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
		if slices.Contains(dataKeys, k) {
			reqData.Data = strings.Trim(strings.Trim(v, "'"), `"`)
		}
	}
	for _, v := range headersSlice {
		strSliceHeader := strings.Split(v, ":")
		if len(strSliceHeader) > 1 {
			if len(strSliceHeader) > 2 {
				match, _ := regexp.MatchString(`^(http)`, strSliceHeader[1])
				if match {
					reqData.Headers[strSliceHeader[0]] = goutils.ConcatSlice([]string{strSliceHeader[1], ":", goutils.ConcatSlice(strSliceHeader[2:])})
				} else {
					reqData.Headers[strSliceHeader[0]] = goutils.ConcatSlice(strSliceHeader[1:])
				}
			} else {
				reqData.Headers[strSliceHeader[0]] = strSliceHeader[1]
			}
		}
	}
	return &reqData
}
