package thinkJson

import (
	"fmt"
	"sort"
)

func MapJsonBySortKey(params map[string]interface{}, sortKey []string) string{
	if len(sortKey) == 0 {
		for key, _ := range params {
			sortKey = append(sortKey, key)
		}
		sort.Strings(sortKey)
	}

	j := "{"
	for _, key := range sortKey {
		s := `"%s":%s,`
		j += fmt.Sprintf(s, key, string(MustMarshal(params[key])))
	}
	j = j[:len(j)-1] + "}"

	//fmt.Println(j)

	return j
}

