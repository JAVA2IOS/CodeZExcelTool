package gloableConfig

import (
	"sync"
)

const (
	TimeForamt_yyyyMMddhhmmss = "2006-01-02 15:04:05"
	TimeForamt_yyyy_MM_dd_hh_mm_ss = "2006_01_02_15_04_05"
	TimeForamt_yyyyMMdd = "2006/01/02"
	TimeForamt_yyyy_MM_dd = "2006-01-02"
	TimeForamt_hhmmss = "15:04:05"
	TimeForamt_hh_mm_ss = "15_04_05"
)

type CZGloable struct {
	
}

var czInstance *CZGloable
var one sync.Once

func Instance() *CZGloable {
	one.Do(func(){
		czInstance = &CZGloable{}
	})
	return czInstance
}