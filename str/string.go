package str

import (
	"fmt"
	"strings"
)

func PrefixAndZeroPad(original interface{}, prefix string, count int, includePrefixInPadLength bool) string {
	padLength := count
	if includePrefixInPadLength {
		padLength = count - len(prefix)
	}

	return fmt.Sprintf("%s%s", prefix, LeftPad(original, "0", padLength))
}

func ZeroPad(original interface{}, count int) string {
	return LeftPad(original, "0", count)
}

func LeftPad(original interface{}, padder string, count int) string {
	padFunc := func(padding, original string) (string, string) { return padding, original }
	return pad(original, padder, padFunc, count)
}

func RightPad(original interface{}, padder string, count int) string {
	padFunc := func(padding, original string) (string, string) { return original, padding }
	return pad(original, padder, padFunc, count)
}

func pad(original interface{}, padder string, orderFunc func(string, string) (string, string), count int) string {
	originalStr := fmt.Sprintf("%s", original)
	if count <= 0 {
		return originalStr
	}

	var needed int
	needed = count - len(originalStr)
	if needed <= 0 {
		return originalStr
	}

	padding := strings.Repeat(padder, needed)
	left, right := orderFunc(padding, originalStr)
	return fmt.Sprintf("%s%s", left, right)
}
