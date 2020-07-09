package app

import "time"

func ParserTime(val string) (*time.Time,error) {
	format := "2006-01-02 15:04:05"
	//format := "20060102"
	theTime,err := time.ParseInLocation(format,val,time.Local)
	return &theTime,err
}

func ParserDate(val string) (*time.Time,error) {
	format := "2006-01-02"
	theTime,err := time.ParseInLocation(format,val,time.Local)
	return &theTime,err
}