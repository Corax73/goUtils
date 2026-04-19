package curl

import (
	"fmt"
	"slices"
	"strings"
)

func ParseCurlString(str string) {
	resp := make(map[string]string)
	strSlice := strings.Split(str, " ")
	usedIndexes := make([]int, 1, len(strSlice))
	for i, v := range strSlice {
		if !slices.Contains(usedIndexes, i) && slices.Contains(validKeys, v) {
			resp[v] = strSlice[i+1]
			usedIndexes = append(usedIndexes, i+1)
		}
	}
	fmt.Println(resp)
}
