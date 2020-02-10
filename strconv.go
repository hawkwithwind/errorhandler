package errorhandler

import "strconv"

func (o *ErrorHandler) ParseInt(s string, base int) int64 {
	if o.Err != nil {
		return 0
	}

	var i64 int64
	i64, o.Err = strconv.ParseInt(s, base, 64)
	return i64
}

func (o *ErrorHandler) ParseUint(s string, base int) uint64 {
	if o.Err != nil {
		return 0
	}

	var u64 uint64
	u64, o.Err = strconv.ParseUint(s, base, 64)
	return u64
}

