package deadlock

import (
	"runtime"
    "bytes"
	"strconv"
)

type GoID struct{}

func (GoID) ExtractGID(s []byte) int64 {
	s = s[len("goroutine "):]
	s = s[:bytes.IndexByte(s, ' ')]
	gid, _ := strconv.ParseInt(string(s), 10, 64)
	return gid
}

func (x GoID) Get() int64 {
	var buf [64]byte
	return x.ExtractGID(buf[:runtime.Stack(buf[:], false)])
}