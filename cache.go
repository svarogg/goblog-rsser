package main

import "time"

type cacheStruct struct {
	time    time.Time
	content string
}

var cache = cacheStruct{}

func isCacheHot() bool {
	return time.Now().Sub(cache.time) < time.Hour*12
}

func getCache() string {
	return cache.content
}

func setCache(content string) {
	cache.content = content
	cache.time = time.Now()
}
