package iotutil

import "strings"

func ConnectPath(params ...string) string {
	res := []string{}
	for _, p := range params {
		if p != "" {
			res = append(res, p)
		}
	}
	return strings.Join(res, "/")
}
