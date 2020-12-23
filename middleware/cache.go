package middleware

import (
	"fmt"
	"net/http"
)

const (
	cacheControlHeaderKey     = "Cache-Control"
	cacheControlNoCacheString = "no-cache"
	cacheControlSecondsString = "public, max-age=%v"

	CacheDuration1Hour     = 3600
	CacheDuration1Minute   = 60
	CacheDuration15Seconds = 15
)

func CacheForNSeconds(w http.ResponseWriter, seconds int) {
	expiry := fmt.Sprintf(cacheControlSecondsString, seconds)
	w.Header().Set(cacheControlHeaderKey, expiry)
}

func DontCache(w http.ResponseWriter) {
	w.Header().Set(cacheControlHeaderKey, cacheControlNoCacheString)
}

func NoCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		DontCache(w)
		next.ServeHTTP(w, r)
	})
}

func CacheFor1Hour(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CacheForNSeconds(w, CacheDuration1Hour)
		next.ServeHTTP(w, r)
	})
}

func CacheFor1Minute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CacheForNSeconds(w, CacheDuration1Minute)
		next.ServeHTTP(w, r)
	})
}

func CacheFor15Seconds(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CacheForNSeconds(w, CacheDuration15Seconds)
		next.ServeHTTP(w, r)
	})
}
