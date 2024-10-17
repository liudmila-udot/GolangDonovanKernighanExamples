package main

import "time"

func IsRFC3339V2(datetime string) bool {
	_, err := time.Parse(time.RFC3339, datetime)
	if nil != err {
		return false
	}
	return true
}

func main() {
	println(IsRFC3339V2("2006-01-02T15:04:05Z07:00"))
}
