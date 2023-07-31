package util

/**
Random Log Record Generator
*/

import (
	"math/rand"
	"time"
)

//GenRandVal interface for getting a random value
type GenRandVal func(*RandLog) interface{}

//RandLog random log simulator
type RandLog struct {
	UrandomSpace []byte
	FrandomSpace []byte
	TmstampEnd   int64
	TmstampMod   int64
	GenRecFields map[string]GenRandVal
	MaxFeeds     int
	MaxUsers     int
	ContextArr   []string
	SeedLen      int
	FeedLen      int

	users []string
	feeds []string
}

// ContextArr  context array
var ContextArr = []string{"like", "view", "share", "save"}

// RandSpace allowed random chars in user and feed keys
var RandSpace = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"

// SeedLen random generator seed length
var SeedLen = 28

// RandRecID create a random record id
func RandRecID() string {
	ret := make([]byte, SeedLen)
	for i := 0; i < SeedLen; i++ {
		ret[i] = RandSpace[rand.Intn(len(RandSpace))]
	}
	return string(ret)
}

//NewRandLog create a new RandLog struct
func NewRandLog() *RandLog {

	ret := &RandLog{
		UrandomSpace: []byte(RandSpace), FrandomSpace: []byte(RandSpace),
		TmstampEnd: time.Now().UTC().UnixNano() / 1000000,
		TmstampMod: 1000, MaxFeeds: 50, MaxUsers: 1000,
		ContextArr: ContextArr, GenRecFields: RecLogsFields,
		SeedLen: 28, FeedLen: 28,
	}

	return ret
}

// Next gen random record
func (rec *RandLog) Next() map[string]interface{} {

	ret := make(map[string]interface{})
	for k, frnd := range rec.GenRecFields {
		ret[k] = frnd(rec)
	}
	return ret
	/*
		return map[string]interface{}{
			"context":  "like",
			"feedId":   "-LTbAwQURqnS6LeEOwTu",
			"loggerId": "simulator",
			"tmstamp":  1544698187530,
			"userId":   "7PijQPUEFXbV3HGedjW9gQSiCNk1",
		}
	*/
}

// {"context":"like","feedId":"-LTbAwQURqnS6LeEOwTu","loggerId":"functions","tmstamp":1544698187530,"userId":"7PijQPUEFXbV3HGedjW9gQSiCNk1"}

// RecLogsFields Generator functions
var RecLogsFields = map[string]GenRandVal{
	"context":  genContext,
	"feedId":   genFeedID,
	"loggerId": genLoggerID,
	"tmstamp":  genTmstamp,
	"userId":   genUserID,
}

func genContext(rec *RandLog) interface{} {
	return rec.ContextArr[rand.Intn(len(rec.ContextArr))]
}
func genFeedID(rec *RandLog) interface{} {

	if len(rec.feeds) == 0 || len(rec.feeds) < rec.MaxFeeds {
		ret := make([]byte, rec.FeedLen)
		for i := 0; i < rec.FeedLen; i++ {
			ret[i] = rec.FrandomSpace[rand.Intn(len(rec.FrandomSpace))]
		}
		rec.feeds = append(rec.feeds, string(ret))
		return string(ret)
	}
	return rec.feeds[rand.Intn(len(rec.feeds))]

}
func genLoggerID(rec *RandLog) interface{} {
	return "ramdomizer"
}
func genTmstamp(rec *RandLog) interface{} {
	return rec.TmstampEnd - rand.Int63n(rec.TmstampMod)
}
func genUserID(rec *RandLog) interface{} {
	if len(rec.users) == 0 || len(rec.users) < rec.MaxUsers {
		ret := make([]byte, rec.SeedLen)
		for i := 0; i < rec.SeedLen; i++ {
			ret[i] = rec.UrandomSpace[rand.Intn(len(rec.UrandomSpace))]
		}
		rec.users = append(rec.users, string(ret))
		return string(ret)
	}
	return rec.users[rand.Intn(len(rec.users))]

}
