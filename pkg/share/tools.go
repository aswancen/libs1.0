package share

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func IsZeroTime(t time.Time) bool {
	// 判断给定时间是否和默认零时间的时间戳相同
	var zero = time.Time{}
	if t.UnixNano() == zero.UnixNano() {
		return true
	}
	return false
}

func PbTimeToTime(t *timestamppb.Timestamp) time.Time {
	if t == nil {
		return time.Time{}
	}
	return time.Unix(t.Seconds, 0)
}

type SliceFunc[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64] func() []T

// AllElementsEqual 判断切片中的所有元素是否相等
func AllElementsEqual[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](f SliceFunc[T]) bool {
	slice := f()
	if len(slice) <= 1 {
		return true
	}
	first := slice[0]
	for _, element := range slice {
		if element != first {
			return false
		}
	}
	return true
}
