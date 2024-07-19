package fttl

import "time"

type Policy struct {
	Expire time.Time     `json:"expire"`
	Access time.Duration `json:"access"`
}
