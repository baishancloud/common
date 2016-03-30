package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func SplitFieldsFloat(s string) (err error, fields map[string]float64) {
	err = nil
	fields = make(map[string]float64)

	s = strings.Replace(s, " ", "", -1)
	if s == "" {
		return
	}
	var vv float64
	fieldSlice := strings.Split(s, ",")
	for _, field := range fieldSlice {
		field_pair := strings.SplitN(field, "=", 2)
		if len(field_pair) == 2 {
			vv, err = strconv.ParseFloat(field_pair[1], 64)
			if err != nil {
				err = fmt.Errorf("bad field %s", field)
				return
			}
			fields[field_pair[0]] = vv
		} else {
			err = fmt.Errorf("bad field %s", field)
			return
		}
	}

	return
}
func RateAllFields(curfields map[string]float64, prefields map[string]float64, diffTs int64) (err error, ratefields map[string]float64) {
	err = nil
	ratefields = make(map[string]float64, len(curfields))
	keys := make([]string, len(curfields))
	i := 0
	for key, _ := range curfields {
		keys[i] = key
		i++
	}
	for _, k := range keys {
		ratefields[k] = (curfields[k] - prefields[k]) / float64(diffTs)
	}
	return
}
